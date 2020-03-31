// 修改gopl.io/ch3/surface (§3.2) 中的corner函数，将返回值命名，并使用bare return。

package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // 画布大小（以像素为单位）
	cells         = 100                 // 网格单元数
	xyrange       = 30.0                // xy的轴范围（-30 30）
	xyscale       = width / 2 / xyrange // xy轴像素单位
	zscale        = height * 0.4        // z轴像素单位
	angle         = math.Pi / 6         // x,y的角度（=30°）
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (sx, sy float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// 计算表面高度
	z := f(x, y)

	// 将（x，y，z）等距投影到二维SVG画布（sx，sy）上
	sx = width/2 + (x-y)*cos30*xyscale
	sy = height/2 + (x+y)*sin30*xyscale - z*zscale
	return
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // 与(0, 0)的距离
	return math.Sin(r) / r
}
