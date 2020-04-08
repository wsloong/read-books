// 可变参数

package main

import "fmt"

func sum(vals ...int) int {
	fmt.Printf("values in sum function: %p\n", vals) // values in sum function: 0xc0000a6040

	var total int
	for _, val := range vals {
		total += val
	}
	return total
}

func main() {
	// total := sum(1, 3, 5)
	// fmt.Printf("1+3+5=%d\n", total)

	valus := []int{2, 4, 6, 8}
	fmt.Printf("values in main function: %p\n", valus) // values in main function: 0xc0000a6040

	total1 := sum(valus...)
	fmt.Printf("total1:%d\n", total1) // total1:20

}
