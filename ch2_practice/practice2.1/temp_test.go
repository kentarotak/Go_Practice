package tempconv

import (
	"math"
	"testing"
)

var epsilon float64

func TestCToF(t *testing.T) {
	var tests = []struct {
		testval Celsius
		expect  Fahrenheit
	}{
		{-17.2, 1.00},
		{-14.4, 6.07},
		{0, 32.00},
	}

	for _, test := range tests {
		got := CToF(test.testval)
		epsilon = math.Abs(float64(got - test.expect))
		if epsilon > 0.1 {
			t.Errorf("CToFTest(%q) = %v", test.expect, got)
		}
	}
}

func TestFToC(t *testing.T) {
	var tests = []struct {
		testval Fahrenheit
		expect  Celsius
	}{
		{1.0, -17.2},
		{6.0, -14.4},
		{32.0, 0},
	}

	for _, test := range tests {
		got := FToC(test.testval)
		epsilon = math.Abs(float64(got - test.expect))
		if epsilon > 0.1 {
			t.Errorf("FToCTest(%q) = %v", test.expect, got)
		}
	}
}

func TestKToC(t *testing.T) {
	var tests = []struct {
		testval Kelvin
		expect  Celsius
	}{
		{0, -273.15},
		{373.15, 100},
	}

	for _, test := range tests {
		if got := KToC(test.testval); got != test.expect {
			t.Errorf("KToCTest(%q) = %v", test.expect, got)
		}
	}
}

func TestCToK(t *testing.T) {
	var tests = []struct {
		testval Celsius
		expect  Kelvin
	}{
		{-273.15, 0},
		{100, 373.15},
	}

	for _, test := range tests {
		if got := CToK(test.testval); got != test.expect {
			t.Errorf("CToKTest(%q) = %v", test.expect, got)
		}
	}
}

func TestKToF(t *testing.T) {
	var tests = []struct {
		testval Kelvin
		expect  Fahrenheit
	}{
		{273.15, 32},
	}

	for _, test := range tests {
		if got := KToF(test.testval); got != test.expect {
			t.Errorf("KToFTest(%q) = %v", test.expect, got)
		}
	}
}

func TestFToK(t *testing.T) {
	var tests = []struct {
		testval Fahrenheit
		expect  Kelvin
	}{
		{32, 273.15},
	}

	for _, test := range tests {
		if got := FToK(test.testval); got != test.expect {
			t.Errorf("KToFTest(%q) = %v", test.expect, got)
		}
	}
}
