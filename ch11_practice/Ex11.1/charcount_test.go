package main

import (
	"bufio"
	"bytes"
	"testing"
	"unicode/utf8"
)

func TestCharcount(t *testing.T) {

	var tests = []struct {
		input     string
		wantcnts  map[string]int
		wantlengs [5]int
		invalid   int
	}{
		{"HelloWorld\n",
			map[string]int{"H": 1, "l": 3, "e": 1, "o": 2, "W": 1, "\n": 1, "r": 1, "d": 1},
			[5]int{0, 11, 0, 0, 0},
			0,
		},
		{"aaaaa\n",
			map[string]int{"a": 5, "\n": 1},
			[5]int{0, 6, 0, 0, 0},
			0,
		},
		// 2バイト文字.
		{"Ā\n",
			map[string]int{"Ā": 1, "\n": 1},
			[5]int{0, 1, 1, 0, 0},
			0,
		},
		// 3バイト文字.
		{"ဣ\n",
			map[string]int{"ဣ": 1, "\n": 1},
			[5]int{0, 1, 0, 1, 0},
			0,
		},
	}

	for _, test := range tests {

		r := bytes.NewReader([]byte(test.input))
		stdin := bufio.NewReader(r)

		counts := make(map[rune]int)    // counts of Unicode characters
		var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
		invalid := 0                    // count of invalid UTF-8 characters

		charcount(stdin, counts, &utflen, &invalid)

		for key, val := range test.wantcnts {
			wantrune := []rune(key)
			if counts[wantrune[0]] != val {
				t.Errorf("str %q:: resunt = %d  want = %d\n", wantrune, counts[wantrune[0]], val)
			}
		}

		for i, val := range test.wantlengs {
			if utflen[i] != val {
				t.Errorf("leng %d:: result = %d  want = %d\n", i, utflen[i], val)
			}
		}

		if test.invalid != invalid {
			t.Errorf("invalid err result = %d  want = %d\n", test.invalid, invalid)
		}

		/*
			for c, n := range counts {
				t.Errorf("%q\t%d\n", c, n)
			}
			fmt.Print("\nlen\tcount\n")
			for i, n := range utflen {
				if i > 0 {
					t.Errorf("%d\t%d\n", i, n)
				}
			}
			if invalid > 0 {
				t.Errorf("\n%d invalid UTF-8 characters\n", invalid)
			}
		*/
	}
}
