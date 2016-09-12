package main

import "testing"

func TestComma(t *testing.T) {

	var tests = []struct {
		testval string
		expect  string
	}{
		{"1", "1"},
		{"12", "12"},
		{"123", "123"},
		{"1234", "1,234"},
		{"+1", "+1"},
		{"+12", "+12"},
		{"+123", "+123"},
		{"+1234", "+1,234"},
		{"-1", "-1"},
		{"-12", "-12"},
		{"-123", "-123"},
		{"-1234", "-1,234"},
		{"+1.01", "+1.01"},
		{"+12.001", "+12.001"},
		{"+123.001", "+123.001"},
		{"+1234.001", "+1,234.001"},
	}

	for _, test := range tests {

		result := comma(test.testval)
		if result != test.expect {
			t.Errorf("err! result = %s, expect = %s\n", result, test.expect)
		}
	}

}
