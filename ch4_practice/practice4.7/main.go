// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 86.

// Rev reverses a slice.
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//!+array
	a := "Hello! 世界 "
	testdata := []byte(a)
	reverse(testdata[:])
	fmt.Printf("%s\n", testdata[:]) // "[5 4 3 2 1 0]"
	//!-array

	// NOTE: ignoring potential errors from input.Err()
}

//!+rev
func reverse(s []byte) {

	for i := 0; i < len(s); {
		_, size := utf8.DecodeRune(s[i:])
		reverse2(s[i : i+size])
		i += size
	}

	reverse2(s[:])

}

func reverse2(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

//!-rev
