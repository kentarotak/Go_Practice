package main

import "testing"

func TestDup(t *testing.T) {

	var tests = []struct {
		testval []string
		expect  []string
	}{
		{[]string{"one"}, []string{"one"}},
		{[]string{"one", "two"}, []string{"one", "two"}},
		{[]string{"one", "two", "three"}, []string{"one", "two", "three"}},
		{[]string{"one", "one", "three"}, []string{"one", "three"}},
		{[]string{"one", "one", "three", "one"}, []string{"one", "three", "one"}},
		{[]string{"one", "one", "one"}, []string{"one"}},
		{[]string{"one", "one", "one", "one"}, []string{"one"}},
		{[]string{"one", "one", "one", "one", "two"}, []string{"one", "two"}},
		{[]string{"one", "one", "one", "one", "two", "two", "two"}, []string{"one", "two"}},
	}

	for _, test := range tests {
		if got := removeDuplicateWord(test.testval); equal(test.expect, got) != true {
			t.Errorf("err %q = %q", test.expect, got)
		}
	}
}

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}

	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
