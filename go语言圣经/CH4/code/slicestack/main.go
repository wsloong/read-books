// 使用[]int类型的 slice模拟stack数据结构

package main

var stack []int

// 压入栈
func push(value int) {
	stack = append(stack, value)
}

// 弹出栈顶的元素
func pop() int {
	top := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	return top
}

// 删除指定索引的元素
func remove(index int) []int {
	copy(stack[index:], stack[index+1:])
	return stack
}
