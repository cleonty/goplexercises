package sexpr

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"text/scanner"
)

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
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{reader: r}
}
func (dec *Decoder) Decode(v interface{}) (err error) {
	lex := &lexer{scan: scanner.Scanner{Mode: scanner.GoTokens}}
	lex.scan.Init(dec.reader)
	lex.next()
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("error in %s: %v", lex.scan.Position, x)
		}
	}()
	read(lex, reflect.ValueOf(v).Elem())
	return nil
}
