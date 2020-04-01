// 练习 5.9： 编写函数expand，将s中的"foo"替换为f("foo")的返回值。

package main

import (
	"fmt"
	"strings"
)

func main() {
	f := func(s string) string {
		return "123"
	}

	res := expand("foobarfool", f)
	fmt.Println(res)
}

func expand(s string, f func(string) string) string {
	res := f("fool")
	return strings.ReplaceAll(s, "foo", res)
}
