package main

import (
	"fmt"
	"os"
)

func sum(vals ...int) int {
	sum := 0
	for _, val := range vals {
		sum += val
	}
	return sum
}

func errorf(linenum int, format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "Строка %d: ", linenum)
	fmt.Fprintf(os.Stderr, format, args)
	fmt.Fprintln(os.Stderr)
}

func min(first int, args ...int) int {
	min := first
	for _, val := range args {
		if val < min {
			min = val
		}
	}
	return min
}

func main() {
	fmt.Println(sum())
	fmt.Println(sum(1, 2))
	fmt.Println(sum(1, 2, 3))
	vals := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(vals...))
	fmt.Println(min(vals[0], vals[1:]...))
}
