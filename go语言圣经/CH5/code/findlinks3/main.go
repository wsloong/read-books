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

func main() {
	// 使用广度优先爬取命令行出入的url
	breadthFirst(crawl, os.Args[1:])
}
