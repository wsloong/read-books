// 练习5.12： gopl.io/ch5/outline2（5.5节）的startElement和endElement共用了全局变量depth，
// 将它们修改为匿名函数，使其共享outline中的局部变量
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}

}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("getting %s:%s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	var depth int
	startElement := func(n *html.Node) {
		if n.Type == html.ElementNode {
			fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
			depth++
		}
	}

	endElement := func(n *html.Node) {
		if n.Type == html.ElementNode {
			depth--
			fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		}
	}

	forEachNode(doc, startElement, endElement)
	return nil
}

// forEachNode 针对每个节点x，都会调用pre(x)和post(x)
// pre和post都是可选的
// 遍历孩子节点之前，pre被调用
// 遍历孩子节点知乎，post被调用
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}

}
