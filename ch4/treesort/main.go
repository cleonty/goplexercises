// main.go
package main

import (
	"fmt"
	"strconv"
)

type tree struct {
	value       int
	left, right *tree
}

func (t *tree) String() string {
	if t != nil {
		return t.left.String() + " " + strconv.Itoa(t.value) + " " + t.right.String()
	}
	return ""
}

func Sort(values []int) {
	var root *tree
	for _, value := range values {
		root = addRoot(root, value)
	}
	appendValues(values[:0], root)
	fmt.Println(root)
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func addRoot(t *tree, value int) *tree {
	if t == nil {
		return &tree{value: value}
	} else if value < t.value {
		t.left = addRoot(t.left, value)
	} else {
		t.right = addRoot(t.right, value)
	}
	return t
}

func main() {
	var array = []int{2, 3, 4, 5, 67, 77, 2, 3, 1}
	Sort(array)
	fmt.Println(array, len(array), cap(array))
}
