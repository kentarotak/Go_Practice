// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package sexpr

import "fmt"

func Example_test() {
	var b bool

	b = true
	data, _ := Marshal(b)
	fmt.Printf("Marshal() = %s\n", data)

	b = false
	data, _ = Marshal(b)
	fmt.Printf("Marshal() = %s\n", data)

	data, _ = Marshal(10)
	fmt.Printf("Marshal() = %s\n", data)

	data, _ = Marshal(10.6)
	fmt.Printf("Marshal() = %s\n", data)

	cmp := complex(3, 5)
	data, _ = Marshal(cmp)
	fmt.Printf("Marshal() = %s\n", data)

	/* Interface??
	var x interface{}
	//x = cmp

	data, _ = Marshal(x)
	fmt.Printf("Marshal() = %s\n", data)
	*/

	// Output:
	// Marshal() = t
	// Marshal() = nil
	// Marshal() = 10
	// Marshal() = 10.600000
	// Marshal() = #C(3.000000 5.000000)
}
