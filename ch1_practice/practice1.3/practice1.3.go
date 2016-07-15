package main

import (
	"fmt"
	"os"
	"strings"
)

//!+
func main() {
	poorLogic(os.Args[1:])
	normalLogic(os.Args[1:])
}

func poorLogic(args []string) {
	var s, sep string
	for i := 0; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	fmt.Println(s)
}

func normalLogic(args []string) {
	fmt.Println(strings.Join(args, " "))
}

//!-
