// 输出曼德布罗特图片

/*
用于遍历1024x1024图像每个点的两个嵌套的循环对应-2到+2区间的复数平面。
程序反复测试每个点对应复数值平方值加一个增量值对应的点是否超出半径为2的圆。
如果超过了，通过根据预设置的逃逸迭代次数对应的灰度颜色来代替。
如果不是，那么该点属于Mandelbrot集合，使用黑色颜色标记。
最终程序将生成的PNG格式分形图像图像输出到标准输出，
*/
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)

			// 图像点（px，py）表示复数值z。
			img.Set(px, py, mandelbrot(z))
		}
	}

	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
