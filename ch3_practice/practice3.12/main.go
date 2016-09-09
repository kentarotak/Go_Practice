package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	result := IsAnagram("abc", "bca")

	fmt.Printf("%d\n", result)
}

func IsAnagram(str1 string, str2 string) bool {

	s1 := strings.Split(str1, "")
	s2 := strings.Split(str2, "")

	sort.Strings(s1)
	sort.Strings(s2)

	if len(s1) != len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}
