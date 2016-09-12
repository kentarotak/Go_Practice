// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

//!+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {

	var sign byte
	var underDecimal []byte
	//符号と小数点を分離する.
	if strings.HasPrefix(s, "+") || strings.HasPrefix(s, "-") {
		sign = s[0]
		s = s[1:]
	} else {
		sign = 0x20
	}

	if strings.Index(s, ".") != -1 {
		temp := strings.Split(s, ".")
		underDecimal = []byte(temp[1])
		s = temp[0]
	}

	n := len(s) - 1
	if n < 3 {
		if len(underDecimal) >= 1 {
			s = string(sign) + s + "." + string(underDecimal)
		} else {
			s = string(sign) + s + string(underDecimal)
		}
		return strings.TrimSpace(s)
	}

	var buf bytes.Buffer

	for i := 0; i <= n; i++ {
		if i%3 == 0 && i != 0 {
			buf.WriteString(",")
		}
		fmt.Fprintf(&buf, "%c", s[n-i])
	}

	s = Reverse(buf.String())

	if len(underDecimal) >= 1 {
		s = string(sign) + s + "." + string(underDecimal)
	} else {
		s = string(sign) + s + string(underDecimal)
	}

	return strings.TrimSpace(s)
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

//!-
