package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n)
	fmt.Printf("\rFibonachi(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, c := range `/|\-` {
			fmt.Printf("\r%c", c)
			time.Sleep(delay)
		}

	}
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
