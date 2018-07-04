package main

import (
	"fmt"
	"os"

	"github.com/cleonty/gopl/ch5/links"
)

const maxdepth = 3

type WorkItem struct {
	link  string
	depth int
}

func crawl(url WorkItem) []WorkItem {
	if url.depth > maxdepth {
		return nil
	}
	fmt.Println(url.link)
	list, err := links.Extract(url.link)
	if err != nil {
		fmt.Println(err)
	}
	var wlist []WorkItem
	for _, link := range list {
		item := WorkItem{link, url.depth + 1}
		wlist = append(wlist, item)
	}
	return wlist
}



func main() {
	worklist := make(chan []WorkItem)
	unseenLinks := make(chan WorkItem)

	go func() {
		var wlist []WorkItem
		for _, link := range os.Args[1:] {
			item := WorkItem{link, 1}
			wlist = append(wlist, item)
		}
		worklist <- wlist
	}()

	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() { worklist <- foundLinks }()
			}
		}()
	}
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link.link] {
				seen[link.link] = true
				unseenLinks <- link
			}
		}
	}
}
