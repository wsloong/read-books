// 练习 8.2： 实现一个并发FTP服务器。服务器应该解析客户端来的一些命令，
// 比如cd命令来切换目录，ls来列出目录内文件，get和send来传输文件，close来关闭连接。
// 你可以用标准的ftp命令来作为客户端，或者也可以自己实现一个

// 客户端

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		conn.Write([]byte(line))

		resp := make([]byte, 1024)
		_, err := conn.Read(resp)
		if err != nil {
			return
		}
		fmt.Println(string(resp))
	}
}

