package main

import (
	"bytes"
	"io"
	"os"
)

func main() {
	var w io.Writer
	w = os.Stdout
	w = new(bytes.Buffer)
	w = nil
}
