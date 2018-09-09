package archive

import (
	"fmt"
	"io"
	"path"
)

// Reader allows to iterate trought files in the archive and read their contents
type Reader interface {
	Next() bool
	io.Reader
}

var formats = make(map[string]func(filename string) (Reader, error))

// RegisterFormat registers new reader for a given archive type
func RegisterFormat(ext string, newReader func(filename string) (Reader, error)) {
	fmt.Printf("register reader for %s\n", ext)
	formats[ext] = newReader
}

// NewReader searches for a suitable Reader for a given archive
func NewReader(filename string) (Reader, error) {
	ext := path.Ext(filename)
	if len(ext) > 0 && ext[0] == '.' {
		ext = ext[1:]
	}
	newReader, ok := formats[ext]
	if !ok {
		return nil, fmt.Errorf("неизвестный формат архива %s", ext)
	}
	return newReader(filename)
}
