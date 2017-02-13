// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package sexpr

import (
	"io"
	"strings"
	"testing"
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
func Test(t *testing.T) {
	type Movie struct {
		Title, Subtitle string
	}
	expect := []Movie{
		{"Title1", "Subtitle1"},
		{"Title2", "Subtitle2"},
		{"Title3", "Subtitle3"},
		{"Title4", "Subtitle4"},
	}

	const sexprStream = `((Title "Title1") (Subtitle "Subtitle1"))
		((Title "Title2") (Subtitle "Subtitle2"))
		((Title "Title3") (Subtitle "Subtitle3"))
		((Title "Title4") (Subtitle "Subtitle4"))
	`
	// Decode it

	dec := NewDecoder(strings.NewReader(sexprStream))

	for i := 0; ; i++ {
		var m Movie
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			t.Logf("%s\n", err)
		}
		if m.Title != expect[i].Title || m.Subtitle != expect[i].Subtitle {
			t.Errorf("expect %s %s : result %s %s", expect[i].Title, expect[i].Subtitle, m.Title, m.Subtitle)
		}

	}

}
