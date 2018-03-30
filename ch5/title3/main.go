package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

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

func title(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		return fmt.Errorf("%s имеет тип %s, не text/html", url, ct)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return fmt.Errorf("анализ %s как HTML: %v", url, err)
	}
	title, err := soleTitle(doc)
	if err != nil {
		return err
	}
	fmt.Printf("url %s title %s\n", url, title)
	return nil
}

func soleTitle(doc *html.Node) (title string, err error) {
	type bailout struct{}
	defer func() {
		switch p := recover(); p {
		case nil:
			//no panic
		case bailout{}:
			err = fmt.Errorf("несколько элементов title")
		default:
			panic(p)
		}
	}()
	forEachNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			if title != "" {
				panic(bailout{})
			}
			title = n.FirstChild.Data
		}
	}, nil)
	if title == "" {
		return "", fmt.Errorf("нет элемента title")
	}
	return title, nil
}

func main() {
	args := []string{"http://google.ru", "http://66.ru", "https://golang.org/doc/gopher/frontpage.png"}
	for _, url := range args {
		err := title(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		}
	}
}
