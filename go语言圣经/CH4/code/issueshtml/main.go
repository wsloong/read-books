// 使用html模板字符串打印bug
package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

const temp2 = `
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
<th>#</th>
<th>State</th>
<th>User</th>
<th>Title</th>
</tr>
{{range .Items}}
<tr>
<td><a href='{{.HTMLURL}}'>{{.Number}}</td>
<td>{{.State}}</td>
<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

var report = template.Must(template.New("issuelist").Parse(temp2))

func main() {
	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	args := query.Get("q")
	if args == "" {
		args = "windows+label:bug"
	}
	
	result, err := SearchIssues(args)
	if err != nil {
		log.Fatal(err)
	}

	if err := report.Execute(w, result); err != nil {
		log.Fatal(err)
	}
}
