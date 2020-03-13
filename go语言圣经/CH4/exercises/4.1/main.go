// 练习 4.1： 编写一个函数，计算两个SHA256哈希码中不同bit的数目。（参考2.6.2节的PopCount函数。)

package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	fmt.Println(CountDiffBit(c1, c2))
}

func CountDiffBit(c1, c2 [32]byte) int {
	var total int

	// 循环字节数组
	for i := 0; i < len(c1); i++ {
		// 1byte = 8bit
		for j := 1; j <= 8; j++ {
			if (c1[i] >> j) != (c2[i] >> j) {
				total++
			}
		}
	}

	return total
}
