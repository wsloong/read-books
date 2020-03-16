// 练习 4.3： 重写reverse函数，使用数组指针代替slice
package main

import "fmt"

func reverse(arr *[5]int) {
	// for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
	// 	arr[i], arr[j] = arr[j], arr[i]
	// }

	for i, j := 0, len(arr)-1; i < j; {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4}
	reverse(&arr)
	fmt.Println(arr)
}
