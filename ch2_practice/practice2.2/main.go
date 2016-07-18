package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/kentarotak/Go_Practice/ch2_practice/practice2.2/tempconv"
)

func main() {

	var input []string
	if len(os.Args) > 1 {
		input = os.Args[1:]
	} else {
		//fmt.Scan(input)
		input = []string{"100", "1000", "10", "1"}
	}

	fmt.Printf("%d\n", len(os.Args))
	fmt.Println(input)

	for _, arg := range input {
		t, err := strconv.ParseFloat(arg, 64)

		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Println(f)
		if err != nil {
			fmt.Fprintf(os.Stderr, "conv: %v\n", err)

		}

		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
}
