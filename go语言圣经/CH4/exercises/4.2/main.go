// 练习 4.2： 编写一个程序，默认情况下打印标准输入的SHA256编码，并支持通过命令行flag定制，输出SHA384或SHA512哈希算法。
package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	d := flag.String("d", "sha256", "哈希算法:sha256(默认)、sha384、sha512")
	str := flag.String("s", "x", "请输入加密字符串")
	flag.Parse()

	flag := strings.ToLower(*d)

	switch flag {
	case "sha512":
		fmt.Fprintf(os.Stdout, "%x\n", sha512.Sum512([]byte(*str)))
	case "sha384":
		fmt.Fprintf(os.Stdout, "%x\n", sha512.Sum384([]byte(*str)))
	default:
		fmt.Fprintf(os.Stdout, "%x\n", sha256.Sum256([]byte(*str)))
	}
}
