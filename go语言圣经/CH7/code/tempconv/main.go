/*
package flag

// Value is the interface to the value stored in a flag.
type Value interface {
	String() string
	Set(string) error
}
该示例实现flag.Value方法的接口
*/

package main

import (
	"flag"
	"fmt"
)

type Celsius float64
type Fahrenheit float64
type Kelvin float64

func (c Celsius) String() string {
	return fmt.Sprintf("%g° C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g° F", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%g° K", k)
}

//  FToC  将华氏温度转换为摄氏温度
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}


// KToC 开尔文温度转换为摄氏温度
func KToC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}


type celsiusFlag struct{ Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)

	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	case "K", "°K":
		f.Celsius = KToC(Kelvin(value))
		return nil
	}

	return fmt.Errorf("invalid temperature %q", s)
}

// CelsiusFlag defines a Celsius flag with the specified name,
// default value, and usage, and returns the address of the flag variable.
// The flag argument must have a quantity and a unit, e.g., "100C".
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

// 使用新标记
var temp = CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}

/*
练习 7.7： 解释为什么帮助信息在它的默认值是20.0没有包含°C的情况下输出了°C。
var temp = CelsiusFlag("temp", 20.0, "the temperature") 该语句已经调用了CelsiusFlag()方法
Celsius=value(这里是20.0)，所以默认类型是`Celsius`。输出也是`Celsius.String()`
*/
