// 通过嵌入结构体来扩展类型

package main

import (
	"fmt"
	"image/color"
	"math"
)

type Point struct{ X, Y float64 }

func (p *Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

type ColoredPoint struct {
	Point
	Color color.RGBA
}

func main() {
	var cp ColoredPoint
	// 可以认为嵌入字段就是ColoredPoint自身的字段
	cp.X = 1
	cp.Point.Y = 2
	fmt.Println(cp.Point.X, cp.Y) // 1 2

	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = ColoredPoint{Point{1, 1}, red}
	var q = ColoredPoint{Point{5, 4}, blue}
	fmt.Println(p.Distance(q.Point)) // 5
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.Point)) // 10

	// 方法值和方法表达式
	p1 := Point{1, 2}
	q1 := Point{4, 6}
	distanceFromP := p1.Distance
	fmt.Println(distanceFromP(q1)) // 5

	var origin Point                   // {0, 0}
	fmt.Println(distanceFromP(origin)) // 2.23606797749979

	scale := (*Point).ScaleBy
	scale(&p1, 2)
	fmt.Println(p1)           // {2 2}
	fmt.Printf("%T\n", scale) // func(*main.Point, float64)
}
