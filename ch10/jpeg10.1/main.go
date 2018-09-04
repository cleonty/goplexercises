package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"

	"io"
	"os"
)

var format string

func init() {
	flag.StringVar(&format, "format", "", "output format png|jpg|gif")
}

func main() {
	flag.Parse()
	if err := convert(os.Stdin, os.Stdout, format); err != nil {
		fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
		os.Exit(1)
	}
}

func convert(in io.Reader, out io.Writer, format string) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintf(os.Stderr, "Входной формат %s\n", kind)
	switch format {
	case "jpeg", "jpg":
		err = jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
	case "png":
		err = png.Encode(out, img)
	case "gif":
		err = gif.Encode(out, img, nil)
	default:
		err = fmt.Errorf("неизвестный формат %v", format)
	}
	return err
}
