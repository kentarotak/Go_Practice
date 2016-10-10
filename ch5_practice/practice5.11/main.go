// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 136.

// The toposort program prints the nodes of a DAG in topological order.
package main

import (
	"fmt"
	"sort"
)

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
	"linear algebra":        {"calculus"},
}

//!-table

const (
	unvisited = iota
	visiting
	visited
)

//!+main
func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]int)
	var visitAll func(items []string)
	var visit func(item string)

	visitAll = func(items []string) {
		for _, item := range items {
			visit(item)
		}
	}

	visit = func(item string) {
		seen[item] = visiting
		for _, node := range m[item] {
			if seen[node] == unvisited {
				visit(node)
			} else if seen[node] == visiting {
				fmt.Printf("★★循環発生!! %sから%sに向かいましたが、\n別のルートからすでに訪れています★★\n", item, node)
			}
		}
		seen[item] = visited
		order = append(order, item)
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	visitAll(keys)
	return order
}

//!-main
