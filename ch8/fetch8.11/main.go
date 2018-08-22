package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type Response struct {
	body string
	err  error
}

func fetch(ctx context.Context, url string, responses chan<- Response) {
	var b bytes.Buffer
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Printf("fetch: NewRequest: %v\n", err)
		responses <- Response{"", err}
		return
	}
	request = request.WithContext(ctx)
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Printf("fetch: do: %v\n", err)
		responses <- Response{"", err}
		return
	}
	defer resp.Body.Close()
	_, err = io.Copy(&b, resp.Body)
	responses <- Response{b.String(), nil}
}

func mirroredQuery() string {
	responses := make(chan Response, 3)
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	go func() {
		wg.Add(1)
		defer wg.Done()
		defer fmt.Println("end 1")
		fetch(ctx, "http://google.ru", responses)
	}()
	go func() {
		wg.Add(1)
		defer wg.Done()
		defer fmt.Println("end 2")
		fetch(ctx, "http://google.ru", responses)
	}()
	go func() {
		wg.Add(1)
		defer wg.Done()
		defer fmt.Println("end 3")
		fetch(ctx, "http://google.ru", responses)
	}()
	fmt.Println("waiting for response")
	errCount := 0
	var response Response
	for {
		response = <-responses
		if response.err != nil {
			fmt.Printf("got error: %v\n", response.err)
			errCount++
			if errCount > 2 {
				break
			}
			continue
		}
		fmt.Println("got response!")
		break
	}
	fmt.Println("cancel other go rounites!")
	cancel()
	fmt.Println("waiting for goroutines")
	wg.Wait()
	fmt.Println("waiting for goroutines done")
	return response.body
}

func main() {
	response := mirroredQuery()
	fmt.Println(response)
	panic("check")
}
