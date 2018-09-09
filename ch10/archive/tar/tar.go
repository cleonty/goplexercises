package tar

import (
	"archive/tar"
	"fmt"
	"os"

	"github.com/cleonty/gopl/ch10/archive"
)

type tarReader struct {
	*tar.Reader
}

func (r *tarReader) Next() bool {
	header, err := r.Reader.Next()
	if (err !=nil) {
		fmt.Printf("tar Next() error: %v\n", err)
		return false
	}
	fmt.Printf("===> file: %s\n", header.Name)
	return true;
}

func (r *tarReader) Read(p []byte) (n int, err error) {
	return r.Reader.Read(p)
}

// NewReader creates a new archive reader for tar archives
func NewReader(filename string) (archive.Reader, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	reader := tar.NewReader(file)
	return &tarReader{reader}, nil
}

func init() {
	archive.RegisterFormat("tar", NewReader)
}
