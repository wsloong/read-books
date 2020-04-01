// 练习5.11： 现在线性代数的老师把微积分设为了前置课程。完善topSort，使其能检测有向图中的环

package main

import (
	"fmt"
	"os"
	"sort"
)

// prereqs记录了每个课程的前置课程
var prereqs = map[string][]string{
	"algorithms":            {"data structures"},
	"calculus":              {"linear algebra"},
	"compilers":             {"data structures", "formal languages", "computer organization"},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
	// // 线性代数的前置课程为微积分。这样形成了一个环
	"linear algebra": {"calculus"},
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(key string, items []string)
	visitAll = func(key string, items []string) {
		for _, item := range items {
			for _, item2 := range m[item] {
				if key == item2 {
					fmt.Println("cycle")
					os.Exit(1)
				}
			}

			if !seen[item] {
				seen[item] = true
				visitAll(item, m[item])
				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	// [algorithms calculus compilers data structures databases discrete math formal languages networks operating systems programming languages]
	sort.Strings(keys)
	visitAll(keys[0], keys)
	return order
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}
