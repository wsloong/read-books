// 练习 3.10： 编写一个非递归版本的comma函数，使用bytes.Buffer代替字符串链接操作。
// 练习 3.11： 完善comma函数，以支持浮点数处理和一个可选的正负号的处理。

package main

import (
	"bytes"
	"strings"
)

func Comma(s string) string {
	var oldStr = s
	var symbol string
	// 先去掉首位的符号
	if s[0] == '+' || s[0] == '-' {
		symbol = string(s[0])
		s = s[1:]
	}

	dotLastIndex := strings.LastIndex(s, ".")
	var decimal string
	if dotLastIndex > 0 {
		decimal = s[dotLastIndex:]
		s = s[:dotLastIndex]
	}

	n := len(s)
	if n <= 3 {
		return oldStr
	}
	return symbol + Comma(s[:n-3]) + "," + s[n-3:] + decimal
}

func CommaWithoutRecursive(s string) string {
	var oldStr = s
	var symbol string
	if s[0] == '+' || s[0] == '-' {
		symbol = string(s[0])
		s = s[1:]
	}

	// 小数部分的数字
	dotLastIndex := strings.LastIndex(s, ".")
	var decimal string
	if dotLastIndex > 0 {
		decimal = s[dotLastIndex:]
		s = s[:dotLastIndex]
	}

	n := len(s)
	if n <= 3 {
		return oldStr
	}

	var buf bytes.Buffer
	if len(symbol) > 0 {
		buf.WriteString(symbol)
	}

	for i := 0; i < n; i++ {
		buf.WriteByte(s[i])

		// 1%3 = 1, 2%3=2, 3%3=0, 4%3=1....循环0,1,2不会超过3
		if (i+1)%3 == n%3 && i != n-1 {
			buf.WriteByte(',')
		}
	}
	if len(decimal) > 0 {
		buf.WriteString(decimal)
	}
	return buf.String()
}
