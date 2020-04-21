package main

import (
	"fmt"
	"time"
)

func main() {
	// 启动一个新的goroutine执行函数 spinner
	go spinner(100 * time.Millisecond)

	const n = 45
	fibN := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	// 这是一个无线循环，但是当main goroutine退出后
	// 这个 程序所在的goroutine会被打断，程序也就退出了
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}

	return fib(x-1) + fib(x-2)
}
