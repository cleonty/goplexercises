package main

import (
	"fmt"
)

func f() (ret int) {
	defer func() {
		recover()
		ret = 1234
	}()
	panic(123)
}

func main() {
	fmt.Println(f())
}
