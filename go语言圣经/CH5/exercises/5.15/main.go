// 练习5.15： 编写类似sum的可变参数函数max和min。考虑不传参时，max和min该如何处理，再编写至少接收1个参数的版本
package main

import "fmt"

func max(vals ...int) (int, error) {
	if len(vals) == 0 {
		return -1, fmt.Errorf("parmars can not empty")
	}

	base := vals[0]
	for _, val := range vals {
		if val > base {
			base = val
			continue
		}
	}
	return base, nil
}

func min(vals ...int) int {
	if len(vals) == 0 {
		panic("parmars can not empty")
	}

	base := vals[0]
	for _, val := range vals {
		if val < base {
			base = val
			continue
		}
	}
	return base
}

func main() {
	max(1)
	min(1)
}
