// 打印标准输入中出现多次的每行文本，并打印出现次数
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	// Scanner 实现接收输入，并把输入打散成行或者单词
	// 通常是处理行形式的输入最简单的方法了
	input := bufio.NewScanner(os.Stdin)

	// Scan 方法读到新行返回true，否则返回false
	for input.Scan() {
		// 当输入q时候结束当前循环
		if input.Text() == "q" {
			break
		}
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
