// main.go
package main

import (
	"fmt"

	"gopl.io/ch2/tempconv"
)

func main() {
	c := tempconv.FToC(212.0)
	fmt.Println(c.String())
	fmt.Printf("%v\n", c)
	fmt.Printf("%s\n", c)
	fmt.Println(c)
	fmt.Printf("%g\n", c)
	fmt.Println(float64(c))
}
