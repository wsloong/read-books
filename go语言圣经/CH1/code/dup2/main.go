// 打印标准输入或者文件中出现多次的每行文本，并打印出现次数
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// map是用make函数创建的数据结构的一个引用
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// 虽然这里的 counts 是一个引用的拷贝，但是这个拷贝和原值都指向了同一个内存块
// 这里对map进行的任何改变都会影响到原来的map
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		counts[input.Text()]++
	}
}
