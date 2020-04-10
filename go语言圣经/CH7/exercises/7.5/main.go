// 练习 7.5： io包里面的LimitReader函数接收一个io.Reader接口类型的r和字节数n，
// 并且返回另一个从r中读取字节但是当读完n个字节后就表示读到文件结束的Reader。
// 实现这个LimitReader函数：func LimitReader(r io.Reader, n int64) io.Read
package main

import (
	"fmt"
	"io"
)

type LimitedReader struct {
	Reader io.Reader
	Limit  int // 限制要读的字节个数
}

func (l *LimitedReader) Read(b []byte) (int, error) {
	if l.Limit <= 0 {
		return 0, io.EOF
	}

	if len(b) > l.Limit {
		b = b[0:l.Limit]
	}

	n, err := l.Reader.Read(b)
	l.Limit -= n
	return n, err
}

func LimitReader(r io.Reader, limit int) io.Reader {
	return &LimitedReader{
		Reader: r,
		Limit:  limit,
	}
}

func main() {
	// reader := strings.NewReader("hello world")
	// lr := LimitReader(reader, 5)
	// b := make([]byte, 11)
	// n, err := lr.Read(b)
	// if err != nil {
	// 	fmt.Println("Get a err: ", err)
	// 	os.Exit(-1)
	// }
	// fmt.Println(string(b), n)

	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	_ = IntSet{}.String()
	var s IntSet

	var _ fmt.Stringer = s
	var _ fmt.Stringer = &s
	fmt.Println(s, &s)

}

type IntSet struct{}

func (i IntSet) String() string {
	return "3232323232"
}
