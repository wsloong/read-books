// 练习 5.3： 编写函数输出所有text结点的内容。注意不要访问<script>和<style>元素,因为这些元素
// 对浏览者是不可见的。

package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "exercises 5.3: %v\n", err)
		os.Exit(1)
	}

	for _, line := range visit(nil, doc) {
		fmt.Println(line)
	}
}

func visit(s []string, n *html.Node) []string {
	if n.Type == html.TextNode {
		s = append(s, n.Data)
	}

	for c := n.FirstChild; c != nil; c = n.NextSibling {
		if c.Data == "script" || c.Data == "style" {
			continue
		}
		s = visit(s, c)
	}
	return s
}
