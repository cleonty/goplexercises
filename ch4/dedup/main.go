// main.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	seen := make(map[string]bool)
	input := bufio.NewScanner(strings.NewReader("мама\nмама\nмама\nмыла\nраму\nмама\nпапа"))
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
	}
}
