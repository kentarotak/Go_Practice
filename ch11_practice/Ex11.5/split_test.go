package split

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	var tests = []struct {
		str     string
		sep     string
		wantlen int
	}{
		{"a:b:c", ":", 3},
		{"1,2,3,4", ",", 4},
		{",,,,", ",", 5},
	}

	for _, test := range tests {
		words := strings.Split(test.str, test.sep)

		if got, want := len(words), test.wantlen; got != want {
			t.Errorf("Split(%q, &q) returned %d words, want %d\n", test.str, test.sep, got, want)
		}
	}
}
