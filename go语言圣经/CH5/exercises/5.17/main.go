// 练习5.17：编写多参数版本的ElementsByTagName，函数接收一个HTML结点树以及任意数量的标签名，
// 返回与这些标签名匹配的所有元素。下面给出了2个例子

package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

const URL = "https://golang.google.cn/"

var nodes []*html.Node

func main() {
	if err := parseURL(os.Args[1:]...); err != nil {
		panic(err)
	}

	for _, n := range nodes {
		fmt.Println(n)
	}

}

func parseURL(name ...string) error {
	if len(name) == 0 {
		return nil
	}

	resp, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("getting %s:%s", URL, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", URL, err)
	}
	ElementsByTagName(doc, name...)
	return nil
}

func ElementsByTagName(n *html.Node, names ...string) []*html.Node {
	for _, name := range names {
		if n.Type == html.ElementNode && n.Data == name {
			nodes = append(nodes, n)
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		nodes = ElementsByTagName(c, names...)
	}
	return nodes
}
