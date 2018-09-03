package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("GOMAXPROCS is %d\n", runtime.GOMAXPROCS(0))
	for i := 0; i < 1000; i++ {
		fmt.Print(1)
		go fmt.Print(0)
	}
}
