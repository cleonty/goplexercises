// main.go
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0
	letterCount := 0
	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
		if unicode.IsLetter(r) {
			letterCount++
		}
	}
	fmt.Printf("Rune\tCount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Printf("Len\tCount\n")
	for i, n := range utflen {
		fmt.Printf("%d\t%d\n", i, n)
	}
	fmt.Printf("\n%d буквенных симовлов UTF-8\n", letterCount)
	if invalid > 0 {
		fmt.Printf("\n%d неверных симовлов UTF-8\n", invalid)
	}
}
