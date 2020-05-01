// ./du2 -v $HOME /usr /bin /etc
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)
// -v 显示程序进度信息，每500ms展示一次
var verbose = flag.Bool("v", false, "show verbose progress messages")

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots)== 0 {
		roots = []string{"."}
	}
	fileSizes := make(chan int64)

	go func() {
		for _, root := range roots {
			walkDir(root, fileSizes)
		}
		close(fileSizes)
	}()

	// tick 默认为nil，select永远不会选中这个case
	var tick <- chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	var nfiles, nbytes int64

loop:
	for {
		select {
			case size, ok := <-fileSizes:
				if !ok {
					// 同时终结select和for两个循环。
					// 如果单纯使用break，for循环会进入下一轮select循环
					break loop
				}
				nfiles++
				nbytes += size
			case <- tick:
				printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

func walkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func dirents(dir string)[]os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du2:%v\n", err)
		return nil
	}
	return entries
}