// 练习 3.13： 编写KB、MB的常量声明，然后扩展到YB。

// 如果常量表达式省略，表示使用前面常量的表达式写法；因为iota又是自增的
// KM = 1 << (10 * 1)
// MB = 1 << (10 * 2)
// GB = 1 << (10 * 3)
// ...
// GB = 1 << (10 * 8)

package main

const (
	_  = iota
	KM = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)
