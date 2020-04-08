// 练习5.18：不修改fetch的行为，重写fetch函数，要求使用defer机制关闭文件。

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

const URL = "https://golang.google.cn/"

func main() {
	filename, n, err := fetch(URL)
	if err != nil {
		panic(err)
	}
	fmt.Printf("http %s; save to %s, n=%d\n", URL, filename, n)
}

// Fetch 下载URL内容到文件，并返回文件名称和长度
func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}

	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}

	defer func() {
		if closeErr := f.Close(); err == nil {
			err = closeErr
		}
	}()

	n, err = io.Copy(f, resp.Body)
	return local, n, err
}
