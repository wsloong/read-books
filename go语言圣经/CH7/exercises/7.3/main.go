// 练习 7.3： 为在gopl.io/ch4/treesort (§4.4)的*tree类型实现一个String方法去展示tree类型的值序列

// s结构体可以包含*s类型的成员，下面演示了二叉树的排序

package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

// Sort 对值进行排序
func Sort(values []int) (*tree, []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	return root, appendValues(values[:0], root)
}

// appendValues 将元素按照顺序增加到t中
// 返回 slice
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}

	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}

	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func (t *tree) String() string {
	var res string
	if t == nil {
		return res
	}

	res += t.left.String()
	res += fmt.Sprintf("%d ", t.value)
	res += t.right.String()
	return res
}

func main() {
	unsortValues := []int{
		1, 9, 7, 3, 5, 23,
	}
	t, sortValues := Sort(unsortValues)
	fmt.Println(t)
	fmt.Println(sortValues)
}
