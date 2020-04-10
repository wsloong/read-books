// 练习 7.2： 写一个带有如下函数签名的函数CountingWriter，传入一个io.Writer接口类型，返回一个新的Writer类型
// 把原来的Writer封装在里面和一个表示写入新的Writer字节数的int64类型指针
// func CountingWriter(w io.Writer) (io.Writer, *int64)
package main

import (
	"fmt"
	"io"
	"os"
)

type CountWriter struct {
	Writer io.Writer
	Count  int64
}

func (c *CountWriter) Write(p []byte) (int, error) {
	n, err := c.Writer.Write(p)
	if err != nil {
		return n, err
	}
	c.Count += int64(n)
	return n, nil
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	newWriter := CountWriter{
		Writer: w,
	}
	return &newWriter, &(newWriter.Count)
}

func main() {
	w, counter := CountingWriter(os.Stdout)
	fmt.Fprintf(w, "%s", "something to do....")
	fmt.Println(*counter)

	w.Write([]byte("hello world"))
	fmt.Println(*counter)
}
