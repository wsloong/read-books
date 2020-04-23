// 使用2个channel将3个goroutine链接起来
// 1, 2, 3 ... => 1, 4, 9...
package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for x := 0; ; x++ {
			naturals <- x
		}
	}()

	go func() {
		for { // 我第一遍默写时候for循环忘记了，导致deadlock
			x := <-naturals
			squares <- x * x
		}
	}()

	for {
		fmt.Println(<-squares)
	}
}
