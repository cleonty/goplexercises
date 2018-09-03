package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup
var count int

func ping(in <-chan int, out chan<- int, cancel <-chan struct{}) {
	defer wg.Done()
	for {
		var ball int
		select {
		case ball = <-in:
			select {
			case out <- ball:
				count++
			case <-cancel:
				return
			}
		case <-cancel:
			return
		}
	}
}

func main() {
	const seconds = 3
	ch1 := make(chan int)
	ch2 := make(chan int)
	cancel := make(chan struct{})
	wg.Add(2)
	go ping(ch1, ch2, cancel)
	go ping(ch2, ch1, cancel)
	ch1 <- 1
	<-time.After(seconds * time.Second)
	close(cancel)
	wg.Wait()
	fmt.Printf("ball pass count %d per second\n", count/seconds)
	fmt.Printf("GOMAXPROCS %d\n", runtime.GOMAXPROCS(0))
}
