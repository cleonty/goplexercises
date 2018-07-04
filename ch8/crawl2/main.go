package main

import (
	"fmt"
	"os"

	"github.com/cleonty/gopl/ch5/links"
)

var n int

// breadthFirst вызывает f для каждого элемента в worklist.
// Все элементы, возвращаемые f, добавляются в worklist.
// f вызывается для каждого элемента не более одного раза
func breadthFirst(f func(item string) []string, worklist chan []string) {
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <- worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- f(link)
				}(link)
			}
		}
	}
}

var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{}
	list, err := links.Extract(url)
	<- tokens
	if err != nil {
		fmt.Println(err)
	}
	return list
}

func main() {
	worklist := make (chan []string)
	n++
	go func (){
		worklist <- os.Args[1:]
	}()
	breadthFirst(crawl, worklist)
}
