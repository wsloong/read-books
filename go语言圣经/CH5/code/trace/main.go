// defer 机制用于记录程序的进入和退出时间

package main

import (
	"log"
	"time"
)

func main() {
	bigSlowOperation()
}

func bigSlowOperation() {
	// trace返回一个函数，该函数会在bigSlowOperation退出时候调用；注意后面的`()`
	defer trace("bigSlowOperation")()

	// ....do works
	time.Sleep(10 * time.Second)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s\n", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}
