// 练习7.17：扩展xmlselect程序以便让元素不仅仅可以通过名称选择，也可以通过它们CSS样式上属性进行选择；
// 例如一个像这样的元素<div id="page" class="wide">可以通过匹配id或者class同时还有它的名称来进行选择。

package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []string

	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlslect:%v\n", err)
			os.Exit(1)
		}

		switch tok := tok.(type) {
		case xml.StartElement:
			item := "<" + tok.Name.Local
			for _, attr := range tok.Attr {
				if attr.Name.Local != "class" && attr.Name.Local != "id" {
					continue
				}
				item += fmt.Sprintf(" %s=\"%s\"", attr.Name.Local, attr.Value)
			}
			item += ">"
			stack = append(stack, item)
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			if containAll(stack, os.Args[1:]) {
				fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
			}
		}
	}
}

func containAll(x, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}

		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]

	}
	return false
}
