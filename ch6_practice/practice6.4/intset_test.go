// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package intset

import "testing"

func TestIntersect(t *testing.T) {
	var IntersectTests = []struct {
		testval []int
	}{
		{
			[]int{},
		},
		{
			[]int{1, 2, 3},
		},
		{
			[]int{10, 20, 30, 40, 50, 1000},
		},
	}

	//!+main

	for _, testval := range IntersectTests {
		var x IntSet

		x.AddAll(testval.testval...)

		result := x.Elems()
		// 要素チェック.
		for i, val := range testval.testval {
			if result[i] != val {
				t.Errorf("err! expect %d, result %d\n", val, result[i])
			}
		}

	}

}
