package eval

import (
	"fmt"
	"strings"
)

func (v Var) Check(vars map[Var]bool) error {
	vars[v] = true
	return nil
}

func (l literal) Check(vars map[Var]bool) error {
	return nil
}

func (u unary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-", u.op) {
		return fmt.Errorf("некорректный унарный оператор %q", u.op)
	}
	return u.x.Check(vars)
}

func (b binary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-*/", b.op) {
		return fmt.Errorf("некорректный бинарный оператор %q", b.op)
	}
	if err := b.x.Check(vars); err != nil {
		return err
	}
	return b.y.Check(vars)
}

func (c call) Check(vars map[Var]bool) error {
	arity, ok := numParams[c.fn]
	if !ok {
		return fmt.Errorf("неизвестная функция %q", c.fn)
	}
	if len(c.args) != arity {
		return fmt.Errorf("вызов %q имеет %d вместо %d аргументов",
			c.fn, len(c.args), arity)
	}
	for _, arg := range c.args {
		if err := arg.Check(vars); err != nil {
			return err
		}
	}
	return nil
}

func (m min) Check(vars map[Var]bool) error {
	if err := m.x.Check(vars); err != nil {
		return err
	}
	return m.y.Check(vars)
}

var numParams = map[string]int{"pow": 2, "sin": 1, "sqrt": 1}
