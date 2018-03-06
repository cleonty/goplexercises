package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinksl: %v\n", err)
		os.Exit(1)
	}
	count := make(map[string]int)
	visit(count, doc)
	fmt.Println(count)
}

func visit(count map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		count[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(count, c)
	}
}
