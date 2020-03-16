// 练习 4.5： 写一个函数在原地完成消除[]string中相邻重复的字符串的操作。
package main

import "fmt"

/*
原地完成消除
遇到相同的先前移一位
下标保持不动，继续检测当前位置是否和下一个元素重复
*/
func removeDuplicates(str []string) []string {
	for i := 0; i < len(str)-1; i++ {
		if str[i] == str[i+1] {
			copy(str[i:], str[i+1:])
			str = str[:len(str)-1]
			i--
		}
	}
	return str
}

func main() {
	s := []string{"a", "b", "b", "b", "a", "c"}
	fmt.Println(s)
	fmt.Println(removeDuplicates(s))

	// copy的规则
	// 1， src依次替换dst相对应元素
	// 2， 按其中较小的那个数组切片的元素个数进行复制
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{5, 4, 3}
	copy(s1, s2) // 5, 4, 3, 4, 5; 5, 4, 3
	fmt.Println(s1, s2)
	copy(s2, s1) // 5, 4, 3, 4, 5; 5, 4, 3
	fmt.Println(s1, s2)
}
