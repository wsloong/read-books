// 演示单方向的channel

package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// 将 chan int类型隐式转换为 Chan<- int单向只发送channel类型
	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}

func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}
