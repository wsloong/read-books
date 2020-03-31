// 练习 5.5： 实现countWordsAndImages

package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		words, images, err := CountWordsAndImages(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "count words and images: %v\n", err)
			continue
		}

		fmt.Printf("url:%s words:%d, images:%d\n", url, words, images)
	}
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("http status code:%d", resp.StatusCode)
		return
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}

	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	s, images := visit(nil, 0, n)

	for _, v := range s {
		words += strings.Count(v, "")
	}
	return
}

func visit(s []string, imgNum int, n *html.Node) ([]string, int) {
	if n.Type == html.TextNode {
		text := strings.Trim(strings.TrimSpace(n.Data), "\r\n")
		if text != "" {
			s = append(s, text)
		}
	}

	if n.Type == html.ElementNode && n.Data == "img" {
		imgNum++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Data == "script" || c.Data == "style" {
			continue
		}

		s, imgNum = visit(s, imgNum, c)
	}
	return s, imgNum
}
