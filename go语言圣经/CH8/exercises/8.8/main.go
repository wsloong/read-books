// 练习 8.8： 使用select来改造8.3节中的echo服务器，为其增加超时，
// 这样服务器可以在客户端10秒中没有任何喊话时自动断开连接。
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	message := make(chan string)
	go func() {
		for {
			select {
			case <- time.After(10 * time.Second):
				c.Close()
			case msg := <-message:
				go echo(c, msg, time.Second)
			}
		}
	}()

	for input.Scan() {
		message <- input.Text()
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintf(c, "\t%s\n", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintf(c, "\t%s\n", shout)
	time.Sleep(delay)
	fmt.Fprintf(c, "\t%s\n", strings.ToLower(shout))
}
