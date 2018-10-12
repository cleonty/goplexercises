package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cleonty/gopl/ch12/params"
)

func search(resp http.ResponseWriter, req *http.Request) {
	var data struct {
		Labels     []string `http:"l"`
		MaxResults int      `http:"max"`
		Exact      bool     `http:"x"`
		Email      string   `http:"email" validator:"email"`
	}
	data.MaxResults = 10
	if err := params.Unpack(req, &data); err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(resp, "Search: %+v\n", data)
}

func main() {
	log.Printf("starting...")
	http.HandleFunc("/", search)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
