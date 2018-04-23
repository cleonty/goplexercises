package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type IntSet struct {
}

func (s *IntSet) String() string {
	return fmt.Sprintf("IntSet %d", s)
}

func main() {
	var w io.Writer
	w = os.Stdout
	w = new(bytes.Buffer)
	//w = time.Second

	var rwc io.ReadWriteCloser
	rwc = os.Stdout
	//rwc = new(bytes.Buffer)

	w = rwc
	w.Write([]byte{'a'})
	//rwc = w

	//var _ = IntSet{}.String()
	var s IntSet
	fmt.Println(s.String())

	var _ fmt.Stringer = &s
	//var _ fmt.Stringer = s

	var any interface{}
	any = 123
	any = "string"
	any = "hello"

	any = map[string]int{"one": 1}
	any = new(bytes.Buffer)
	fmt.Println(any)
	{
		var _ io.Writer = new(bytes.Buffer)
		var _ io.Writer = (*bytes.Buffer)(nil)

	}

}
