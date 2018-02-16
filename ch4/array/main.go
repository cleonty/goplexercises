// main.go
package main

import (
	"crypto/sha256"
	"fmt"
)

func one() {
	fmt.Println("Start one")
	defer fmt.Println("End one")
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])
	for i, v := range a {
		fmt.Println(i, v)
	}
	for _, v := range a {
		fmt.Println(v)
	}
	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	fmt.Println(r[2], q[2])
	s := [...]int{1, 2, 3}
	//s = [...]int{1, 2, 3, 4} error
	fmt.Println(s)

	type Currency int
	const (
		USD Currency = iota
		EUR
		GBR
		RUR
	)
	symbol := [...]string{USD: "USD", EUR: "EUR", GBR: "GBR", RUR: "RUR", 99: "None"}
	fmt.Println(symbol)
	{
		a := [2]int{1, 2}
		b := [...]int{1, 2}
		c := [2]int{1, 3}
		fmt.Println(a == b, b == c, a == c, a != b)
	}
	{
		c1 := sha256.Sum256([]byte{'x'})
		c2 := sha256.Sum256([]byte{'X'})
		fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	}
}
func zero1(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}

func zero2(ptr *[32]byte) {
	*ptr = [32]byte{}
}

func main() {
	one()
}
