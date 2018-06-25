package main

import (
	"fmt"
)

func main() {
	naturals := make(chan int)
	doubles := make(chan int)
	identities := make(chan int)
	triples := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go doubler(doubles, naturals)
	go identer(identities, doubles)
	go tripler(triples, identities)
	go squarer(squares, triples)
	printer(squares)
}

func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

func doubler(out chan<- int, in <-chan int) {
	for x := range in {
		out <- 2 * x
	}
	close(out)
}

func identer(out chan<- int, in <-chan int) {
	for x := range in {
		out <- x
	}
	close(out)
}

func tripler(out chan<- int, in <-chan int) {
	for x := range in {
		out <- 3 * x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}pS7529Pvd4

func printer(in <-chan int) {
	for x := range in {
		fmt.Println(x)
	}
}
