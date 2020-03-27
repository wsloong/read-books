// 演示信任html
// ./autoescape > autoescape.html；打开autoescape.html查看效果
package main

import (
	"html/template"
	"os"
)

func main() {
	const temp1 = `<p>A: {{.A}}</p><p>B:{{.B}}</p>`
	t := template.Must(template.New("escape").Parse(temp1))

	var data struct {
		A string        // 不可信的文本
		B template.HTML // 信任的HTML
	}
	data.A = "<b>Hello!</b>"
	data.B = "<b>Hello!</b>"

	if err := t.Execute(os.Stdout, data); err != nil {
		panic(err)
	}
}
