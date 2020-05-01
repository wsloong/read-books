// 演示一个倒计时程序
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Commencing countdown.")
	// Tick返回一个无缓冲的通道
	// 程序会周期性的向这个通道发送值
	tick := time.Tick(time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		<- tick
	}
	launch()
}

func launch() {

}