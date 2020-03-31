// 练习 5.2： 编写函数，记录在HTML树中出现的同名元素的次数。

package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "exercises 5.2: %v\n", err)
		os.Exit(1)
	}

	for k, v := range count(nil, doc) {
		fmt.Println(k, ">>>>", v)
	}
}

func count(m map[string]int, n *html.Node) map[string]int {
	if m == nil {
		m = make(map[string]int)
	}
	if n.Type == html.ElementNode {
		m[n.Data]++
	}

	if n.FirstChild != nil {
		m = count(m, n.FirstChild)
	}

	if n.NextSibling != nil {
		m = count(m, n.NextSibling)
	}
	return m
}
