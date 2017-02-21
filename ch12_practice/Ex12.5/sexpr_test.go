// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package sexpr

import (
	"encoding/json"
	"fmt"
	"log"
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
		Title, Subtitle string
		Year            int
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
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
		log.Fatalf("Marshal failed: %v", err)
	}

	jsonBytes := ([]byte)(data)
	result := new(Movie)

	if err := json.Unmarshal(jsonBytes, result); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return
	}

	fmt.Println(result.Title)
	fmt.Println(result.Subtitle)
	fmt.Println(result.Year)

	var sortmapkey []string
	for key, _ := range result.Actor {
		sortmapkey = append(sortmapkey, key)
	}

	sort.Strings(sortmapkey)

	for _, key := range sortmapkey {
		fmt.Printf("%s : %s\n", key, result.Actor[key])
	}

	for _, val := range result.Oscars {
		fmt.Printf("%s\n", val)
	}

	// Output:
	//Dr. Strangelove
	//How I Learned to Stop Worrying and Love the Bomb
	//1964
	//Brig. Gen. Jack D. Ripper : Sterling Hayden
	//Dr. Strangelove : Peter Sellers
	//Gen. Buck Turgidson : George C. Scott
	//Grp. Capt. Lionel Mandrake : Peter Sellers
	//Maj. T.J. "King" Kong : Slim Pickens
	//Pres. Merkin Muffley : Peter Sellers
	//Best Actor (Nomin.)
	//Best Adapted Screenplay (Nomin.)
	//Best Director (Nomin.)
	//Best Picture (Nomin.)
}