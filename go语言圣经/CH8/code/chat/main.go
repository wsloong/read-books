// 一个聊天的示例
// 客户端使用 code/netcat3的代码
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// 消息输出通道
type client chan<- string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // 所有的消息通道
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}
}

// broadcaster 广播
func broadcaster() {
	clients := make(map[client]bool) // 所有连接的client
	for {
		select {
		case msg := <-messages:
			// 将消息广播到所有的客户端
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWrite(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	// NOTE: 忽略来自输入的潜在错误

	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

func clientWrite(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
