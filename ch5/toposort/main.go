package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}
var prereqsCyclic = map[string][]string{
	"algorithms":     {"data structures"},
	"calculus":       {"linear algebra"},
	"linear algebra": {"calculus"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}

func topoSort2(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}
	for key := range m {
		visitAll([]string{key})
	}
	return order
}

func find(items []string, item string) bool {
	for _, i := range items {
		if i == item {
			return true
		}
	}
	return false
}

func topoSort3(m map[string][]string) ([]string, error) {
	var order []string
	resolved := make(map[string]bool)
	var visitAll func(items []string, parents []string) error
	visitAll = func(items []string, parents []string) error {
		if items == nil {
			return nil
		}
		for _, item := range items {
			itemResolved, ok := resolved[item]
			if ok && !itemResolved {
				return fmt.Errorf("cycle detected: %s", strings.Join(append(parents, item), " -> "))
			}
			if !ok {
				resolved[item] = false
				if err := visitAll(m[item], append(parents, item)); err != nil {
					return err
				}
				resolved[item] = true
				order = append(order, item)
			}
		}
		return nil
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return order, visitAll(keys, nil)
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
	for i, course := range topoSort2(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
	order, err := topoSort3(prereqsCyclic)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	for i, course := range order {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}
