// 练习 8.2： 实现一个并发FTP服务器。服务器应该解析客户端来的一些命令，
// 比如cd命令来切换目录，ls来列出目录内文件，get和send来传输文件，close来关闭连接。
// 你可以用标准的ftp命令来作为客户端，或者也可以自己实现一个

// 服务端
package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
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
	defer c.Close()

	for {
		b := make([]byte, 1024)
		_, err := c.Read(b)
		if err != nil {
			log.Println("read error:", err)
			return
		}

		s := string(b)
		if strings.Contains(s, "ls") {
			files, err := ioutil.ReadDir("./")
			if err != nil {
				log.Println("read dir error:", err)
				return
			}

			var b bytes.Buffer

			for _, f := range files {
				_s := fmt.Sprintf("dir:%t\t%d\t%s\n",f.IsDir(), f.Size(), f.Name())
				b.Write([]byte(_s))
			}
			c.Write(b.Bytes())
		} else if strings.Contains(s, "get") {
			// 这里固定返回一个文件作为演示
			f, err := os.OpenFile("1.txt", os.O_RDONLY, 0666)
			if err != nil {
				log.Println("open file:", err)
				return
			}
			io.Copy(c, f)
			f.Close()
		}
		// ....
	}
}
