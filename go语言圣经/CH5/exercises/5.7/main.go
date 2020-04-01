// 练习 5.7： 完善startElement和endElement函数，使其成为通用的HTML输出器。
// 要求：输出注释结点，文本结点以及每个元素的属性（< a href='...'>）。
// 使用简略格式输出没有孩子结点的元素（即用<img/>代替<img></img>
// 编写测试，验证程序输出的格式正确。（详见11章）
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

var depth int

func main() {
	for _, url := range os.Args[1:] {
		if err := findElements(url); err != nil {
			fmt.Fprintf(os.Stderr, "exercises 5.7: %v\n", err)
			continue
		}
	}
}

func findElements(url string) error {
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

	forEachNode(doc, startElement, endElement)
	return nil
}

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

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		var attr string
		for _, a := range n.Attr {
			attr += " " + a.Key + "=" + "\"" + a.Val + "\" "
		}
		fmt.Printf("%*s<%s%s", depth*2, "", n.Data, attr)
		depth++
	}

	if n.Type == html.ElementNode && n.FirstChild == nil && n.Data != "script" {
		fmt.Printf("/>\n")
	} else if n.Type == html.ElementNode {
		fmt.Printf(">\n")
	}

	if n.Type == html.TextNode {
		fmt.Printf("%*s %s\n", depth*2, "", n.Data)
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode && n.FirstChild == nil && n.Data != "script" {
		depth--
		fmt.Printf("\n")
		return
	}

	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}
