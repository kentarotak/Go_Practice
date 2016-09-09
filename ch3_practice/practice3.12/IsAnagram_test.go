package main

import "testing"

func TestIsAnagram(t *testing.T) {

	var tests = []struct {
		testval1 string
		testval2 string
		result   bool
	}{
		{"abc", "cba", true},
		{"", "", true},
		{"10_", "01_", true},
		{"abcd", "abc", false},
		{"abc", "abcd", false},
	}

	for _, test := range tests {

		result := IsAnagram(test.testval1, test.testval2)
		if result != test.result {
			t.Errorf("err! s1 =%s, s2 = %s, expect %t, result %t\n", test.testval1, test.testval2, test.result, result)
		}
	}

}
