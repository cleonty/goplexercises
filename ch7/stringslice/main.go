package main

import (
	"fmt"
	"sort"
)

type StringSlice []string

func (s StringSlice) Len() int {
	return len(s)
}

func (s StringSlice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s StringSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	s := StringSlice([]string{"hello", "world", "opera"})
	sort.Sort(s)
	fmt.Println(s)
}
