// 练习 3.12： 编写一个函数，判断两个字符串是否是是相互打乱的，也就是说它们有着相同的字符，但是对应不同的顺序。
// TODO: 考虑用map记录各个字符和对应的位置，但是到这里还没有讲到map这个类型
package main

import (
	"strings"
)

func IsOutOrder(s1, s2 string) bool {
	if s1 == s2 {
		return false
	}

	if len(s1) != len(s2) {
		return false
	}

	// 中文情况，range会隐式进行解码
	for _, v := range s1 {
		if !strings.Contains(s2, string(v)) {
			return false
		}
	}

	return true
}
