// 练习5.16：编写多参数版本的strings.Join。

package main

import (
	"fmt"
	"strings"
)

func join(sep string, a ...string) string {
	if len(a) == 0 {
		return ""
	}
	var b strings.Builder
	b.WriteString(a[0])
	for _, s := range a[1:] {
		b.WriteString(sep)
		b.WriteString(s)
	}
	return b.String()
}

func main() {
	s := join(",", "a", "b", "c", "d", "e", "f")
	fmt.Println(s)
}
