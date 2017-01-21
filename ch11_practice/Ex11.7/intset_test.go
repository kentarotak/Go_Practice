// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package intset

import (
	"math/rand"
	"testing"
	"time"
)

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

const (
	//length = 0xFFFFFFFF

	//BenchmarkIntSetBench-4        	 5000000	       243 ns/op
	//BenchmarkMapBench-4           	 5000000	       266 ns/op
	//BenchmarkInsetUnionWithBench-4	       2	 610791550 ns/op
	//BenchmarkMapUnionWithBench-4  	   10000	    362959 ns/op

	length = 0xFF

//BenchmarkIntSetBench-4        	50000000	        32.4 ns/op
//BenchmarkMapBench-4           	20000000	        74.8 ns/op
//BenchmarkInsetUnionWithBench-4	20000000	        71.3 ns/op
//BenchmarkMapUnionWithBench-4  	  100000	     15755 ns/op

// 組み込みマップは桁数が大きいと優位
// 桁数が少ないとIntsetのほうが優位
)

func BenchmarkIntSetBench(b *testing.B) {
	//!+main

	var iset IntSet
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < b.N; i++ {
		n := rng.Intn(length)
		iset.Add(n)
	}

}

func BenchmarkMapBench(b *testing.B) {
	//!+main

	mset := make(map[int]bool)
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < b.N; i++ {
		n := rng.Intn(length)
		mset[n] = true
	}

}

func BenchmarkInsetUnionWithBench(b *testing.B) {
	//!+main

	var xset IntSet
	var yset IntSet

	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < b.N; i++ {
		m := rng.Intn(length)
		n := rng.Intn(length)
		xset.Add(m)
		yset.Add(n)

		xset.UnionWith(&yset)
	}

}

func BenchmarkMapUnionWithBench(b *testing.B) {
	//!+main

	xmap := make(MapSet)
	ymap := make(MapSet)

	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < b.N; i++ {
		m := rng.Intn(length)
		n := rng.Intn(length)
		xmap[m] = true
		ymap[n] = true
		xmap.UnionWith(ymap)
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
