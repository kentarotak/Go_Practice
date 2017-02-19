// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package sexpr

import (
	"fmt"
	"sort"
)

// Test verifies that encoding and decoding a complex data value
// produces an equal result.
//
// The test does not make direct assertions about the encoded output
// because the output depends on map iteration order, which is
// nondeterministic.  The output of the t.Log statements can be
// inspected by running the test with the -v flag:
//
// 	$ go test -v gopl.io/ch12/sexpr
//
func Example_Test() {
	type Movie struct {
		Title    string            `sexpr:"title"`
		Subtitle string            `sexpr:"subtitle"`
		Year     int               `sexpr:"year"`
		Actor    map[string]string `sexpr:"actor"`
		Oscars   []string          `sexpr:"oscars"`
		Sequel   *string           `sexpr:"sequel"`
	}
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}

	// Encode it
	data, err := Marshal(strangelove)

	if err != nil {
		fmt.Printf("err : %v\n", err)
		return
	}

	fmt.Printf("%s\n", data)

	test := new(Movie)

	Unmarshal(data, test)

	fmt.Printf("Unmarshal Title: %s\n", test.Title)
	fmt.Printf("Unmarshal Subtitle: %s\n", test.Subtitle)
	fmt.Printf("Unmarshal Year: %d\n", test.Year)

	var sortmapkey []string
	for key, _ := range test.Actor {
		sortmapkey = append(sortmapkey, key)
	}

	sort.Strings(sortmapkey)

	for _, key := range sortmapkey {
		fmt.Printf("Unmarshal Actor: %s\n", test.Actor[key])
	}

	for _, val := range test.Oscars {
		fmt.Printf("Unmarshal Oscars: %s\n", val)
	}

	// output:
	//Mapが順不同なので..
}
