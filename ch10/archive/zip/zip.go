package zip

import (
    "fmt"
	"io"
	"archive/zip"

	"github.com/cleonty/gopl/ch10/archive"
)

type zipReader struct {
    zrc *zip.ReadCloser
    fileIndex int
    frc io.ReadCloser
}

func (r *zipReader) Next() bool {
	if r.fileIndex < len(r.zrc.File) - 1 {
        r.fileIndex++
        file := r.zrc.File[r.fileIndex]
        rc, err := file.Open()
        if err != nil {
            fmt.Printf("zip error: %v\n", err)
            return false
        }
        r.frc = rc
        fmt.Printf("===> file: %s\n", file.Name)
        return true
    }
    return false
}

func (r *zipReader) Read(p []byte) (n int, err error) {
    if r.frc == nil {
        return 0, fmt.Errorf("no current file")
    }
    return r.frc.Read(p)
}

// NewReader creates a new archive reader for tar archives
func NewReader(filename string) (archive.Reader, error) {
	rc, err := zip.OpenReader(filename)
	if err != nil {
		return nil, err
	}
	return &zipReader{rc, -1, nil}, nil
}

func init() {
	archive.RegisterFormat("zip", NewReader)
}
