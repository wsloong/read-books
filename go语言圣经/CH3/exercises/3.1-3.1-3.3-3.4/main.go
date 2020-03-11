// 使用 SVG(可缩放的矢量图，矢量线绘制的XML标准) 渲染 3D 表面图

package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
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
type typeFunc func(x, y float64) (float64, bool)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/egg", eggHandler)
	http.HandleFunc("/saddle", saddleHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	surface(w, f)
}

func eggHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	surface(w, eggbox)
}

func saddleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	surface(w, saddle)
}

func surface(w io.Writer, f typeFunc) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, isNaN := corner(i+1, j, f)
			if isNaN {
				continue
			}
			bx, by, isNaN := corner(i, j, f)
			if isNaN {
				continue
			}
			cx, cy, isNaN := corner(i, j+1, f)
			if isNaN {
				continue
			}
			dx, dy, isNaN := corner(i+1, j+1, f)
			if isNaN {
				continue
			}
			fmt.Fprintf(w, "<polygon style='stroke: %s;'  points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				color(i, j), ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprint(w, "</svg>")
}

func corner(i, j int, f typeFunc) (float64, float64, bool) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// 计算表面高度
	z, isNaN := f(x, y)
	if isNaN {
		return 0, 0, isNaN
	}

	// 将（x，y，z）等距投影到二维SVG画布（sx，sy）上
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, isNaN
}

// 练习 3.1： 如果f函数返回的是无限制的float64值，
//那么SVG文件可能输出无效的多边形元素（虽然许多SVG渲染器会妥善处理这类问题
// 修改程序跳过无效的多边形。
// 这里增加一个bool值用于确认生成的float64为有效的值
func f(x, y float64) (float64, bool) {
	r := math.Hypot(x, y) // 与(0, 0)的距离
	res := math.Sin(r)
	return res, math.IsNaN(res)
}

// 输出鸡蛋壳
func eggbox(x, y float64) (float64, bool) {
	res := 0.2 * (math.Cos(x) + math.Cos(y))
	return res, math.IsNaN(res)
}

// 输出马鞍
func saddle(x, y float64) (float64, bool) {
	a := 25.0
	b := 17.0
	a2 := a * a
	b2 := b * b
	res := y*y/a2 - x*x/b2
	return res, math.IsNaN(res)
}

// 练习 3.3： 根据高度给每个多边形上色，那样峰值部将是红色(#ff0000)，谷部将是蓝色(#0000ff)
// TOOD: 这里偷懒了，原因是我不会-_-
func color(i, j int) string {
	if (i+j)&1 == 0 {
		return "#ff0000"
	}
	return "#0000ff)"
}
