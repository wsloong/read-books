// 练习 8.12： 使broadcaster能够将arrival事件通知当前所有的客户端。为了达成这个目的，你
// 需要有一个客户端的集合，并且在entering和leaving的channel中记录客户端的名字。
// TODO broadcaster的clients难道不是？

// 练习 8.13： 使聊天服务器能够断开空闲的客户端连接，比如最近五分钟之后没有发送任何消
//息的那些客户端。提示：可以在其它goroutine中调用conn.Close()来解除Read调用，就像
//input.Scanner()所做的那样。

//练习 8.14： 修改聊天服务器的网络协议这样每一个客户端就可以在entering时可以提供它们
//的名字。将消息前缀由之前的网络地址改为这个名字。
// TODO 网络协议?

//练习 8.15： 如果一个客户端没有及时地读取数据可能会导致所有的客户端被阻塞。修改
//broadcaster来跳过一条消息，而不是等待这个客户端一直到其准备好写。或者为每一个客户
//端的消息发出channel建立缓冲区，这样大部分的消息便不会被丢掉；broadcaster应该用一个
//非阻塞的send向这个channel中发消息

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

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
			// 将消息广播到所有的客户端 for 8.15
			for cli, bool := range clients {
				if !bool {
					continue
				}
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
	// for 8.13; 创建一个通道，将客户端的输入放到这个channel中
	clientMsg := make(chan string)
	go func() {
		for {
			select {
			case <-time.After(5 * time.Minute):
				conn.Close()
				return
			case s := <-clientMsg:
				messages <- s

			}
		}
	}()
	for input.Scan() {
		clientMsg <- who + ": " + input.Text()
	}

	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

func clientWrite(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
