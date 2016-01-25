package main

import (
	"fmt"
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

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
	// should be
	//
	// 1:	intro to programming
	// 2:	discrete math
	// 3:	data structures
	// 4:	algorithms
	// 5:	linear algebra
	// 6:	calculus
	// 7:	formal languages
	// 8:	computer organization
	// 9:	compilers
	// 10:	databases
	// 11:	operating systems
	// 12:	networks
	// 13:	programming languages
	//
	// but
	//
	// 1:	computer organization
	// 2:	intro to programming
	// 3:	discrete math
	// 4:	data structures
	// 5:	operating systems
	// 6:	networks
	// 7:	programming languages
	// 8:	algorithms
	// 9:	linear algebra
	// 10:	calculus
	// 11:	formal languages
	// 12:	compilers
	// 13:	databases
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items map[string]bool)

	visitAll = func(items map[string]bool) {
		for item, _ := range items {
			if !seen[item] {
				seen[item] = true
				// visitAll(m[item])
				mm := make(map[string]bool, len(m[item]))
				for _, p := range m[item] {
					mm[p] = true
				}
				visitAll(mm)
				order = append(order, item)
			}
		}
	}

	keys := make(map[string]bool, len(m))
	for key := range m {
		keys[key] = true
	}

	visitAll(keys)
	return order
}
