package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"sync"

	"golang.org/x/net/html"
)

func crawl(ctx context.Context, url string) []string {
	if cancelled(ctx) {
		fmt.Println("crawl: context cancelled")
		return nil
	}
	fmt.Println(url)
	list, err := Extract(ctx, url)
	if err != nil {
		fmt.Println(err)
	}
	return list
}

var done = make(chan struct{})

func cancelled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

func main() {
	var wg sync.WaitGroup
	worklist := make(chan []string)
	unseenLinks := make(chan string)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		worklist <- os.Args[1:]
	}()
	go func() {
		os.Stdin.Read(make([]byte, 1))
		cancel()
	}()

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			defer func() { wg.Done(); fmt.Printf("goroutine %d ends\n", i) }()
			for link := range unseenLinks {
				if cancelled(ctx) {
					return
				}
				foundLinks := crawl(ctx, link)
				if foundLinks != nil && !cancelled(ctx) {
					go func() { worklist <- foundLinks }()
				}
			}
		}(i)
	}
	seen := make(map[string]bool)
loop:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("ctx.done")
			break loop
		case list, ok := <-worklist:
			if !ok {
				fmt.Println("break loop !ok")
				break loop
			}
			if cancelled(ctx) {
				fmt.Println("break loop cancelled(ctx) worklist")
				break loop
			}
			for _, link := range list {
				if cancelled(ctx) {
					fmt.Println("irerating range list but context is cancelled")
					close(unseenLinks)
					break loop
				}
				if !seen[link] {
					seen[link] = true
					unseenLinks <- link
				}
			}
		}
	}
	fmt.Println("wait for gouroutines")
	wg.Wait()
	fmt.Println("waiting done")
	panic("check if there is a single goroutine with name 'main.main()'")
}

// Extract выполняет HTTP-запрос GET по определенному URL, выполняет
// синтаксический анализ HTML и возвращает ссылки в HTML-документе.
func Extract(ctx context.Context, url string) ([]string, error) {
	if cancelled(ctx) {
		fmt.Println("Extract: context cancelled")
		return nil, fmt.Errorf("cancelled")
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("extract: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("получение %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("анализ %s как HTML: %v", url, err)
	}
	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
