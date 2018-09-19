package charcount

import (
	"io"
	"strings"
	"testing"
)

func TestCharCount(t *testing.T) {
	tests := []struct {
		name string
		r    io.Reader
		want int
	}{
		{"one letter", strings.NewReader("a"), 1},
		{"two letters", strings.NewReader("ab"), 2},
		{"russian", strings.NewReader("Леонтий"), 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CharCount(tt.r); got != tt.want {
				t.Errorf("CharCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
