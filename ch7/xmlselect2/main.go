package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []xml.StartElement
	var selectors []selector
	for _, arg := range os.Args[1:] {
		s, err := parse(arg)
		if err != nil {
			fmt.Println(err)
			return
		}
		selectors = append(selectors, s)
	}
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("%s: %s\n", os.Args[0], err)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok)
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			if containsAll(stack, selectors) {
				fmt.Printf("%s: %s\n", strings.Join(os.Args[1:], " "), tok)
			}
		}
	}
}

func containsAll(x []xml.StartElement, y []selector) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if matchSelector(x[0], y[0]) {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}

func matchSelector(x xml.StartElement, s selector) bool {
	if s.name != "" && x.Name.Local != s.name {
		return false
	}
	if s.attrs != nil {
		for _, attr := range x.Attr {
			if val, ok := s.attrs[attr.Name.Local]; ok {
				if attr.Value != val {
					return false
				}
			}
		}
	}
	return true
}
