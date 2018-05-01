package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func main() {
	db := database{}
	db.Items = map[string]dollars{"shoes": 50, "socks": 30}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/update", db.update)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database struct {
	sync.Mutex
	Items map[string]dollars
}

func (db *database) list(w http.ResponseWriter, r *http.Request) {
	db.Lock()
	defer db.Unlock()
	if err := dbTemplate.Execute(w, db); err != nil {
		fmt.Fprint(w, err)
		return
	}
}

func (db *database) price(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	db.Lock()
	price, ok := db.Items[item]
	db.Unlock()
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "нет товара %q\n", item)
		return
	}
	fmt.Fprintf(w, "цена %q: %s\n", item, price)
}

func (db *database) update(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price, err := strconv.Atoi(r.URL.Query().Get("price"))
	if err != nil {
		fmt.Fprintf(w, "ошибка  в цене %s\n", r.URL.Query().Get("price"))
		return
	}
	db.Lock()
	defer db.Unlock()
	_, ok := db.Items[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "нет товара %q\n", item)
		return
	}
	db.Items[item] = dollars(price)
	fmt.Fprintf(w, "новая цена для %q: %s\n", item, dollars(price))
}

var dbTemplate = template.Must(template.New("db").Parse(`<h1>Магазин</h1>
	<table>
	  <tr>
		<th>Товар</th>
		<th>Цена</th>
	  </tr>
	  {{range $key, $value := .Items}}
	  <tr>
	  	<td>{{$key}}</td>
	  	<td>{{$value}}</td>
	  </tr>
	  {{end}}
	</table>
	`))
