package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func fetch(url string) (body string) {
	var b bytes.Buffer
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return ""
	}
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	_, err = io.Copy(&b, resp.Body)
	return b.String()
}

func mirroredQuery() string {
	responses := make(chan string, 3)
	go func() { responses <- fetch("asia.gopl.io") }()
	go func() { responses <- fetch("europe.gopl.io") }()
	go func() { responses <- fetch("americas.gopl.io") }()
}

func main() {
	for _, url := range os.Args[1:] {

	}
}
