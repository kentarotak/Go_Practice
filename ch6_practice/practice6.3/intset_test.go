// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package intset

import (
	"fmt"
	"testing"
)

func TestIntersect(t *testing.T) {
	var IntersectTests = []struct {
		testval    []int
		testval2   []int
		result_val []int
		result_len int
	}{
		{
			[]int{},
			[]int{},
			[]int{},
			0,
		},
		{
			[]int{1, 2, 3},
			[]int{},
			[]int{},
			0,
		},
		{
			[]int{},
			[]int{1, 2, 3},
			[]int{},
			0,
		},

		{
			[]int{10, 20, 30, 40, 50, 1000},
			[]int{40, 50, 60, 70, 80, 2000},
			[]int{40, 50},
			2,
		},
	}

	//!+main

	for _, testval := range IntersectTests {
		var x IntSet
		var y IntSet

		x.AddAll(testval.testval...)
		y.AddAll(testval.testval2...)

		x.IntersectWith(&y)
		fmt.Printf("目視確認 : IntersetWith %s\n", x.String())

		// 要素のサイズ数を確認.
		if x.Len() != testval.result_len {
			t.Errorf("err! size different:: result %d, expect %d\n", x.Len(), testval.result_len)
		}

		// 要素チェック.
		for _, exp := range testval.result_val {
			if x.Has(exp) != true {
				t.Errorf("err! must have this element : %d\n", exp)
			}
		}

	}

}

func TestSDifference(t *testing.T) {
	var SDifferenceTests = []struct {
		testval    []int
		testval2   []int
		result_val []int
		result_len int
	}{
		{
			[]int{},
			[]int{},
			[]int{},
			0,
		},
		{
			[]int{1, 2, 3},
			[]int{},
			[]int{1, 2, 3},
			3,
		},
		{
			[]int{},
			[]int{1, 2, 3},
			[]int{1, 2, 3},
			3,
		},
		{
			[]int{10, 20, 30, 40, 50, 1000},
			[]int{40, 50, 60, 70, 80, 2000, 3000},
			[]int{10, 20, 30, 60, 70, 80, 1000, 2000, 3000},
			9,
		},
		{
			[]int{40, 50, 60, 70, 80, 2000, 3000},
			[]int{10, 20, 30, 40, 50, 1000},
			[]int{10, 20, 30, 60, 70, 80, 1000, 2000, 3000},
			9,
		},
	}

	//!+main

	for _, testval := range SDifferenceTests {
		var x IntSet
		var y IntSet

		x.AddAll(testval.testval...)
		y.AddAll(testval.testval2...)

		x.SymmetricDifference(&y)
		fmt.Printf("目視確認 : SymmetricDifference %s\n", x.String())

		// 要素のサイズ数を確認.
		if x.Len() != testval.result_len {
			t.Errorf("err! size different:: result %d, expect %d\n", x.Len(), testval.result_len)
		}

		// 要素チェック.
		for _, exp := range testval.result_val {
			if x.Has(exp) != true {
				t.Errorf("err! must have this element : %d\n", exp)
			}
		}

	}

}

func TestDifference(t *testing.T) {
	var DifferenceTests = []struct {
		testval    []int
		testval2   []int
		result_val []int
		result_len int
	}{
		{
			[]int{},
			[]int{},
			[]int{},
			0,
		},
		{
			[]int{1, 2, 3},
			[]int{},
			[]int{1, 2, 3},
			3,
		},
		{
			[]int{},
			[]int{1, 2, 3},
			[]int{},
			0,
		},
		{
			[]int{10, 20, 30, 40, 50, 1000},
			[]int{40, 50, 60, 70, 80, 2000},
			[]int{10, 20, 30, 1000},
			4,
		},
		{
			[]int{40, 50, 60, 70, 80, 2000},
			[]int{10, 20, 30, 40, 50, 1000},
			[]int{60, 70, 80, 2000},
			4,
		},
	}

	//!+main

	for _, testval := range DifferenceTests {
		var x IntSet
		var y IntSet

		x.AddAll(testval.testval...)
		y.AddAll(testval.testval2...)

		x.DifferenceWith(&y)
		fmt.Printf("目視確認 : DifferenceWith %s\n", x.String())

		// 要素のサイズ数を確認.
		if x.Len() != testval.result_len {
			t.Errorf("err! size different:: result %d, expect %d\n", x.Len(), testval.result_len)
		}

		// 要素チェック.
		for _, exp := range testval.result_val {
			if x.Has(exp) != true {
				t.Errorf("err! must have this element : %d\n", exp)
			}
		}

	}

}
