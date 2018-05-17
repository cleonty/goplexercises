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
	var root = Element{}
	parseXML(&root, dec)
	fmt.Printf("%s\n", root)
}

func parseXML(parent *Element, dec *xml.Decoder) {
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			return
		} else if err != nil {
			panic(fmt.Sprintf("%s", err))
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			// fmt.Printf("begin StartElement %s\n", tok.Name.Local)
			element := &Element{tok.Name, tok.Attr, nil}
			parent.Children = append(parent.Children, Node(element))
			parseXML(element, dec)
			// fmt.Printf("end StartElement %s\n", tok.Name.Local)
		case xml.CharData:
			// fmt.Printf("CharData %s\n", tok)
			parent.Children = append(parent.Children, Node(CharData(tok)))
			continue
		case xml.EndElement:
			// fmt.Printf("EndElement %s\n", tok.Name.Local)
			return
		}
	}
}
