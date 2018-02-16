// main.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	freq := make(map[string]int)
	input := bufio.NewScanner(strings.NewReader("мама\nмама\nмама\nмыла\nраму мама папа"))
	input.Split(bufio.ScanWords)
	for input.Scan() {
		word := input.Text()
		freq[word]++
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "wordfreq: %v\n", err)
	}
	for w, f := range freq {
		fmt.Printf("%s\t%d\n", w, f)
	}
}
