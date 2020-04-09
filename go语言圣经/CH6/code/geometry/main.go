// 方法
package geometry

import "math"

type Point struct{ X, Y float64 }

// 传统的函数；geometry包级别的函数
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// 类型 Point 的方法；
// p 方法的接收器
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// 不同类型可以定义相同的方法名。因为每种类型都有其方法的命名空间
type Path []Point

func (path Path) Distance() float64 {
	var sum float64
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}
