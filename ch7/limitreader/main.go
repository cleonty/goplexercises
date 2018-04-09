package main

import (
	"fmt"
	"io"
	"strings"
)

type LimitedReader struct {
	rd    io.Reader
	limit int
}

func (reader *LimitedReader) Read(p []byte) (n int, err error) {
	if reader.limit > 0 {
		if len(p) <= reader.limit {
			n, err = reader.rd.Read(p)
		} else {
			n, err = reader.rd.Read(p[0:reader.limit])
		}
		reader.limit -= n
	} else {
		return 0, io.EOF
	}
	return
}

func LimitReader(rd io.Reader, n int) io.Reader {
	return &LimitedReader{rd, n}
}

var a io.LimitedReader

func main() {
	rd := LimitReader(strings.NewReader("Leonty Chudinov"), 5)
	p := make([]byte, 10)
	n, _ := rd.Read(p)
	fmt.Println(string(p[0:n]))
}
