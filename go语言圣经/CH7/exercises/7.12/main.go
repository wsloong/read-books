// 练习 7.12： 修改/list的handler让它把输出打印成一个HTML的表格而不是文本。
// html/template包(§4.6)可能会对你有帮助

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
)

const temp2 = `
<table>
<tr style='text-align: left'>
<th>ITEM</th>
<th>PRICE</th>
</tr>
{{range $k, $v := .db}}
<tr>
<td>{{$k}}</td>
<td>{{$v}}</td>
</tr>
{{end}}
</table>`

var showHtml = template.Must(template.New("price").Parse(temp2))

func main() {
	err1 := fmt.Errorf("1")
	err2 := fmt.Errorf("2")
	fmt.Println(err1 == err2)

	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/update", db.update)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

var mu sync.Mutex

type dollars float64

func (d dollars) String() string { return fmt.Sprintf("$%.2f\n", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, r *http.Request) {
	res := map[string]database{"db": db}
	if err := showHtml.Execute(w, res); err != nil {
		log.Fatal(err)
	}
}

func (db database) price(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db database) update(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	item := q.Get("item")

	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}

	price, err := strconv.ParseFloat(q.Get("price"), 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "invalidate price: %s\n", price)
		return
	}

	mu.Lock()
	defer mu.Unlock()
	db[item] = dollars(price)
	fmt.Fprintf(w, "success")
}
