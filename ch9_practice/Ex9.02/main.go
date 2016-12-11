// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 45.

// (Package doc comment intentionally malformed to demonstrate golint.)
//!+
package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

var once sync.Once

// pc[i] is the population count of i.
var pc [256]byte

func loadCounter() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	once.Do(func() { loadCounter() })

	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])

}

func PopCountByLoop(x uint64) int {
	var temp int
	for i := 0; i < 8; i++ {
		temp += int(pc[byte(x>>(uint(i)*8))])
	}
	return temp
}

func main() {
	str := os.Args[1]
	num, _ := strconv.Atoi(str)

	fmt.Printf("popcount = %d, PopcountLoop = %d\n", PopCount(uint64(num)), PopCountByLoop(uint64(num)))
}

//!-
