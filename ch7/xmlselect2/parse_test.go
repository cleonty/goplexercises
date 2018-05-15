package main

import (
	"testing"
)

func TestParse(t *testing.T) {
	text := `div[a="hello",be="world"]`
	s, err := parse(text)
	if err != nil {
		t.Error(err)
		return
	}
	if s.name != "div" {
		t.Errorf(`expected name "div", got %q"`, s.name)
	}
	if s.attrs["a"] != "hello" {
		t.Errorf(`expected attr a="hello", got %q"`, s.attrs["a"])
	}
	if s.attrs["be"] != "world" {
		t.Errorf(`expected attr be="world", got %q"`, s.attrs["be"])
	}
}
