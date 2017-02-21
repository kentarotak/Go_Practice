// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package sexpr

import "fmt"

func Example_test() {
	var b bool

	b = true

	data, err := Marshal(b)
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Printf("Marshal() = %s\n", data)

	data, err = Marshal(10)
	fmt.Printf("Marshal() = %s\n", data)

	data, err = Marshal(10.6)
	fmt.Printf("Marshal() = %s\n", data)

	cmp := complex(3, 5)
	data, err = Marshal(cmp)
	fmt.Printf("Marshal() = %s\n", data)

	// Output:
	// Marshal() = t
	// Marshal() = 10
	// Marshal() = 10.600000
	// Marshal() = #C(3.000000 5.000000)
}

func Example_Unmarshal_test() {
	var b bool

	b = true

	data, err := Marshal(b)
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Printf("Marshal() = %s\n", data)

	data, err = Marshal(10)
	fmt.Printf("Marshal() = %s\n", data)

	data, err = Marshal(10.6)
	fmt.Printf("Marshal() = %s\n", data)
	var fl float64
	Unmarshal(data, &fl)
	fmt.Printf("Unmarshal() = %v\n", fl)

	cmp := complex(3, 5)
	data, err = Marshal(cmp)
	fmt.Printf("Marshal() = %s\n", data)
	var cmp2 complex128
	Unmarshal(data, &cmp2)
	fmt.Printf("Unmarshal() = %v\n", cmp2)

	// Output:
	// Marshal() = t
	// Marshal() = 10
	// Marshal() = 10.600000
	// Marshal() = #C(3.000000 5.000000)
}
