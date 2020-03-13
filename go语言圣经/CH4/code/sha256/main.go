// crypto/sha256包的Sum256函数对一个任意的字节slice类型的数据生成一个对应的消息摘要。
// 消息摘要有256bit大小，因此对应[32]byte数组类型。
// 如果两个消息摘要是相同的，那么可以认为两个消息本身也是相同
package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)

	//输出
	// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
	// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	// false
	// [32]uint8

	// %x 十六进制格式打印
	// %t 打印布尔类型
	// %T 显示一个值对应的数据类型
}
