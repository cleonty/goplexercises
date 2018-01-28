// ch1 project main.go
package main

import (
	"fmt"
	"os"
	"strings"
)

func printArgs1() {
	fmt.Println(strings.Join(os.Args[0:(len(os.Args))], " "))
}

func printArgs2() {
	for i, arg := range os.Args[1:] {
		fmt.Printf("%d %s\n", i+1, arg)
	}
}

func main() {
	printArgs1()
	printArgs2()
}
