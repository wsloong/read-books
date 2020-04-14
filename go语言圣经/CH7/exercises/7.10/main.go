// 练习 7.10： sort.Interface类型也可以适用在其它地方。
// 编写一个IsPalindrome(s sort.Interface) bool函数表明序列s是否是回文序列，
// 换句话说反向排序不会改变这个序列。假设如果!s.Less(i, j) && !s.Less(j, i)则索引i和j上的元素相等。

package main

import (
	"fmt"
	"sort"
)

func IsPalindrome(s sort.Interface) bool {
	if s.Len() == 0 {
		return false
	}

	i, j := 0, s.Len()-1
	for i < j {
		if !s.Less(i, j) && !s.Less(j, i) {
			i++
			j--
			continue
		}
		return false
	}

	// halfLen := s.Len() / 2

	// for i := 0; i < halfLen; i++ {
	// 	before := i
	// 	end := s.Len() - 1 - i
	// 	if !s.Less(before, end) && !s.Less(end, before) {
	// 		continue
	// 	}

	// 	return false
	// }

	return true
}

func main() {
	a := []int{1, 2, 3, 2, 1}
	fmt.Println(IsPalindrome(sort.IntSlice(a))) // true

	a = []int{1, 2, 3, 4, 5}
	fmt.Println(IsPalindrome(sort.IntSlice(a))) //false

	a = []int{1}
	fmt.Println(IsPalindrome(sort.IntSlice(a))) //false

	a = []int{}
	fmt.Println(IsPalindrome(sort.IntSlice(a))) // false

}
