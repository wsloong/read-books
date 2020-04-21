package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	// 监听 localhost:8000上的tcp链接
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		// 程序会在这里阻塞，直到一个链接
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		// 处理链接
		handleConn(conn)
	}
}

// 处理过来的链接
func handleConn(c net.Conn) {
	// 使用defer确保程序退出时候关闭链接
	defer c.Close()
	for {
		// 链接实现了reader/writer/closer接口
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(time.Second * 1)
	}
}
