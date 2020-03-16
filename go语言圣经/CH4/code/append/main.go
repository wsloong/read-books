// 理解slice的append底层是如何工作
package main

import "fmt"

func appendInt(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)

	if zlen <= cap(x) {
		// 有足够的空间,原底层数组基础上扩展slice
		z = x[:zlen]
	} else {
		// 空间不足，2倍增长空间
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		// 分配一个新切片，该切片和原切片有不同的底层数组
		z = make([]int, zlen, zcap)
		// copy原数据到新切片
		copy(z, x)
	}

	copy(z[len(x):], y)
	return z
}

func main() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
}
