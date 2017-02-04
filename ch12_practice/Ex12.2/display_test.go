// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package display

// NOTE: we can't use !+..!- comments to excerpt these tests
// into the book because it defeats the Example mechanism,
// which requires the // Output comment to be at the end
// of the function.

func Example_expr() {
	type Cycle struct {
		Value int
		Tail  *Cycle
	}

	var c Cycle
	c = Cycle{42, &c}

	Display("c", c)

	// Output:
	//
}
