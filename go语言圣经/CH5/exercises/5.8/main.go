// 练习 5.8： 修改pre和post函数，使其返回布尔类型的返回值。返回false时，中止 forEachNoded 的遍历。
// 使用修改后的代码编写 ElementByID 函数，根据用户输入的id查找第一个拥有该id元素的HTML元素，查找成功后，停止遍历

package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

var url = flag.String("url", "https://golang.google.cn/", "the url to get")
var id = flag.String("id", "", "the id attr in url content")

func main() {
	flag.Parse()

	if *id == "" {
		fmt.Fprintf(os.Stderr, "id can not empty")
		os.Exit(1)
	}

	_, err := findElements(*url, *id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "exercises 5.7: %v\n", err)
		os.Exit(1)
	}
}

func findElements(url, id string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getting %s:%s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	return ElementByID(doc, id, startElement), nil
}

func ElementByID(n *html.Node, id string, pre func(n *html.Node, id string) bool) *html.Node {
	if pre != nil {
		if pre(n, id) {
			return n
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		n = ElementByID(c, id, pre)

	}
	return n
}

// id为HTML元素的属性
func startElement(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == id {
				return true
			}
		}
	}

	return false
}
