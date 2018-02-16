// main.go
package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	const noDelay time.Duration = 0
	const timeout = 5 * time.Minute
	fmt.Printf("%T %[1]v\n", noDelay)
	fmt.Printf("%T %[1]v\n", timeout)
	fmt.Printf("%T %[1]v\n", time.Minute)
	const (
		a = 1
		b
		c = 2
		d
	)
	fmt.Println(a, b, c, d)
	type Weekday int
	const (
		Monday Weekday = iota
		Thusday
	)
	fmt.Printf("%T %[1]v\n", Thusday)
	type Flags int
	const (
		FirstFlag Flags = 1 << iota
		SecondFlag
		ThirdFlag
		FourFlag
		FiveFlag
	)
	fmt.Printf("%T %[1]v\n", FirstFlag)
	fmt.Printf("%T %[1]v\n", SecondFlag)
	fmt.Printf("%T %[1]v\n", ThirdFlag)
	fmt.Printf("%T %[1]v\n", FourFlag)
	fmt.Printf("%T %[1]v\n", FiveFlag)
	const (
		_ = 1 << (10 * iota)
		KiB
		MiB
		GiB
		TiB
		EiB
		ZiB
		YiB
	)
	fmt.Printf("%T %[1]v\n", KiB)
	fmt.Printf("%T %[1]v\n", MiB)
	fmt.Printf("%T %[1]v\n", GiB)
	fmt.Printf("%T %[1]v\n", TiB)
	fmt.Printf("%T %[1]v\n", YiB/MiB)

	const (
		n1 float32    = math.Pi
		n2 float64    = math.Pi
		n3 complex128 = math.Pi
	)
	fmt.Println(n1, n2, n3)

	var f float64 = 3 + 0i // Нетипизированное комплексное -> float64
	f = 2                  // Нетипизированное целое -> float64
	f = 1e123              // Нетипизированное действительное -> float64
	f = 'a'                // Нетипизированная руна -> float64
	fmt.Println(f)
}
