package eval

import (
	"fmt"
	"strings"
)

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return fmt.Sprintf("%g", l)
}

func (u unary) String() string {
	return fmt.Sprintf("%c%s", u.op, u.x)
}

func (b binary) String() string {
	if b.op == '-' || b.op == '+' {
		return fmt.Sprintf("(%s %c %s)", b.x, b.op, b.y)
	}
	return fmt.Sprintf("%s %c %s", b.x, b.op, b.y)
}

func (c call) String() string {
	var args []string
	for _, arg := range c.args {
		args = append(args, arg.String())
	}
	return fmt.Sprintf("%s(%s)", c.fn, strings.Join(args, ", "))
}

func (m min) String() string {
	return fmt.Sprintf("min(%s, %s)", m.x, m.y)
}
