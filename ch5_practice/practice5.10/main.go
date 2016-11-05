// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 136.

// The toposort program prints the nodes of a DAG in topological order.
package main

import "fmt"

//!+table
// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string]map[string]bool{
	"algorithms": {"data structures": true},
	"calculus":   {"linear algebra": true},

	"compilers": {
		"data structures":       true,
		"formal languages":      true,
		"computer organization": true,
	},

	"data structures":       {"discrete math": true},
	"databases":             {"data structures": true},
	"discrete math":         {"intro to programming": true},
	"formal languages":      {"discrete math": true},
	"networks":              {"operating systems": true},
	"operating systems":     {"data structures": true, "computer organization": true},
	"programming languages": {"data structures": true, "computer organization": true},
}

//!-table

//!+main
func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
	//topoSort(prereqs)
}

func topoSort(m map[string]map[string]bool) []string {
	var order []string
	seen := make(map[string]bool)

	var visitAll func(items map[string]bool)

	visitAll = func(items map[string]bool) {

		for key, _ := range items {
			if !seen[key] {
				seen[key] = true
				visitAll(m[key])
				order = append(order, key)
			}

		}
	}

	for key, _ := range m {
		visitAll(m[key])

		if !seen[key] {
			seen[key] = true
			order = append(order, key)
		}
	}

	return order
}

//!-main
