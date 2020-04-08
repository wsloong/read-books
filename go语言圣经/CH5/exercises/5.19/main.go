// 练习5.19： 使用panic和recover编写一个不包含return语句但能返回一个非零值的函数。

package main

import "fmt"

func main() {
	fmt.Println(noReturn())
}

func noReturn() (err error) {
	type bailout struct{}

	defer func() {
		switch p := recover(); p {
		case nil:
		// no panic
		case bailout{}:
			// "expected" panic
			err = fmt.Errorf("get a bailout panic")
		default:

		}

	}()
	panic(bailout{})
}
