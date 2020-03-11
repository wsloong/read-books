// Echo4 打印命令行参数
package main

import (
	"flag"
	"fmt"
	"strings"
)

// 创建一个对应布尔型标志参数的变量n，默认false，说明`omit trailing newling`
var n = flag.Bool("n", false, "omit trailing newling")

// 创建一个对应字符串型标志参数的变量sep，默认" "，说明`separator`
var sep = flag.String("s", " ", "separator")

func main() {
	//更新每个标志参数对应变量的值(之前是默认值).
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
