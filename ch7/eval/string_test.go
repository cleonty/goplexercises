package eval

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	tests := []struct {
		expr string
		want string
	}{
		{"sqrt(A / pi)", "sqrt(A / pi)"},
		{"pow(x, 3) + pow(y, 3)", "(pow(x, 3) + pow(y, 3))"},
		{"5 / 9 * (F - 32)", "5 / 9 * (F - 32)"},
	}
	for _, test := range tests {
		fmt.Printf("\n%s\n", test.expr)
		expr, err := Parse(test.expr)
		if err != nil {
			t.Error(err)
			continue
		}
		got := expr.String()
		fmt.Printf("\t => %s\n", got)
		if got != test.want {
			t.Errorf("%s.String() = %q, требуется %q\n",
				test.expr, got, test.want)
		}
	}
}
