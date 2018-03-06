package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func startElementWithID(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		for _, attr := range n.Attr {
			if attr.Key == "id" && attr.Val == id {
				fmt.Printf("startElementWithID %q found tag is %q\n", id, n.Data)
				return true
			}
		}
	}
	return false
}

func endElementWithID(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		for _, attr := range n.Attr {
			if attr.Key == "id" && attr.Val == id {
				fmt.Printf("endElementWithID %s found\n", id)
				return true
			}
		}
	}
	return false
}

func main() {
	url := "http://www.e1.ru"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		fmt.Println(err)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	n := ElementByID(doc, "fb-root")
	if n != nil {
		fmt.Printf("найден %q\n", n.Data)
	} else {
		fmt.Printf("не найден %q\n", "lffb-root")
	}

}

// ElementByID возврщает элемент по ID
func ElementByID(doc *html.Node, id string) *html.Node {
	startElement := func(n *html.Node) bool { return startElementWithID(n, id) }
	return forEachNode(doc, startElement, nil)
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) *html.Node {
	node := n
	if pre != nil {
		if pre(n) {
			fmt.Printf("(pre)найден узел %q\n", node.Data)
			return node
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		node = forEachNode(c, pre, post)
		if node != nil {
			fmt.Printf("передаю найденный тег %q на один уровень вверх, родительский тег %q\n", node.Data, n.Data)
			return node
		}
	}
	if post != nil {
		if post(n) {
			fmt.Printf("(post)found node %v", node.Data)
			return node
		}
	}
	return nil
}
