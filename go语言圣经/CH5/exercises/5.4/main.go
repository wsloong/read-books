// 练习 5.4： 扩展vist函数，使其能够处理其他类型的结点，如images、scripts和style sheets。

package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1:%v\n", err)
		os.Exit(1)
	}

	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

// visit附加到链接，在n中找到每个链接，并返回结果
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && (n.Data == "link" || n.Data == "a") {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "img") {
		for _, a := range n.Attr {
			if a.Key == "src" {
				links = append(links, a.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
