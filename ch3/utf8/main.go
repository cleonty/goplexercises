// main.go
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Hello, \U00004e16\U0000754c"
	fmt.Println(len(s))                    // "13"
	fmt.Println(utf8.RuneCountInString(s)) // "9"
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	for i, r := range "Hello, \U00004e16\U0000754c" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
	var n int
	for range s {
		n++
	}
	fmt.Println(n)
	s = "プログラム"
	fmt.Printf("% x\n", s)
	r := []rune(s)
	fmt.Printf("%x\n", r)
	fmt.Println(string(r))
	fmt.Println(string(65))
	fmt.Println(string(1234567))
}

func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

func Contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}
