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

func TestPopcountByShifting(t *testing.T) {
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
		if got := PopCountByShifting(test.testval); got != test.expect {
			t.Errorf("err %v = %v", test.expect, got)
		}
	}
}

func TestPopcountByClearing(t *testing.T) {
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
		if got := PopCountByClearing(test.testval); got != test.expect {
			t.Errorf("err %v = %v", test.expect, got)
		}
	}
}

const (
	//testval = 0xFFFFFFF00000000 // 56
	testval = 0xFF0000000000000 // 16
	//testval = 0xF00000000000000 // 8
	//testval = 0x000000000000000 // 0
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(testval)
	}
}

func BenchmarkPopCountByLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByLoop(testval)
	}
}

func BenchmarkPopCountByShifting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByShifting(testval)
	}
}

func BenchmarkPopCountByClearing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByClearing(testval)
	}
}

func BenchmarkPopCountAllF(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(0xFFFFFFFFFFFFFFFF)
	}
}

func BenchmarkPopCountByLoopAllF(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByLoop(0xFFFFFFFFFFFFFFFF)
	}
}

func BenchmarkPopCountByShiftingAllF(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByShifting(0xFFFFFFFFFFFFFFFF)
	}
}

func BenchmarkPopCountByClearingAllF(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByClearing(0xFFFFFFFFFFFFFFFF)
	}
}
