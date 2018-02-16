// main.go
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma2("123456789"))
	s := "abc"
	b := []byte(s)
	_ = string(b)
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + comma(s[n-3:])
}

func comma2(s string) string {
	var buf bytes.Buffer
	for i := 0; i < len(s); i++ {
		if i > 0 && (len(s)-i)%3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(s[i])
	}
	return buf.String()
}
