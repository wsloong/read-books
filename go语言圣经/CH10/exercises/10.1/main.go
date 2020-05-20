// 练习 10.1： 扩展jpeg程序，以支持任意图像格式之间的相互转换，
// 使用image.Decode检测支持的格式类型，然后通过flag命令行标志参数选择输出的格式

package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png" // register PNG decoder, 如果没有会报错"jpeg: image: unknown formatf"
	"io"
	"os"

	"golang.org/x/image/bmp"
)

var outputFormat string

func init() {
	flag.StringVar(&outputFormat, "f", "gif", "output format; forexample:png")
}

func main() {
	flag.Parse()

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

	switch outputFormat {
	case "png":
		return png.Encode(out, img)
	case "jpg":
		return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
	case "gif":
		return gif.Encode(out, img, &gif.Options{NumColors: 95})
	case "bmp":
		return bmp.Encode(out, img)
	}
	return fmt.Errorf("unknow format")
}
