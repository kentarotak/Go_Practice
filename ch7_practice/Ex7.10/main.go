package main

import "sort"

func IsPalindrome(s sort.Interface) bool {

	length := s.Len()

	for i := 0; i < length; i++ {
		if !(!s.Less(i, length-i-1) && !s.Less(length-i-1, i)) {
			return false
		}

	}

	return true

}
