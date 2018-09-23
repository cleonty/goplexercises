package sprint

import (
	"fmt"
	"testing"
)

type MyInt int

func (i MyInt) String() string {
	return fmt.Sprintf("MyInt(%d)", int(i))
}
func TestSprint(t *testing.T) {
	type args struct {
		x interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"integer", args{1}, "1"},
		{"string", args{"abc"}, "abc"},
		{"bool", args{true}, "true"},
		{"bool", args{false}, "false"},
		{"stringer", args{MyInt(5)}, "MyInt(5)"},
		{"unknown", args{struct{}{}}, "???"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sprint(tt.args.x); got != tt.want {
				t.Errorf("Sprint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleSprint() {
	fmt.Println(Sprint(123))
	fmt.Println(Sprint(struct{}{}))
	fmt.Println(Sprint(MyInt(77)))
	// Output:
	// 123
	// ???
	// MyInt(77)
}

func BenchmarkSprint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sprint(123)
	}
}
