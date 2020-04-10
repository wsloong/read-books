// 练习 7.4： strings.NewReader函数通过读取一个string参数返回一个满足io.Reader接口类型的值（和其它值）。
// 实现一个简单版本的NewReader，并用它来构造一个接收字符串输入的HTML解析器（§5.2）

package main

import (
	"fmt"
	"io"
)

type StringReader struct {
	content string
	n       int
}

func (s *StringReader) Read(p []byte) (int, error) {
	if s.n >= len(s.content) {
		return 0, io.EOF
	}

	data := s.content[s.n:]
	n := copy(p, data)
	s.n += n
	return n, nil
}

func NewStringReader(s string) *StringReader {
	r := &StringReader{
		content: s,
	}
	return r
}

func main() {
	reader := NewStringReader("hello world")
	data := make([]byte, 11)
	n, err := reader.Read(data)
	for err == nil {
		fmt.Println(n, string(data[0:n]))
		n, err = reader.Read(data)
	}
}
