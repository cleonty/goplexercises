package main

import (
	"bytes"
	"fmt"
	"testing"
)

func Test_echo(t *testing.T) {
	type args struct {
		newline bool
		sep     string
		args    []string
	}
	tests := []struct {
		args    args
		want    string
		wantErr bool
	}{
		{args{true, "", []string{}}, "\n", false},
		{args{false, "", []string{}}, "", false},
		{args{true, "\t", []string{"one", "two", "three"}}, "one\ttwo\tthree\n", false},
		{args{true, ",", []string{"a", "b", "c"}}, "a,b,c\n", false},
		{args{false, ":", []string{"1", "2", "3"}}, "1:2:3", false},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("echo(%v, %q, %q)", tt.args.newline, tt.args.sep, tt.args.args)
		t.Run(name, func(t *testing.T) {
			out = new(bytes.Buffer)
			if err := echo(tt.args.newline, tt.args.sep, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("echo() error = %v, wantErr %v", err, tt.wantErr)
			}
			got := out.(*bytes.Buffer).String()
			if got != tt.want {
				t.Errorf("%q, want %q", got, tt.want)
			}
		})
	}
}
