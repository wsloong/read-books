// 练习 8.10： HTTP请求可能会因http.Request结构体中Cancel channel的关闭而取消。修改
//8.6节中的web crawler来支持取消http请求。（提示：http.Get并没有提供方便地定制一个请
//求的方法。你可以用http.NewRequest来取而代之，设置它的Cancel字段，然后用
//http.DefaultClient.Do(req)来进行这个http请求。）
package main

import (
	"context"
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
	"sync/atomic"
	"time"
)

// 限制20个并发
// 类似令牌桶的做法
var tokens = make(chan struct{}, 20)

// FOR 练习8.6
const DEPTH = 6

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	worklist := make(chan []string)
	unseenLinks := make(chan string)

	// 这里设置1s后，执行取消函数，这时候ctx.Done()会收到一个信号
	// 这种处理，还是会发送请求，但是中途会报 context canceled 的错误
	time.AfterFunc(time.Second, cancelFunc)

	go func() {
		worklist <- os.Args[1:]
	}()
	var depth int32

	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(ctx, link)
				// 爬到的链接在一个专有的goroutine中发送到worklist来避免死锁
				go func() {
					atomic.AddInt32(&depth, 1)
					worklist <- foundLinks
				}()
			}
		}()
	}

	seen := make(map[string]bool)

	for {
		select {
		case list := <-worklist:
			if depth == DEPTH {
				close(unseenLinks)
				break
			}
			for _, link := range list {
				if !seen[link] {
					seen[link] = true
					unseenLinks <- link
				}
			}
		case <-ctx.Done():
			return
		}
	}
}

// 抓取页面
func crawl(ctx context.Context, url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // 获取一个token
	list, err := Extract(ctx, url)
	<-tokens // 释放一个token
	if err != nil {
		log.Print(err)
	}
	return list
}

// 解析HTTP GET到URL的内容，并返回HTML中的链接
func Extract(ctx context.Context, url string) ([]string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s:%s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML:%v", url, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}

				// 使用 resp.Request.Url.Parse可以解析成完成的url
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue
				}
				links = append(links, link.String())
			}
		}
	}

	forEachNode(doc, visitNode, nil)
	return links, nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}
