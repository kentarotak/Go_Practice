// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package word

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

//!+bench

//!-bench

//!+test
func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
		{"detartrated", true},
		{"A man, a plan, a canal: Panama", true},
		{"Evil I did dwell; lewd did I live.", true},
		{"Able was I ere I saw Elba", true},
		{"été", true},
		{"Et se resservir, ivresse reste.", true},
		{"palindrome", false}, // non-palindrome
		{"desserts", false},   // semi-palindrome
	}
	for _, test := range tests {
		if got := IsPalindrome(test.input); got != test.want {
			t.Errorf("IsPalindrome(%q) = %v", test.input, got)
		}
	}
}

//!-test

//!+bench
func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome("A man, a plan, a canal: Panama")
	}
}

//!-bench

//!+example

func ExampleIsPalindrome() {
	fmt.Println(IsPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(IsPalindrome("palindrome"))
	// Output:
	// true
	// false
}

//!-example

/*
//!+random
import "math/rand"

//!-random
*/

//!+random
// randomPalindrome returns a palindrome whose length and contents
// are derived from the pseudo-random number generator rng.
func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // random length up to 24
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000)) // random rune up to '\u0999'
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

func TestRandomPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}

func TestNotRandomPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := notRandomPalindrome(rng)
		if IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = true", p)
		}
	}
}

// 回文ではない文字列を生成する.
func notRandomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) + 5 // random between 5 from 30

	runes := make([]rune, n)
	for i := 0; i < n; i++ {
		r := rune(rng.Intn(0x1000)) // random rune up to '\u0999'
		runes[i] = r
	}

	// 必ず回文にならないように最初と最後の値を異なる値にする.
	runes[0] = 0x041   // A
	runes[n-1] = 0x042 // B
	return string(runes)
}

func TestRandomPalindromesIncludeSpaceAndPunctuation(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomPalindromeIncludeSpaceAndPunctuation(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = true", p)
		}
	}
}

// 句読点や空白を必ず含む回文を生成する.
func randomPalindromeIncludeSpaceAndPunctuation(rng *rand.Rand) string {
	n := rng.Intn(25) // random length up to 24
	m := rng.Intn(9)

	l := n + m

	runes := make([]rune, l)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000)) // random rune up to '\u0999'
		runes[i] = r
		runes[n-1-i] = r
	}

	insertlist := []rune{0x020, 0x0A0, 0x2002, 0x2003, 0x2004, 0x2005, 0x2006, 0x2007, 0x2008, 0x2009, 0x200A, 0x200B, 0x3000, 0xFEFF, 0x0009}

	// 句読点やスペースをm領域文挿入する.
	for i := 0; i < m; i++ {
		insert := rng.Intn(n + m)
		copy(runes[insert+1:], runes[insert:])

		runes[insert] = insertlist[rng.Intn(len(insertlist))]
	}

	return string(runes)
}

//!-random

// Answer for Exercicse 11.1: Modify randomPalindrome to exercise
// IsPalindrome's handling of punctuation and spaces.

// WARNING: the conversion r -> upper -> lower doesn't preserve
// the value of r in some cases, e.g., µ Μ, ſ S, ı I

// randomPalindrome returns a palindrome whose length and contents
// are derived from the pseudo-random number generator rng.
/*
func randomNoisyPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // random length up to 24
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x200)) // random rune up to \u99
		runes[i] = r
		r1 := r
		if unicode.IsLetter(r) && unicode.IsLower(r) {
			r = unicode.ToUpper(r)
			if unicode.ToLower(r) != r1 {
				fmt.Printf("cap? %c %c\n", r1, r)
			}
		}
		runes[n-1-i] = r
	}
	return "?5gigu" + string(runes) + "!"
}

func TestRandomNoisyPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	n := 0
	for i := 0; i < 1000; i++ {
		p := randomNoisyPalindrome(rng)
		if IsPalindrome(p) {
			t.Errorf("IsNoisyPalindrome(%q) = false", p)
			n++
		}
	}
	fmt.Fprintf(os.Stderr, "fail = %d\n", n)
}
*/
