package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	var w Wheel
	// 下面是两种通过点操作赋值匿名字段
	w.X = 8      // 等价于 w.Circle.Point.X = 8
	w.Y = 8      // 等价于 w.Circle.Point.Y = 8
	w.Radius = 5 // 等价于 w.Circle.Radius = 5
	w.Spokes = 20

	/*

		结构体字面值并没有简短表示匿名成员的语法， 因此下面的语句都不能编译通过：
		w = Wheel{8, 8, 5, 20}                       // compile error: unknown fields
		w = Wheel{X: 8, Y: 8, Radius: 5, Spokes: 20} // compile error: unknown fields
		结构体字面值必须遵循形状类型声明时的结构，所以我们只能用下面的两种语法，它们彼此是等价的：
	*/

	w1 := Wheel{Circle{Point{8, 8}, 5}, 20}
	w2 := Wheel{
		Circle: Circle{
			Point:  Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20,
	}

	fmt.Printf("%#v\n", w1)

	w2.X = 42
	fmt.Printf("%#v\n", w2)
}
