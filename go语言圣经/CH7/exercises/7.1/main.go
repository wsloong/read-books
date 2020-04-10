// 练习 7.1： 使用来自ByteCounter的思路，实现一个针对对单词和行数的计数器。你会发现bufio.ScanWords非常的有用。
package main

import (
	"bufio"
	"fmt"
	"strings"
)

type WordsCounter int

func (w *WordsCounter) Write(p []byte) (int, error) {
	bs := bufio.NewScanner(strings.NewReader(string(p)))
	bs.Split(bufio.ScanWords)
	var total int
	for bs.Scan() {
		total++
	}
	*w += WordsCounter(total)
	return total, nil
}

type LineCounter int

func (l *LineCounter) Write(p []byte) (int, error) {
	bs := bufio.NewScanner(strings.NewReader(string(p)))
	bs.Split(bufio.ScanLines)
	var total int
	for bs.Scan() {
		total++
	}
	*l += LineCounter(total)
	return total, nil
}

func main() {
	var words WordsCounter
	words.Write([]byte("hello world"))
	words.Write([]byte("hello"))
	fmt.Println(words)

	var lines LineCounter
	lines.Write([]byte("good luck\r\ngood bay"))
	fmt.Println(lines)
}
