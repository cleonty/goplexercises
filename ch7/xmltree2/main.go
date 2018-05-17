package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"strings"
)

// Node is CharData or *Element
type Node interface {
	String() string
}

// CharData is a string
type CharData string

// Element is an XML element
type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func (e *Element) String() string {
	var children []string
	for _, c := range e.Children {
		children = append(children, c.String())
	}
	return fmt.Sprintf("[%s: %s]", e.Type.Local, strings.Join(children, " "))
}

func (c CharData) String() string {
	return string("\"" + c + "\"")
}

type xmlError string

func main() {
	dec := xml.NewDecoder(strings.NewReader("<div><a>hello</a><b>x</b>jeo<x></x></div>"))
	root := parseXML(dec)
	fmt.Printf("%s\n", root)
}

func parseXML(dec *xml.Decoder) Node {
	var root Node
	var stack []*Element
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(fmt.Sprintf("%s", err))
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			element := &Element{tok.Name, tok.Attr, nil}
			if root != nil {
				stack[len(stack)-1].Children = append(stack[len(stack)-1].Children, element)
			} else {
				root = element
			}
			stack = append(stack, element)
		case xml.CharData:
			stack[len(stack)-1].Children = append(stack[len(stack)-1].Children, CharData(tok))
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		}
	}
	return root
}
