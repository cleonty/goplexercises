package archive_test

import (
	"fmt"
	"io/ioutil"
	"testing"

	. "github.com/cleonty/gopl/ch10/archive"
	_ "github.com/cleonty/gopl/ch10/archive/tar"
	_ "github.com/cleonty/gopl/ch10/archive/zip"
)

func TestNewReader(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"tar test",
			args{"./test/test.tar"},
			false,
		},
		{
			"zip test",
			args{"./test/test.zip"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewReader(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewReader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil {
				t.Errorf("NewReader() is nil")
			}
			testReader(t, got)
		})
	}
}

func testReader(t *testing.T, reader Reader) {
	for reader.Next() {
		data, err := ioutil.ReadAll(reader)
		if err != nil {
			fmt.Printf("error %v\n", err)
			continue
		}
		fmt.Println(string(data))
	}
}
