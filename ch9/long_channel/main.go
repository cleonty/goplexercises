package main

import (
	"fmt"
	"time"
)

func f(in <-chan int, out chan<- int) {
	out <- <-in
}

func main() {
	in := make(chan int)
	ch := in
	var out chan int
	const n = 1000000
	for i := 0; i < n; i++ {
		out = make(chan int)
		go f(in, out)
		in = out
	}
	start := time.Now()
	ch <- 1
	<-out
	fmt.Printf("time=%s, n=%d\n", time.Since(start), n)
}
