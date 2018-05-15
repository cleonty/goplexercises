package main

import (
	"fmt"
	"strings"
	"text/scanner"
)

type selector struct {
	name  string
	attrs map[string]string
}

type lexer struct {
	scan  scanner.Scanner
	token rune // current lookahead token
}

func (lex *lexer) next()        { lex.token = lex.scan.Scan() }
func (lex *lexer) text() string { return lex.scan.TokenText() }

type lexPanic string

// describe returns a string describing the current token, for use in errors.
func (lex *lexer) describe() string {
	switch lex.token {
	case scanner.EOF:
		return "end of file"
	case scanner.Ident:
		return fmt.Sprintf("identifier %s", lex.text())
	case scanner.Int, scanner.Float:
		return fmt.Sprintf("number %s", lex.text())
	}
	return fmt.Sprintf("%q", rune(lex.token)) // any other rune
}

func parse(input string) (_ selector, err error) {
	defer func() {
		switch x := recover().(type) {
		case nil:
			// no panic
		case lexPanic:
			err = fmt.Errorf("%s", x)
		default:
			// unexpected panic: resume state of panic.
			panic(x)
		}
	}()
	lex := new(lexer)
	lex.scan.Init(strings.NewReader(input))
	lex.scan.Mode = scanner.ScanIdents | scanner.ScanStrings
	lex.next() // initial lookahead
	s := parseSelector(lex)
	if lex.token != scanner.EOF {
		return selector{}, fmt.Errorf("unexpected %s", lex.describe())
	}
	return s, nil
}

func parseSelector(lex *lexer) selector {
	var tagName string
	var attrs map[string]string
	if lex.token == scanner.Ident {
		tagName = lex.text()
		lex.next()
	}
	if lex.token == '[' {
		lex.next()
		if lex.token != ']' {
			attrs = make(map[string]string)
			name, value := parseAttr(lex)
			attrs[name] = value[1 : len(value)-1]
			for lex.token == ',' {
				lex.next()
				name, value := parseAttr(lex)
				attrs[name] = value[1 : len(value)-1]
			}
			if lex.token != ']' {
				msg := fmt.Sprintf("got %s, want ']'", lex.describe())
				panic(lexPanic(msg))
			}
			lex.next()
		}
	}
	return selector{tagName, attrs}
}

func parseAttr(lex *lexer) (name, value string) {
	if lex.token != scanner.Ident {
		msg := fmt.Sprintf("unexpected %s", lex.describe())
		panic(lexPanic(msg))
	}
	name = lex.text()
	lex.next()
	if lex.token != '=' {
		msg := fmt.Sprintf("got %q, want '='", lex.token)
		panic(lexPanic(msg))
	}
	lex.next()
	if lex.token != scanner.String {
		msg := fmt.Sprintf("want string, got %s", lex.describe())
		panic(lexPanic(msg))
	}
	value = lex.text()
	lex.next()
	return
}
