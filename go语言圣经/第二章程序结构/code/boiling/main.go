// bioling：打印水的沸点

package main

import (
	"fmt"
	"path"
)

const boilingF = 212.0

func main() {
	fname := "xxx.csv"
	fmt.Println(path.Ext(fname))

	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g° F or %g° C \n", f, c)
}
