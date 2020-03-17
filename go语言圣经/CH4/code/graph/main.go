package main

import "fmt"

var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

// 即使from到to的边不存在，graph[from][to]依然可以返回一个有意义的结果。
func hasEdge(from, to string) bool {
	return graph[from][to]
}

func main() {
	addEdge("west", "east")
	fmt.Println(hasEdge("north", "south"))
	fmt.Println(hasEdge("west", "east"))
}
