// 练习 4.6： 编写一个函数，原地将一个UTF-8编码的[]byte类型的slice中相邻的空格（参考unicode.IsSpace）替换成一个空格返回
package main

import (
	"fmt"
	"unicode"
)

func removeSpace(b []byte) []byte {
	for i := 0; i < len(b)-1; i++ {
		if unicode.IsSpace(rune(b[i])) && unicode.IsSpace(rune(b[i+1])) {
			copy(b[i:], b[i+1:])
			b = b[:len(b)-1]
			i--
		}
	}
	return b
}

func main() {
	b := []byte("hello world  world 1  2  中   国")
	fmt.Printf("%q\n", removeSpace(b))
}
