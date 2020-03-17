// 练习 4.9： 编写一个程序 wordfreq 程序，报告输入文本中每个单词出现的频率。
// 在第一次调用Scan前先调用 input.Split(bufio.ScanWords )函数，这样可以按单词而不是按行输入。

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		counts[input.Text()]++
	}

	for k, v := range counts {
		fmt.Println(k, " == ", v)
	}
}
