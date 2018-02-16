// main.go
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(basename1("Hello World"))
	fmt.Println(basename1("meduza/a/b.go"))
	fmt.Println(basename2("Hello World"))
	fmt.Println(basename2("meduza/a/b.go"))
}

func basename1(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func basename2(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}
