// 非负十进制整数字符串中插入逗号。"12345"  => "12,345"
package main

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	return comma(s[:n-3]) + "," + s[n-3:]
}
