package main

import (
	"fmt"
	"unicode"
)

func main() {

	str := "世界! 　"

	fmt.Printf("%c\n", []byte(str))

	result := spaceCompress([]byte(str))

	fmt.Printf("%c\n", result)

	fmt.Printf("%s\n", string(result))
}

func spaceCompress(x []byte) []byte {
	var out string
	str := string(x)

	flag := false

	for _, rstr := range str {
		isspace := unicode.IsSpace(rstr)
		fmt.Printf("%d\n", isspace)
		if isspace == true {
			if flag == false {
				out += " "
			}
			flag = true
		} else {
			out += string(rstr)
			flag = false
		}
		fmt.Printf("%d\n", out)
	}

	return []byte(out)
}
