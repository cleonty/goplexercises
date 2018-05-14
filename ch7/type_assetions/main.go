package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer
	w = os.Stdout
	f := w.(*os.File)

	//b := w.(*bytes.Buffer)
	fmt.Printf("%#v %T\n", f, f)
	//fmt.Printf("%#v %T\n", b, b)
	//os.PathError
	_, err := os.Open("/no/such/file")
	fmt.Printf("%#v %s %t", err, err, os.IsNotExist(err))
}
