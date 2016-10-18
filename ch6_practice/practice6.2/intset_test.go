// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package intset

import "testing"

func TestLen(t *testing.T) {
	var Lentests = []struct {
		testval []int
		result  int
	}{
		{
			[]int{},
			0,
		},
		{
			[]int{10, 10, 10},
			1,
		},
		{
			[]int{10, 20, 70, 80, 100, 200, 1000, 2000},
			8,
		},
	}

	//!+main

	for _, lentest := range Lentests {
		var x IntSet

		x.AddAll(lentest.testval...)
		result := x.Len()

		if result != lentest.result {
			t.Errorf("err! expect %d, result %d\n", lentest.result, result)
		}

	}

}
