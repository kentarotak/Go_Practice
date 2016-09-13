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
// 題意は満たせてないかも...
func reverse(s []byte) {

	// runeにキャスト.
	tmp := []rune(string(s))
	for i, j := 0, len(tmp)-1; i < j; i, j = i+1, j-1 {
		tmp[i], tmp[j] = tmp[j], tmp[i]
	}

	tmp2 := []byte(string(tmp))

	fmt.Printf("%s\n", string(tmp))

	for i := 0; i < len(tmp2); i++ {
		s[i] = tmp2[i]
	}
	//s = []byte(string(tmp))

	fmt.Printf("%s\n", string(s))

}

func reverse2(s []byte) {

	for i, j := 0, len(s)-1; i < j; {
		_, size := utf8.DecodeRuneInString(string(s[i:]))
		fmt.Printf("%d\n", size)

		i += size
		j -= size
	}

}

//!-rev
