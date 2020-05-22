package charcount

import (
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

func CharCount(input string) {
	// 定义一个map，用于统计Unicode码点的数量
	counts := make(map[rune]int)
	// 定义一个数组，统计UTF8编码的数量
	var utflen [utf8.UTFMax + 1]int
	// 统计UTF8编码中错误的数量
	var invalid int
	in := strings.NewReader(input)
	for {
		// 执行UTF-8解码并返回三个值：
		// 1， 解码的rune字符的值
		// 2， 字符UTF-8编码后的长度
		// 3， 一个错误值
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}

		// 如果输入的是无效的UTF-8编码的字符，
		// 返回的将是unicode.ReplacementChar表示无效字符，并且编码长度是1。
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}

	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}

	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
