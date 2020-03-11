// 练习 3.6： 升采样技术可以降低每个像素对计算颜色值和平均值的影响。简单的方法是将每个像素分成四个子像素，实现它。
// TODO:说实话，我没懂，参考了https://blog.csdn.net/q1576962841/article/details/86084461这篇文章

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
		epsX                   = (xmax - xmin) / width
		epsY                   = (ymax - ymin) / height
	)

	offX := []float64{-epsX, epsX}
	offY := []float64{-epsY, epsY}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			subPixels := make([]color.Color, 0)

			for i := 0; i < 2; i++ {
				for j := 0; j < 2; j++ {
					z := complex(x+offX[i], y+offY[j])
					subPixels = append(subPixels, mandelbrot(z))
				}
			}
			img.Set(px, py, avg(subPixels))
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
			// TODO 这里可以组合各种色值
			return color.RGBA{iterations - n, iterations / (n + 1), n, 255}
		}
	}
	return color.Black
}

// 4组颜色拆分r、g、b、a，算出平均值
func avg(colors []color.Color) color.Color {
	var r, g, b, a uint16
	n := len(colors)
	for _, c := range colors {
		_r, _g, _b, _a := c.RGBA()
		r += uint16(_r / uint32(n))
		g += uint16(_g / uint32(n))
		b += uint16(_b / uint32(n))
		a += uint16(_a / uint32(n))
	}
	return color.RGBA64{r, g, b, a}
}
