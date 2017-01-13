// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package intset

import "testing"

func TestIntSetandMap(t *testing.T) {
	//!+main

	var iset IntSet
	mset := make(map[int]bool)

	tests := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, test := range tests {
		iset.Add(test)
		mset[test] = true

		if iset.Has(test) != mset[test] {
			t.Errorf("str %d:: resunt = %q  want = %q\n", iset.Has(test), mset[test])
		}

	}

}

type MapSet map[int]bool

func (s MapSet) UnionWith(t MapSet) {

	for key, val := range t {
		s[key] = val
	}

}

func TestUnionWith(t *testing.T) {
	var x, y IntSet
	xmap := make(MapSet)
	ymap := make(MapSet)

	testsets1 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	testsets2 := [10]int{20, 21, 22, 23, 24, 25, 26, 27, 28, 29}

	for _, test := range testsets1 {
		x.Add(test)
		xmap[test] = true
	}

	for _, test := range testsets2 {
		y.Add(test)
		ymap[test] = true
	}

	x.UnionWith(&y)
	xmap.UnionWith(ymap)

	for i := 0; i < 30; i++ {
		if x.Has(i) != xmap[i] {
			t.Errorf("str %d:: resunt = %q  want = %q\n", x.Has(i), xmap[i])
		}
	}

}
