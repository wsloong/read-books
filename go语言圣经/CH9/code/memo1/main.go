// 版本1的缓存函数， 非并发安全
package memo

import (
	"io/ioutil"
	"net/http"
)

// Func 一个函数的类型
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

// Memo 缓存函数f的结果
type Memo struct {
	f     Func
	cache map[string]result
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

func (memo *Memo) Get(key string) (interface{}, error) {
	res, ok := memo.cache[key]
	if !ok {
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
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

func Do1() {

}

func Do2() {

}
