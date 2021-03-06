package weightconv

import (
	"math"
	"testing"
)

func TestPtoG(t *testing.T) {
	var tests = []struct {
		testval Pound
		expect  Gram
	}{
		{0, 0},
		{1000, 453592},
	}

	for _, test := range tests {
		if got := PtoG(test.testval); got != test.expect {
			t.Errorf("MToFTest(%q) = %v", test.expect, got)
		}
	}
}

func TestGtoP(t *testing.T) {
	var tests = []struct {
		testval Gram
		expect  Pound
	}{
		{0, 0},
		{1000, 2},
	}

	for _, test := range tests {
		if got := GtoP(test.testval); Pound(math.Trunc(float64(got))) != test.expect {
			t.Errorf("MToFTest(%q) = %v", test.expect, got)
		}
	}
}
