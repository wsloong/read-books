// reverse 原内存空间翻转slice
package main

import "fmt"

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	s := [...]int{0, 1, 2, 3, 4, 5, 6}
	reverse(s[:])
	fmt.Println(s)
	fmt.Println()

	// 将slice元素循环向左旋转n个元素的方法就是三次调用revers翻转函数
	// 第一次反转开头的n个元素，然后反转剩下的元素，最后反转整个slice元素
	// 如果是向右旋转，则将第三个函数调用移到第一个调用位置就好
	// 下面演示 向左旋转2个元素
	s1 := []int{0, 1, 2, 3, 4, 5}
	reverse(s1[:2])
	reverse(s1[2:])
	reverse(s1)
	fmt.Println(s1) // [2 3 4 5 0 1]
}
