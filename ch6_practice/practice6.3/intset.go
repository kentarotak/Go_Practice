// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 165.

// Package intset provides a set of integers based on a bit vector.
package intset

import (
	"bytes"
	"fmt"
)

//!+intset

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)

	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) Len() int {
	var count int
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				count++
			}
		}
	}

	return count
}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)

	//Addされていない場合は、何もせずに終了.
	if word >= len(s.words) {
		return
	}
	s.words[word] &= ^(1 << bit)

}

func (s *IntSet) Clear() {
	s.words = nil
}

func (s *IntSet) Copy() *IntSet {
	var y IntSet
	y.words = make([]uint64, len(s.words), cap(s.words))

	copy(y.words, s.words)

	return &y
}

func (s *IntSet) AddAll(xs ...int) {
	for _, x := range xs {
		word, bit := x/64, uint(x%64)

		for word >= len(s.words) {
			s.words = append(s.words, 0)
		}
		s.words[word] |= 1 << bit
	}

}

func (s *IntSet) IntersectWith(t *IntSet) {

	if len(s.words) == 0 || len(t.words) == 0 {
		s.words = nil
		t.words = nil
		return
	}

	var tmp, tmp2 *IntSet

	// 大小関係を見て、ループさせる要素を決める.
	if len(s.words) > len(t.words) {
		tmp, tmp2 = s, t
	} else {
		tmp, tmp2 = t, s
	}

	// 共通して含まれる要素を抽出するので要素の値が少ない方で回す
	for i := range tmp2.words {
		tmp2.words[i] &= tmp.words[i]
	}

	s.words = tmp2.words
}

func (s *IntSet) SymmetricDifference(t *IntSet) {

	if len(s.words) == 0 && len(t.words) == 0 {
		// 両方の要素が空集合の場合は何もせずに終了.
		return
	} else if len(s.words) == 0 {
		// sの要素が0の場合は、tの値をsとして終了.
		s.words = t.words
		return
	} else if len(t.words) == 0 {
		// tの要素が0の場合はそのまま終了.
		return
	}

	var tmp, tmp2 *IntSet

	// 大小関係を見て、ループさせる要素を決める.
	if len(s.words) > len(t.words) {
		tmp, tmp2 = s, t
	} else {
		tmp, tmp2 = t, s
	}

	// 要素が小さい方でループを回して共通する要素を探す.
	for i := range tmp2.words {
		// 共通する要素を抜き出して
		tmp3 := tmp2.words[i] & tmp.words[i]
		// 論理和を取って.
		tmp.words[i] |= tmp2.words[i]
		// 共通する要素と排他的論理和を取る.
		tmp.words[i] ^= tmp3
	}

	s.words = tmp.words
}

func (s *IntSet) DifferenceWith(t *IntSet) {

	if len(s.words) == 0 || len(t.words) == 0 {
		return
	}

	// 大小関係を見て、ループさせる要素を決める.
	var size []uint64
	if len(s.words) > len(t.words) {
		size = t.words
	} else {
		size = s.words
	}

	for i := range size {
		tmp := s.words[i] & t.words[i]
		s.words[i] ^= tmp
	}

}

//!-intset

//!+string

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

//!-string
