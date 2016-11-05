// "$foo $bar $xxx -> "FOO BAR XXX"
// になるようにする.

// $を検出して、定義したfunctionに引数を渡して変換させる.

package main

import (
	"fmt"
	"strings"
)

func main() {

	f := func(data string) string {
		return strings.ToUpper(data)
	}

	fmt.Printf("%s\n", expand("$foo $bar xxx", f))

}

func expand(s string, f func(string) string) string {

	sep := strings.Split(s, " ")

	var temp []string
	for _, i := range sep {
		if strings.HasPrefix(i, "$") {
			temp = append(temp, f(i))
		} else {
			temp = append(temp, i)
		}
	}

	result := strings.Join(temp, " ")
	return result
}
