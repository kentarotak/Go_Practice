package main

import "fmt"

func main() {
	//!+array
	a := [...]int{0, 1, 2, 3, 4, 5}
	rotate(a[:], 2)
	fmt.Println(a)
	//!-array

	// NOTE: ignoring potential errors from input.Err()
}

//!+rev
// reverse reverses a slice of ints in place.
func rotate(s []int, times int) {

	reverse(s[:times])
	reverse(s[times:])
	reverse(s)

}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
