// 这个示例演示了，使用关闭一个通道，通知所有的goroutine
// 同时管理多个goroutine，可以使用context（上下文管理器）
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

var done = make(chan struct{})
var fileSizes = make(chan int64)

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func main() {
	// 从标准输入中读取
	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()

	for {
		select {
		case <-done:
			for range fileSizes {
				//
			}
			return
		case size, ok := <-fileSizes:
			fmt.Println(size, ok)
		}
	}
}

func walkDir(dir string, wg *sync.WaitGroup, fileSizes chan<- int64) {
	defer wg.Done()
	if cancelled() {
		return
	}
	for _, entry := range dirents(dir) {
		entry = entry
	}
}

func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du4:%v\n", err)
		return nil
	}
	return entries
}
