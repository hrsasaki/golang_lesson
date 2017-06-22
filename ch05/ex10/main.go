// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 136.

// The toposort program prints the nodes of a DAG in topological order.
package main

import "fmt"

//!+table
// prereqs maps computer science courses to their prerequisites.
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

//!-table

//!+main
func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	depthMap := make(map[string]int)
	seen := make(map[string]bool)
	var visitAll func(items []string)
	depth := 0

	visitAll = func(items []string) {
		depth++
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				depthMap[item] = depth
			}
		}
		depth--
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	visitAll(keys)
	found := false
	i := 1
	var order []string
	for {
		for key, depth := range depthMap {
			if depth == i {
				order = append([]string{key}, order...)
				found = true
			}
		}
		if !found {
			break
		}
		i++
		found = false
	}

	return order
}

//!-main
