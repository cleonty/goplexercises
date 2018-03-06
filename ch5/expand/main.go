package main

import (
	"fmt"
	"strings"
)

func expand(s string, f func(s string) string) string {
	return strings.Replace(s, "$foo", f("foo"), -1)
}

func main() {
	s := "$foo bar baz $foo"
	f := func(_ string) string { return fmt.Sprintf("%q", s) }
	fmt.Printf("%q -> %q", s, expand(s, f))
}
