package main

import (
	"sort"
	"testing"
)

func TestPalindrome(t *testing.T) {
	var PalindromeTests = []struct {
		testval sort.IntSlice
		result  bool
	}{
		{
			sort.IntSlice{1, 2, 3, 4, 3, 2, 1}, true,
		},
		{
			sort.IntSlice{1, 2, 3, 3, 2, 1}, true,
		},
		{
			sort.IntSlice{1, 2, 3}, false,
		},
		{
			sort.IntSlice{10, 20, 30, 40, 50, 1000}, false,
		},
	}

	for _, test := range PalindromeTests {
		if IsPalindrome(test.testval) != test.result {
			t.Errorf("err! %v", test.testval)
		}
	}

}
