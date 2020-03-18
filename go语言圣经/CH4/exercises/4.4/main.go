// 练习 4.4： 编写一个rotate函数，通过一次循环完成旋转。
package main

import "fmt"

/*
s []int 切片
r int	右移几位
借助一个新切片，新切片的下标为原数组下标加上偏移量
如果超过切片的最大长度，则从最左边开始
*/
func rotate(s []int, r int) []int {
	slen := len(s)
	s1 := make([]int, slen)
	for k := range s {
		index := r + k
		if index >= slen {
			index -= slen
		}
		s1[index] = s[k]
	}
	return s1
}

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	s1 := rotate(s, 2)
	fmt.Println(s1)

	type Employee struct {
		Name string
		Age int
	}

	var dilbert Employee
	fmt.Println(dilbert)
	dilbert.Name = "ahah"

	var  employeeOfTheMonth *Employee = &dilbert
	(*employeeOfTheMonth).Age = 10
	fmt.Println(*employeeOfTheMonth)
}
