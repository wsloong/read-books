package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		cnn, err := listener.Accept()
		if err != nil {
			log.Panicln(err)
			continue
		}
		go handleConn(cnn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()

	for {
		if _, err := io.WriteString(c, time.Now().Format("15:04:05\n")); err != nil {
			return
		}
		time.Sleep(time.Second)
	}
}
