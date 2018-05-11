package main

import (
	"fmt"

	"github.com/cleonty/gopl/ch7/eval"
)

func main() {
	var line string
	fmt.Printf("Введите выражение: ")
	if _, err := fmt.Scanln(&line); err != nil {
		fmt.Println(err)
		return
	}
	expr, err := eval.Parse(line)
	if err != nil {
		fmt.Println(err)
		return
	}
	vars := make(map[eval.Var]bool)
	if err := expr.Check(vars); err != nil {
		fmt.Println(err)
		return
	}
	var env eval.Env = make(map[eval.Var]float64)
	for name := range vars {
		fmt.Printf("введите %s = ", name)
		var value float64
		if _, err := fmt.Scanf("%g\n", &value); err != nil {
			fmt.Println(err)
			return
		}
		env[name] = value
	}
	fmt.Printf("значение %s в среде %v равно %g\n", expr.String(), env, expr.Eval(env))

}
