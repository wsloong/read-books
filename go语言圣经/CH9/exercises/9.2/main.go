// 练习 9.2： 重写2.6.2节中的PopCount的例子，使用sync.Once，只在第一次需要用到的时候进行初始化。
// (虽然实际上，对PopCount这样很小且高度优化的函数进行同步可能代价没法接受)
package main

import (
	"fmt"
	"sync"
)

var once sync.Once
var pc [256]byte

func init() {
	once.Do(func() {
		for i := range pc {
			// 得出整数中二进制`1的个数`字典表
			// 整数i的1的个数 = 整数i/2中1的个数 + i在低1位的1的个数
			pc[i] = pc[i/2] + byte(i&1)
		}
	})
}

// 通过8次右移操作分别求低8位的1数量，然后相加
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func main() {
	fmt.Println(PopCount(10))
}
