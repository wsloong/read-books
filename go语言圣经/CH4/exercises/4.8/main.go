// 练习 4.8： 修改charcount程序，使用unicode.IsLetter等相关的函数，统计字母、数字等Unicode中不同的字符类别。

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	letters := make(map[rune]int)
	numbers := make(map[rune]int)
	others := make(map[rune]int)
	var invalid int

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}

		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		if unicode.IsLetter(r) {
			letters[r]++
			continue
		}

		if unicode.IsNumber(r) {
			numbers[r]++
			continue
		}

		others[r]++
	}

	fmt.Printf("letters\tcount\n")
	for c, n := range letters {
		fmt.Printf("%q\t%d\n", c, n)
	}

	fmt.Print("\numbers\tcount\n")
	for i, n := range numbers {
		if i > 0 {
			fmt.Printf("%q\t%d\n", i, n)
		}
	}

	fmt.Print("\nothers\tcount\n")
	for i, n := range others {
		if i > 0 {
			fmt.Printf("%q\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
