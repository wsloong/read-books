// 练习 7.9： 使用html/template包 (§4.6) 替代printTracks将tracks展示成一个HTML表格。
// 将这个解决方案用在前一个练习中，让每次点击一个列的头部产生一个HTTP请求来排序这个表格

package main

import (
	"html/template"
	"log"
	"net/http"
	"sort"
	"time"
)

const temp2 = `
<table>
<tr style='text-align: left'>
<th><a href='/?sort_by=title'>Title</a></th>
<th><a href='/?sort_by=artist'>Artist</a></th>
<th><a href='/?sort_by=album'>Album</a></th>
<th><a href='/?sort_by=year'>Year</a></th>
<th><a href='/?sort_by=length'>Length</a></th>
</tr>
{{range .T}}
<tr>
<td>{{.Title}}</td>
<td>{{.Artist}}</td>
<td>{{.Album}}</td>
<td>{{.Year}}</td>
<td>{{.Length}}</td>
</tr>
{{end}}
</table>`

var showHtml = template.Must(template.New("sorting").Parse(temp2))

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("4m24s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("3m37s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

type lessFunc func(i, j int) bool

type tableSort struct {
	T         []*Track
	lessFuncs []lessFunc
}

var ts = tableSort{T: tracks}

var columnFuns = map[string]lessFunc{
	"title":  ts.byTitle,
	"artist": ts.byArtist,
	"album":  ts.byAlbum,
	"year":   ts.byYear,
	"length": ts.byLength,
}

func (x tableSort) Len() int { return len(x.T) }
func (x tableSort) Less(i, j int) bool {
	for k := len(x.lessFuncs) - 1; k >= 0; k-- {
		// 如果最后一个直接可以确定
		if x.lessFuncs[k](i, j) {
			return true
		} else if !x.lessFuncs[k](j, i) { // 判断是否相等，相等就执行下一个less函数
			continue
		}

		return false
	}
	return false
}

func (x tableSort) Swap(i, j int) { x.T[i], x.T[j] = x.T[j], x.T[i] }

func (x tableSort) byTitle(i, j int) bool {
	return x.T[i].Title < x.T[j].Title
}

func (x tableSort) byArtist(i, j int) bool {
	return x.T[i].Artist < x.T[j].Artist
}

func (x tableSort) byAlbum(i, j int) bool {
	return x.T[i].Album < x.T[j].Album
}

func (x tableSort) byYear(i, j int) bool {
	return x.T[i].Year < x.T[j].Year
}

func (x tableSort) byLength(i, j int) bool {
	return x.T[i].Length < x.T[j].Length
}

func main() {
	http.HandleFunc("/", IndexHandler)
	log.Fatal(http.ListenAndServe(":8090", nil))
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	p := q.Get("sort_by")

	f, ok := columnFuns[p]
	if !ok {
		f = ts.byTitle
	}
	// TODO随着次数增加，单纯的往里面append这种方式并不是最好的，这里只是演示
	ts.lessFuncs = append(ts.lessFuncs, f)
	sort.Sort(ts)

	if err := showHtml.Execute(w, ts); err != nil {
		log.Fatal(err)
	}
}
