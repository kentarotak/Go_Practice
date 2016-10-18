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
		for _, val := range lentest.testval {
			x.Add(val)
		}
		result := x.Len()

		if result != lentest.result {
			t.Errorf("err! expect %d, result %d\n", lentest.result, result)
		}

	}

}

func TestRemove(t *testing.T) {
	var Lentests = []struct {
		testval   []int
		removeval []int
		result    int
	}{
		{
			[]int{10, 10, 10},
			[]int{10},
			0,
		},
		{
			[]int{},
			[]int{10},
			0,
		},

		{
			[]int{10, 20, 70, 80, 100, 200, 1000, 2000},
			[]int{10, 2000},
			6,
		},
	}

	//!+main

	for _, lentest := range Lentests {
		var x IntSet
		for _, val := range lentest.testval {
			x.Add(val)
		}
		for _, val := range lentest.removeval {
			x.Remove(val)
			// 保持しているか確認.
			chk := x.Has(val)
			if chk == true {
				t.Errorf("err! bag Can't remove: removeval : %d %t\n", val)
			}
		}

		result := x.Len()

		if result != lentest.result {
			t.Errorf("err! expect %d, result %d\n", lentest.result, result)
		}
	}

}

func TestClear(t *testing.T) {
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
			0,
		},
		{
			[]int{10, 20, 70, 80, 100, 200, 1000, 2000},
			0,
		},
	}

	//!+main

	for _, lentest := range Lentests {
		var x IntSet
		for _, val := range lentest.testval {
			x.Add(val)
		}
		x.Clear()
		result := x.Len()
		if result != lentest.result {
			t.Errorf("err! expect %d, result %d\n", lentest.result, result)
		}

	}

}

func TestCopy(t *testing.T) {
	var Lentests = []struct {
		testval []int
	}{
		{
			[]int{},
		},
		{
			[]int{10, 10, 10},
		},
		{
			[]int{10, 20, 70, 80, 100, 200, 1000, 2000},
		},
	}

	//!+main

	for _, lentest := range Lentests {
		var x IntSet
		for _, val := range lentest.testval {
			x.Add(val)
		}
		y := x.Copy()
		base := x.Len()
		copy := y.Len()

		if base != copy {
			t.Errorf("err! because of different size. base %d, copy %d\n", base, copy)
		}

		for _, val := range lentest.testval {
			if x.Has(val) != y.Has(val) {
				t.Errorf("err! element error %d\n", val)
			}
		}

	}

}
