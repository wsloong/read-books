// 上手实践前面提到的strings.Join和直接Println，并观察输出结果的区别

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args, " "))
	fmt.Println(os.Args)
}
