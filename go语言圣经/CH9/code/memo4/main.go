// 该版本在memo3的基础上 进行重复抑制
// 每个map都是指向一个条目的指针，每个条目包含对函数f调用结果内容缓存
// 这次entry还包含一个ready的channel，在条目的结果被设置之后，这个channel就会被关闭
// 可以向其他goroutine广播去读取该条目内的结果是安全的
package memo

import (
	"io/ioutil"
	"net/http"
	"sync"
)

// Func 一个函数的类型
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

type Memo struct {
	f     Func
	mu    sync.Mutex // 使用排它锁保护map
	cache map[string]*entry
}

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	e := memo.cache[key]

	if e == nil {
		// 第一次请求这个key的值 这个goroutine负责计算值和广播值已经就绪的条件
		e = &entry{ready: make(chan struct{})}
		// 插入一条没有准备好的条目，当前goroutine负责调用慢函数，跟新条目，广播值已经就绪了
		memo.cache[key] = e
		memo.mu.Unlock()

		e.res.value, e.res.err = memo.f(key)
		close(e.ready)		// broadcast ready condition
	} else {
		// 重复的请求key，值还没有写入(其他goroutine在调用f这个慢函数)
		memo.mu.Unlock()
		<-e.ready		// 等待就绪状态,在channel关闭之前一直等待
	}	
	return e.res.value, e.res.err
}

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func Do1() {}

func Do2() {}
