// mian.go
package main

import (
	"fmt"
)

func test() {
	var fruits []string = []string{"apples", "bananas", "oranges"}
	for i := len(fruits) - 1; i >= 0; i-- {
		fmt.Println(i)
		fmt.Println(fruits[i])
	}
	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o) // "438 666 0666"
	x := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)
	// Вывод:
	// 3735928559 deadbeef 0xdeadbeef 0XDEADBEEF
	ascii := 'а'
	Unicode := '🟊'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)
	fmt.Printf("%d %[1]c %[1]q\n", Unicode)
	fmt.Printf("%d %[1]q\n", newline)
}

func main() {
	test()
}
