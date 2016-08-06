package main

import "fmt"

// nonempty returns a slice holding only the non-empty strings.
// The underlying array is modified during the call.
func removeDuplicateWord(strings []string) []string {

	if len(strings) <= 1 {
		return strings
	}
	var flag bool
	out := strings[:0]
	for i := 0; i < len(strings)-1; i++ {
		if strings[i] != strings[i+1] {
			out = append(out, strings[i])
			if len(strings)-2 == i {
				out = append(out, strings[i+1])
			}
			flag = false
		} else {
			if len(strings)-2 == i && flag == true {
				out = append(out, strings[i+1])
			}
			flag = true
		}
	}
	return out
}

//!-nonempty

func main() {
	//!+main
	data := []string{"one", "one", "get", "get", "get", "three", "get"}
	fmt.Printf("%q\n", removeDuplicateWord(data)) // `["one" "three"]`
	fmt.Printf("%q\n", data)                      // `["one" "three" "three"]`
	//!-main
}
