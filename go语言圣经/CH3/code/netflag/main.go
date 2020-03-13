// 演示iota的使用

package main

import "fmt"

type Flags uint

const (
	FlagUp           Flags = 1 << iota // 1 is up
	FlagBroadcast                      // 10 supports broadcast access capability
	FlagLoopback                       // 100 is a loopback interface
	FlagPointToPoint                   // 1000 belongs to a point-to-point link
	FloagMulticast                     // 10000 supports multicast access capability
)

func IsUp(v Flags) bool     { return v&FlagUp == FlagUp }
func TurnDown(v *Flags)     { *v &^= FlagUp }
func SetBroadcast(v *Flags) { *v |= FlagBroadcast }
func IsCast(v Flags) bool   { return v&(FlagBroadcast|FloagMulticast) != 0 }

func main() {
	var v Flags = FloagMulticast | FlagUp // 10001
	fmt.Printf("%b %t\n", v, IsUp(v))     // 10001 true
	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v)) // 10000 false
	SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))   // 10010 false
	fmt.Printf("%b %t\n", v, IsCast(v)) // 10010 true

}
