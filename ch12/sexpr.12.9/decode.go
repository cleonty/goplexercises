package sexpr

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"text/scanner"
)

type Token interface{}

type Symbol struct {
	Name string
}
type String struct {
	Name string
}

type Int struct {
	Number int
}

type StartList struct {
}

type EndList struct {
}

type lexer struct {
	scan  scanner.Scanner
	token rune // current token
}

func (lex *lexer) next()        { lex.token = lex.scan.Scan() }
func (lex *lexer) text() string { return lex.scan.TokenText() }
func (lex *lexer) consume(want rune) {
	if lex.token != want {
		panic(fmt.Sprintf("got %q, need %q", lex.text(), want))
	}
	lex.next()
}

func read(lex *lexer, v reflect.Value) {
	switch lex.token {
	case scanner.Ident:
		if lex.text() == "nil" {
			v.Set(reflect.Zero(v.Type()))
			lex.next()
			return
		}
	case scanner.String:
		s, _ := strconv.Unquote(lex.text())
		v.SetString(s)
		lex.next()
		return
	case scanner.Int:
		i, _ := strconv.Atoi(lex.text())
		v.SetInt(int64(i))
		lex.next()
		return
	case '(':
		lex.next()
		readList(lex, v)
		lex.consume(')')
		return
	}
}

func readList(lex *lexer, v reflect.Value) {
	switch v.Kind() {
	case reflect.Array: // (item ...)
		for i := 0; !endList(lex); i++ {
			read(lex, v.Index(i))
		}
	case reflect.Slice: // (item ...)
		for !endList(lex) {
			item := reflect.New(v.Type().Elem()).Elem()
			read(lex, item)
			v.Set(reflect.Append(v, item))
		}
	case reflect.Struct: // ((name value) ...)
		for !endList(lex) {
			lex.consume('(')
			if lex.token != scanner.Ident {
				panic(fmt.Sprintf("got %q, need field name", lex.text()))
			}
			name := lex.text()
			lex.next()
			read(lex, v.FieldByName(name))
			lex.consume(')')
		}
	case reflect.Map: // ((key value) ...)
		v.Set(reflect.MakeMap(v.Type()))
		for !endList(lex) {
			lex.consume('(')
			key := reflect.New(v.Type().Key()).Elem()
			read(lex, key)
			value := reflect.New(v.Type().Elem()).Elem()
			read(lex, value)
			v.SetMapIndex(key, value)
			lex.consume(')')
		}
	default:
		panic(fmt.Sprintf("can't decode list into %v", v.Type()))
	}

}

func endList(lex *lexer) bool {
	switch lex.token {
	case scanner.EOF:
		panic("unexpected end of file")
	case ')':
		return true
	}
	return false
}

// Unmarshal анализирует данные S-выражения и заполняет переменную,
// адресом которой является ненулевой указатель out.
func Unmarshal(data []byte, out interface{}) (err error) {
	decoder := NewDecoder(bytes.NewReader(data))
	return decoder.Decode(out)
	return nil
}

type Decoder struct {
	reader io.Reader
	lex    *lexer
}

func NewDecoder(r io.Reader) *Decoder {
	lex := &lexer{scan: scanner.Scanner{Mode: scanner.GoTokens}}
	dec := &Decoder{reader: r, lex: lex}
	lex.scan.Init(dec.reader)
	lex.next()
	return dec
}

func (dec *Decoder) Decode(v interface{}) (err error) {
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("error in %s: %v", dec.lex.scan.Position, x)
		}
	}()
	read(dec.lex, reflect.ValueOf(v).Elem())
	return nil
}

func (dec *Decoder) Token() (Token, error) {
	lex := dec.lex
	var token interface{}
	switch lex.token {
	case scanner.Ident:
		token = Symbol{lex.text()}
	case scanner.String:
		s, _ := strconv.Unquote(lex.text())
		token = String{s}
	case scanner.Int:
		i, _ := strconv.Atoi(lex.text())
		token = Int{i}
	case '(':
		token = StartList{}
	case ')':
		token = EndList{}
	default:
		return nil, fmt.Errorf("invalid token %q", lex.text())
	}
	lex.next()
	return token, nil
}

func (dec *Decoder) consume(token Token) {
	t, err := dec.Token()
	if err != nil {
		panic(err)
	}
	if reflect.TypeOf(t) != reflect.TypeOf(token) {
		panic(fmt.Sprintf("wrong type %q, expected %q", reflect.TypeOf(t), reflect.TypeOf(token)))
	}
}
