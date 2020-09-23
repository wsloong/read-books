package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"sync"
	"time"
)

var (
	enteringChannel = make(chan *User)		// 新用户到来
	leavingChannel = make(chan *User)		// 用户离开
	messageChannel = make(chan Message, 8)		// 广播专用的用户普通信息,缓冲是尽可能的避免出现异常情况阻塞
)

// 用户信息
type User struct {
	ID int
	Addr string
	EnterAt time.Time
	MessageChannel chan string
}

func (u User) String() string {
	return u.Addr + ", UID:"+ strconv.Itoa(u.ID) + ", Enter At:" + u.EnterAt.Format("2006-01-02 15:04:05+8000")
}

// 消息内容
type Message struct {
	OwnerID int
	Content string
}


func main() {
	listener, err := net.Listen("tcp", ":2020")
	if err != nil {
		panic(err)
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

// broadcaster 用于记录聊天室用户，并进行消息广播
// 1. 新用户进来
// 2 用户普通消息
// 3 用户离开
func broadcaster() {
	users := make(map[*User]struct{})
	for {
		select {
		case user := <- enteringChannel:
			// 新用户进入
			users[user] = struct{}{}
		case user := <-leavingChannel:
			// 用户离开
			delete(users, user)
			// 避免 goroutine 泄露
			close(user.MessageChannel)
		case msg := <- messageChannel:
			// 给所有在线用户发送消息
			for user := range users {
				if user.ID == msg.OwnerID {
					continue
				}

				user.MessageChannel <- msg.Content
			}
		}
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	// 1 新用户进来，构建新用户实例
	user := &User{
		ID: GenUserID(),
		Addr: conn.RemoteAddr().String(),
		EnterAt: time.Now(),
		MessageChannel: make(chan string, 8),
	}

	// 由于当前是在一个新的 goroutine 中进行读操作，所以需要开一个 goroutine 用于写
	// 读写中间通过 channel 进行通信
	// 这里要优先写，不然会收到自己到来的信息
	go sendMessage(conn, user.MessageChannel)

	// 给当前用户发送欢迎消息，想所有用户告知新用户到来
	user.MessageChannel <- "Welcome, " + user.String()
	msg := Message{
		OwnerID: user.ID,
		Content: "user: `" + strconv.Itoa(user.ID) + "` has enter",
	}
	messageChannel <- msg

	// 记录到全局用户列表中，避免用锁
	enteringChannel <- user

	// 踢出超时的用户
	var userActive = make(chan struct{})
	go func() {
		d := 5 * time.Minute
		timer := time.NewTimer(d)
		for {
			select {
			case <- timer.C:
				conn.Close()
			case <- userActive:
				timer.Reset(d)
			}
		}
	}()

	// 循环读取用户输入
	input := bufio.NewScanner(conn)
	for input.Scan() {
		msg.Content =  strconv.Itoa(user.ID) + ":" + input.Text()
		messageChannel <- msg

		// 用户活跃
		userActive <- struct{}{}
	}

	if err := input.Err(); err != nil {
		log.Println("读取错误: ", err)
	}

	// 用户离开
	leavingChannel <- user
	msg.Content = "user: `" + strconv.Itoa(user.ID) + "` has left"
	messageChannel <- msg
}

func sendMessage(conn net.Conn, ch <- chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

// 生成用户 ID
var (
	globalID int
	idLocker sync.Mutex
)

func GenUserID() int {
	idLocker.Lock()
	defer idLocker.Unlock()

	globalID++
	return globalID
}