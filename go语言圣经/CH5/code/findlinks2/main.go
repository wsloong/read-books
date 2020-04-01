package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		links, err := findLinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
			continue
		}

		fmt.Println("url:", url)
		for _, link := range links {
			fmt.Println(link)
		}
	}
}

// HTTP GET 获取URL页面，并返回链接
func findLinks(url string) ([]string, error) {
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

	return visit(nil, doc), nil
}

func visit(s []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				s = append(s, attr.Val)
			}
		}
	}

	if n.FirstChild != nil {
		s = visit(s, n.FirstChild)
	}

	if n.NextSibling != nil {
		s = visit(s, n.NextSibling)
	}
	return s
}
