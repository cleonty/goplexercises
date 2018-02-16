// main.go
package main

import (
	"fmt"
	"sort"
)

func main() {
	{
		ages := make(map[string]int)
		fmt.Println(ages)
	}
	{
		var ages map[string]int = make(map[string]int)
		ages["alice"] = 10
		fmt.Println(ages)
		fmt.Println("noney:", ages["nonkey"])
	}
	{
		var ages map[string]int = map[string]int{}
		ages["alice"] = 10
		fmt.Println(ages)
	}
	{
		ages := map[string]int{
			"alice": 10,
			"bob":   5,
			"joe":   4,
			"don":   4,
		}
		delete(ages, "alice")
		delete(ages, "carl")
		ages["carl"] = ages["carl"] + 1
		ages["carl"] += 1
		ages["carl"]++
		fmt.Println(ages)
		for name, age := range ages {
			fmt.Printf("%s\t%d\n", name, age)
		}

		names := make([]string, 0, len(ages))
		for name := range ages {
			names = append(names, name)
		}
		sort.Strings(names)
		for _, name := range names {
			fmt.Printf("%s\t%d\n", name, ages[name])
		}
		age, ok := ages["leo"]
		if !ok {
			fmt.Println("age for leo not found")
		} else {
			fmt.Println("age for leo found", age)
		}
		if age, ok := ages["leo"]; ok {
			fmt.Println("age for leo found", age)
		}
	}
	{
		ages1 := map[string]int{
			"alice": 10,
			"bob":   5,
			"joe":   4,
			"don":   4,
			"leo":   7,
		}
		ages2 := map[string]int{
			"alice": 10,
			"bob":   5,
			"joe":   4,
			"don":   4,
			"op":    6,
		}

		fmt.Println(equal(ages1, ages2))

	}
	{
		Add([]string{"aba", "baba", "haba"})
		Add([]string{"aba", "baba", "haba"})
		Add([]string{"aba", "baba", "haba"})
		fmt.Println(Count([]string{"aba", "baba", "haba"}))
	}
}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || xv != yv {
			return false
		}
	}
	return true
}

var m = make(map[string]int)

func k(list []string) string {
	return fmt.Sprintf("%q", list)
}

func Add(list []string) {
	m[k(list)]++
}

func Count(list []string) int {
	return m[k(list)]
}
