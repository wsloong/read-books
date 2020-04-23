// 退出时候(Ctrl+d),友好的处理后台goroutine
package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})

	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{}
	}()

	// 当关闭了标准输入(Ctrl+d,不是Ctrl+c)，该函数会返回，然后调用Close()关闭读写方向的网络链接
	// server程序会收到一个文件(end-of-file)的结束信号
	// 后台goroutine的io.Copy()函数返回一个错误(这里没有处理错误)
	// 然后打印done，向无缓存通道写入空的struct
	mustCopy(conn, os.Stdin)
	conn.Close()
	// 这里会阻塞，直到有goroutine发送数据
	<-done
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
