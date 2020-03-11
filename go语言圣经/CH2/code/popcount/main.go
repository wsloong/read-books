// 计算一个无符号整数的二进制中1的个数
package main

import "fmt"

/*
生成辅助数组的代码可以用匿名函数来封装
var pc [256]byte = func() (pc [256]byte) {
	for i := range pc {
		pc[i] = pc[i+2] + byte(i&1)
	}
}()

初始化构造一个长度256的数组
为什么是256呢？
结合主程序PopCount，将64为的无符号整数切割成8个8位，而8位的无符号整数取值范围是0-255[2^8 -1]
*/
var pc [256]byte

// 生成辅助表格pc
/*
	25|2		1
	12|2		0
	6|2			0
	3|1			1
	1|2			1
*/
func init() {
	for i := range pc {
		// 得出整数中二进制`1的个数`字典表
		// 整数i的1的个数 = 整数i/2中1的个数 + i在低1位的1的个数
		pc[i] = pc[i/2] + byte(i&1)
	}
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

// 练习 2.3：重写PopCount函数，用一个循环代替单一的表达式
func PopCountWithLoop(x uint64) int {
	var total int
	for i := 0; i < 8; i++ {
		total += int(pc[byte(x>>(i*8))])
	}
	return total
}

// 练习 2.4： 用移位算法重写PopCount函数，每次测试最右边的1bit，然后统计总数。
func PopCountWithBitMove(x uint64) int {
	var total int
	for i := 0; i < 64; i++ {
		// &1；只有最右边第一位是1是结果才是1
		if x&1 == 1 {
			total++
		}

		// 右移以为
		x = x >> 1
	}
	return total
}

// 练习 2.5： 表达式x&(x-1)用于将x的最低的一个非零的bit位清零。
func PopCountWithCleanLowest(x uint64) int {
	// 3	0011
	// 3-1	0010
	// &	0010
	// 2-1  0001
	// &	0000
	var total int
	for x != 0 {
		x = x & (x - 1)
		total++
	}
	return total
}

func main() {
	fmt.Println(PopCount(10))
	fmt.Println(PopCountWithLoop(10))
	fmt.Println(PopCountWithBitMove(10))
	fmt.Println(PopCountWithCleanLowest(10))
}
