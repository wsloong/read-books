// 练习 7.11： 增加额外的handler让客户端可以创建，读取，更新和删除数据库记录。
// 例如，一个形如 /update?item=socks&price=6 的请求会更新库存清单里一个货品的价格
// 并且当这个货品不存在或价格无效时返回一个错误值。（注意：这个修改会引入变量同时更新的问题）
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func main() {
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
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
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
