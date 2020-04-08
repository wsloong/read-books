// 练习5.13： 修改crawl，使其能保存发现的页面，必要时，可以创建目录来保存这些页面。
// 只保存来自原始域名下的页面。假设初始页面在golang.org下，就不要保存vimeo.com下的页面。
package main

import (
	"fmt"
	"log"
	"os"
)

// breadthFirst 为每个元素调用一次f函数,并将结果放到slice中
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

// 抓取页面
func crawl(url string) []string {
	fmt.Println(url)
	list, err := Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

// 保存初始的URL
var BaseUrls []string

// 当前要解析的url是否在元素URL中
func UrlInBases(url string) bool {
	for _, u := range BaseUrls {
		if u == url {
			return true
		}
	}
	return false
}

func main() {
	BaseUrls = os.Args[1:]
	// 使用广度优先爬取命令行出入的url
	breadthFirst(crawl, BaseUrls)
}
