package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <- ticker.C:
			fmt.Println("in tick...")
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		}
	}
	launch()
}

func launch() {

}

