package main

import (
	"fmt"
	"os"

	"github.com/cleonty/gopl/ch5/links"
)

// breadthFirst вызывает f для каждого элемента в worklist.
// Все элементы, возвращаемые f, добавляются в worklist.
// f вызывается для каждого элемента не более одного раза
func breadthFirst(f func(item string) []string, worklist chan []string) {
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					worklist <- f(link)
				}(link)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		fmt.Println(err)
	}
	return list
}

func main() {
	worklist := make (chan []string)
	go func (){
		worklist <- os.Args[1:]
	}()
	breadthFirst(crawl, worklist)
}
