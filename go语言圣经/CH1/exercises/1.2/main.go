// 修改echo程序，使其打印value和index，每个value和index显示一行
package main

import "os"

import "fmt"

func main() {
	for index, value := range os.Args {
		fmt.Println(index, value)
	}
}
