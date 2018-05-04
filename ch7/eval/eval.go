package eval

import (
	"fmt"
	"math"
)

// Expr представляет арифметическое выражение
type Expr interface {
	// Eval возвращает значение данного Ехрг в среде env.
	Eval(env Env) float64
	// Check сообщает об ошибках в данном Ехрг и добавляет свои Vars.
	Check(vars map[Var]bool) error
}

// Var определяет переменную, например x.
type Var string

// literal представляет собой числовую константу, например 3.141.
type literal float64

// unary представляет выражение с унарным оператором, например -х.
type unary struct {
	op rune // + или -
	x  Expr
}

// binary представляет выражение с бинарным оператором, например х+у.
type binary struct {
	op   rune // -, +, *, /
	x, y Expr
}

// call представляет выражение вызова функции, например sin(x).
type call struct {
	fn   string // одно из "pow", "sin", "sqrt"
	args []Expr
}

// Env отображает имена переменных на значения
type Env map[Var]float64

func (v Var) Eval(env Env) float64 {
	return env[v]
}

func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("неподдерживаемый унарный оператор: %q", u.op))
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("неподдерживаемый бинарный оператор: %q", b.op))
}

func (с call) Eval(env Env) float64 {
	switch с.fn {
	case "sin":
		return math.Sin(с.args[0].Eval(env))
	case "pow":
		return math.Pow(с.args[0].Eval(env), с.args[1].Eval(env))
	case "sqrt":
		return math.Sqrt(с.args[0].Eval(env))
	}
	panic(fmt.Sprintf("неподдерживаемый вызов функции: %q", с.fn))
}
