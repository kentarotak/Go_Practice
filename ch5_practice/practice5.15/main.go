// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 142.

// The sum program demonstrates a variadic function.
package main

import "fmt"

func min(vals ...int) int {
	result := 0
	if len(vals) == 0 {
		return 0
	} else {
		result = vals[0]
	}

	for _, val := range vals {
		if result > val {
			result = val
		}
	}
	return result
}

func max(vals ...int) int {
	result := 0
	if len(vals) == 0 {
		return 0
	} else {
		result = vals[0]
	}

	for _, val := range vals {
		if result < val {
			result = val
		}
	}
	return result
}

//!-

func main() {
	//!+main
	fmt.Println(min())           //  "0"
	fmt.Println(min(3))          //  "3"
	fmt.Println(min(1, 2, 3, 4)) //  "10"
	//!-main

	fmt.Println(max())           //  "0"
	fmt.Println(max(3))          //  "3"
	fmt.Println(max(1, 2, 3, 4)) //  "10"

	//!+slice
	values := []int{1, 2, 3, 4}
	fmt.Println(min(values...)) // "10"
	//!-slice
}
