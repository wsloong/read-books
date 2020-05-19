// 练习 9.3： 扩展Func类型和 (*Memo).Get 方法，
// 支持调用方提供一个可选的done channel，
// 使其具备通过该channel来取消整个操作的能力(§8.9)。
// 一个被取消了的Func的调用结果不应该被缓存

// 该程序包含多个goroutine，更应该使用context上下文管理器来管理

package memo

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Func func(ctx context.Context, key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{}
}

type request struct {
	key      string
	response chan<- result
	ctx      context.Context
}

type Memo struct {
	requests chan request
}

func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(ctx context.Context, key string) (interface{}, error) {
	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("canceled")
	default:
		response := make(chan result)
		memo.requests <- request{key, response, ctx}
		res := <-response
		return res.value, res.err
	}
}

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		select {
		case <-req.ctx.Done():
			continue
		default:
			e := cache[req.key]
			if e == nil {
				e = &entry{ready: make(chan struct{})}
				cache[req.key] = e
				go e.call(req.ctx, f, req.key)
			}
			go e.deliver(req.ctx, req.response)
		}
	}
}

func (e *entry) call(ctx context.Context, f Func, key string) {
	e.res.value, e.res.err = f(ctx, key)
	close(e.ready)
}

func (e *entry) deliver(ctx context.Context, response chan<- result) {
	select {
	case <-ctx.Done():
		return
	default:
		// 等待值就绪
		<-e.ready
		// 发送结果到客户端
		response <- e.res
	}
}

func httpGetBody(ctx context.Context, url string) (interface{}, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
