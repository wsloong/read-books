// 该版本，调用 get 的goroutine会两次获取锁
// 查询阶段获取一次，如果没有，进入更新会再次获取锁。
// 中间阶段的其他goroutine可以随意使用cache

// 依然存留的问题
// 如果两个以上的goroutine同一时刻调用Get来请求同样的URL时候，
// 多个goroutine一起查询cache，发现没有值
// 然后一起调用f这个函数，得到结果后都会更新map，其中一个会覆盖另一个结果
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
	cache map[string]result
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	res, ok := memo.cache[key]
	memo.mu.Unlock()

	if !ok {
		res.value, res.err = memo.f(key)
		memo.mu.Lock()
		memo.cache[key] = res
		memo.mu.Unlock()
	}
	return res.value, res.err
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
