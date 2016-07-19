package main

import "testing"

func TestPopcount(t *testing.T) {
	var tests = []struct {
		testval uint64
		expect  int
	}{
		{0x00, 0},
		{0x01, 1},
		{0x02, 1},
		{0x85362, 8},
	}

	for _, test := range tests {
		if got := PopCount(test.testval); got != test.expect {
			t.Errorf("err %v = %v", test.expect, got)
		}
	}
}

func TestPopcountByLoop(t *testing.T) {
	var tests = []struct {
		testval uint64
		expect  int
	}{
		{0x00, 0},
		{0x01, 1},
		{0x02, 1},
		{0x85362, 8},
	}

	for _, test := range tests {
		if got := PopCountByLoop(test.testval); got != test.expect {
			t.Errorf("err %v = %v", test.expect, got)
		}
	}
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByLoop(0x1234567890ABCDEF)
	}
}
