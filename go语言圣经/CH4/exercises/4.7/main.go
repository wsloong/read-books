// 练习 4.7： 修改reverse函数用于原地反转UTF-8编码的[]byte。是否可以不用分配额外的内存？

package main

import (
	"fmt"
	"unicode/utf8"
)

func reverseUTF8(b []byte) {
	for l := len(b); l > 0; {
		r, size := utf8.DecodeRuneInString(string(b[0:]))
		copy(b[0:], b[size:])
		copy(b[l-size:], []byte(string(r)))
		l -= size
	}

}

func main() {
	b := []byte("中国")
	reverseUTF8(b)
	fmt.Println(string(b))
}
