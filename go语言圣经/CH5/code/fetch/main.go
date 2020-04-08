// 1.5节fetch的改动版，将http响应信息写入本地文件而不是从标准输出流输出
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

	n, err = io.Copy(f, resp.Body)
	// Close file, but prefer error from Co[]py, if any
	// NFS文件系统，写入发生的错误会延迟到关闭才反馈， 所以这里检查关闭是否有错误
	// 如果io.Copy和f.Close都发生错误了，我们倾向于将io.Copy错误反馈给调用者
	// 因为它先于f.Close发生，更接近问题的本质
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	return local, n, err
}
