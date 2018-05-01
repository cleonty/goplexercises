package main

import (
	"bytes"
	"fmt"
	"io"
)

const debug = false

func main() {
	//var buf *bytes.Buffer
	var buf io.Writer
	if debug {
		buf = new(bytes.Buffer)
	}
	f(buf)
	if debug {
		fmt.Println(buf)
	}
}

func f(out io.Writer) {
	if out != nil {
		out.Write([]byte("выполнено!\n"))
	}
}
