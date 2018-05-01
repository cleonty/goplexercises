package main

import (
	"fmt"
	"sort"
)

func IsPalindrom(s sort.Interface) bool {
	for i, j := 0, s.Len()-1; i < j; i, j = i+1, j-1 {
		if !s.Less(i, j) && !s.Less(j, i) {
			continue
		} else {
			return false
		}
	}
	return true
}

func main() {
	v := []int{1, 2, 3, 3, 2, 1}
	fmt.Println(IsPalindrom(sort.IntSlice(v)))
}
