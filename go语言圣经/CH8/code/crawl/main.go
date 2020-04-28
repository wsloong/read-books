package main

import (
	"fmt"
	"log"
	"os"
	"sync/atomic"
)

// 限制20个并发
// 类似令牌桶的做法
var tokens = make(chan struct{}, 20)

// 抓取页面
func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // 获取一个token
	list, err := Extract(url)
	<-tokens // 释放一个token
	if err != nil {
		log.Print(err)
	}
	return list
}

func main1() {
	worklist := make(chan []string)

	var n int // 等待发送到worklist的个数
	n++       // 默认情况下会把命令行的参数发送worklist中， 为什么不写成n := 1?

	// start with the command-line arguments
	go func() {
		worklist <- os.Args[1:]
	}()

	// Crawl  the web concurrently
	seen := make(map[string]bool)

	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}

	// 下面的代码程序永远不会终止
	// 即使爬到了所有初始链接衍生出的链接
	// 所以使用上面的程序代替
	// for list := range worklist {
	// 	for _, link := range list {
	// 		if !seen[link] {
	// 			seen[link] = true
	// 			go func(link string) {
	// 				worklist <- crawl(link)
	// 			}(link)
	// 		}
	// 	}
	// }
}

// 另一个版本的限制并发
// Go语言实战中有类似的例子
// 相比之下，我更喜欢这个版本

// FOR 练习8.6
const DEPTH = 2

func main() {
	worklist := make(chan []string)
	unseenLinks := make(chan string)

	go func() {
		worklist <- os.Args[1:]
	}()
	var depth int32

	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				// 爬到的链接在一个专有的goroutine中发送到worklist来避免死锁
				go func() {
					atomic.AddInt32(&depth, 1)
					worklist <- foundLinks
				}()
			}
		}()
	}

	seen := make(map[string]bool)
	for list := range worklist {
		// 到达指定深度，关闭(unseenLinks),break
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
	}
}
