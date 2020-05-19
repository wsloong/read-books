package memo

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

var urls = []string{
	"https://golang.google.cn/",
	"https://golang.google.cn/pkg/",
	"https://golang.google.cn/doc/",
	"https://golang.google.cn/",
	"https://golang.google.cn/pkg/",
	"https://golang.google.cn/doc/",
}

// 顺序调用
func TestDo1(t *testing.T) {
	m := New(httpGetBody)
	ctx := context.TODO()

	for _, url := range urls {
		start := time.Now()
		value, err := m.Get(ctx, url)
		if err != nil {
			fmt.Printf("error:%s\n", err)
		}
		fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
	}
}

// 并发调用，因为不是并发安全的，会出现bug
func TestDo2(t *testing.T) {
	m := New(httpGetBody)
	var wg sync.WaitGroup
	ctx := context.TODO()

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			start := time.Now()
			value, err := m.Get(ctx, url)

			if err != nil {
				fmt.Printf("url:%s, error:%s\n", url, err)
				return
			}
			fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))

		}(url)
	}
	wg.Wait()
}
