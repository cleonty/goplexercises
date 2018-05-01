package main

import (
	"fmt"
	"net/http"
)

func main() {
	db := database{"shoes": 50, "socks": 30}
	http.ListenAndServe(":8080", db)
}

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	case "/price":
		item := r.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "нет товара %q\n", item)
			return
		}
		fmt.Fprintf(w, "цена %q: %s\n", item, price)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "нет страницы: %s\n", r.URL.Path)
	}
}
