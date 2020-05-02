// 练习 8.11： 紧接着8.4.4中的mirroredQuery流程，实现一个并发请求url的fetch的变种。当第
//一个请求返回时，直接取消其它的请求。

package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res := mirroredQuery()
	fmt.Println(res)
}

func mirroredQuery() string {
	responses := make(chan string, 3)
	ctx, cancelFunc := context.WithCancel(context.Background())
	go func() { responses <- request(ctx, "https://www.baidu.com") }()
	go func() { responses <- request(ctx, "https://golang.google.cn") }()
	go func() { responses <- request(ctx, "https://americas.gopl.io") }()

	for {
		select {
		case res := <-responses:
			cancelFunc()
			return res
		}
	}
}
func request(ctx context.Context, hostname string) (response string) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, hostname, nil)
	if err != nil {
		fmt.Printf("host:%s, err:%v\n", hostname, err)
		return ""
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("do host:%s, err:%v\n", hostname, err)
		return ""
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read body host:%s, err:%v\n", hostname, err)
		return ""
	}
	return string(b)
}
