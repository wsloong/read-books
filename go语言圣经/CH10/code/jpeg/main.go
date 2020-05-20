// 演示匿名导入
// 读取一个PNG图片，将至转换为JPGE图片

// 使用
// go build -o mandelbrot ../CH3/code/mandelbrot
// go build -o jpeg ./code/jpeg
// ./mandelbrot| ./jpeg > mandelbrot.jpg

package main

import (
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png" // register PNG decoder, 如果没有会报错"jpeg: image: unknown format"
	"io"
	"os"
)

func main() {
	if err := toJPEG(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
		os.Exit(1)
	}
}

func toJPEG(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format = ", kind)
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}
