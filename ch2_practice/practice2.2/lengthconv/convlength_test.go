package lengthconv

import (
	"math"
	"testing"
)

func TestFtoM(t *testing.T) {
	var tests = []struct {
		testval Feet
		expect  Meter
	}{
		{0, 0},
		{36, 10},
	}

	for _, test := range tests {
		if got := FtoM(test.testval); Meter(math.Trunc(float64(got))) != test.expect {
			t.Errorf("MToFTest(%q) = %v", test.expect, got)
		}
	}
}

func TestMtoF(t *testing.T) {
	var tests = []struct {
		testval Meter
		expect  Feet
	}{
		{0, 0},
		{11, 36},
	}

	for _, test := range tests {
		if got := MtoF(test.testval); Feet(math.Trunc(float64(got))) != test.expect {
			t.Errorf("FToMTest(%q) = %v", test.expect, got)
		}
	}
}
