// 练习 8.5： 使用一个已有的CPU绑定的顺序程序，
// 比如在3.3节中我们写的Mandelbrot程序或者3.2节中的3-D surface计算程序，并将他们的主循环改为并发形式，
// 使用channel来进行通信。在多核计算机上这个程序得到了多少速度上的改进？使用多少个goroutine是最合适的呢？
package main

import (
	"image"
	"image/color"
	"math/cmplx"
	"sync"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
)

func ManderbrotWithGoroutine() {
	type pointColor struct {
		x, y  int
		color color.Color
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	var wg sync.WaitGroup
	ch := make(chan pointColor)

	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			wg.Add(1)
			go func(px, py int) {
				defer wg.Done()
				var p pointColor
				p.x = px
				p.y = py
				p.color = mandelbrot(complex(x, y))
				ch <- p

			}(px, py)
		}
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for p := range ch {
		img.Set(p.x, p.y, p.color)
	}

	// 这里为了性能测试，将其注释
	//png.Encode(os.Stdout, img)
}

func ManderbrotSingle() {
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

	// 这里为了性能测试，将其注释
	// png.Encode(os.Stdout, img)
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
