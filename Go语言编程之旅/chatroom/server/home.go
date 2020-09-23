package server

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/wsloong/chatroom/global"
)

func homeHandleFunc(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles(global.RootDir + "/template/home.html")
	if err != nil {
		fmt.Fprint(w, "模板解析错误")
		return
	}

	if err := tpl.Execute(w, nil); err != nil {
		fmt.Fprintf(w, "模板执行错误")
		return
	}
}
