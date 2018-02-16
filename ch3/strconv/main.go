// main.go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x))
	fmt.Println(strconv.FormatInt(int64(x), 2))
	fmt.Println(fmt.Sprintf("x=%b", x))
	x, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println(err)
	}
	z, err := strconv.ParseInt("123", 10, 16)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(z)
}
