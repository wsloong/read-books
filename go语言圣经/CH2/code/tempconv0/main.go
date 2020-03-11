// 摄氏温度和华氏温度的计算
package main

import "fmt"

// Celsius和Fahrenheit虽然具有相同的底层类型float64
// 但是它们是不同的数据类型
type Celsius float64    // 摄氏温度
type Fahrenheit float64 // 华氏温度

// 为类型Celsius 定义一个String()方法，
func (c Celsius) String() string {
	return fmt.Sprintf("%g° C", c)
}

const (
	AbsoluteZeroC Celsius = -273.15 // 绝对零度
	FreezingC     Celsius = 0       // 结冰点温度
	BiilingC      Celsius = 100     // 沸点温度
)

// Celsius(t)和Fahrenheit(t)是类型转换操作，它们并不是函数调用。
// 类型转换不会改变值本身，但是会使它们的语义发生变化
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
