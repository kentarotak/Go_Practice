package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/kentarotak/Go_Practice/ch2_practice/practice2.2/lengthconv"
	"github.com/kentarotak/Go_Practice/ch2_practice/practice2.2/tempconv"
	"github.com/kentarotak/Go_Practice/ch2_practice/practice2.2/weightconv"
)

func main() {

	var input []string
	if len(os.Args) > 1 {
		input = os.Args[1:]
	} else {
		temp := bufio.NewScanner(os.Stdin)
		temp.Scan()
		input = strings.Split(temp.Text(), " ")
	}

	for _, arg := range input {
		t, err := strconv.ParseFloat(arg, 64)

		if err != nil {
			fmt.Fprintf(os.Stderr, "conv: %v\n", err)
			os.Exit(1)
		}

		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("---temp -----\n")
		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))

		feet := lengthconv.Feet(t)
		meter := lengthconv.Meter(t)
		fmt.Printf("---- length -----\n")
		fmt.Printf("%s = %s, %s = %s\n", feet, lengthconv.FtoM(feet), meter, lengthconv.MtoF(meter))

		pound := weightconv.Pound(t)
		gram := weightconv.Gram(t)
		fmt.Printf("---- weight -----\n")
		fmt.Printf("%s = %s, %s = %s\n", pound, weightconv.PtoG(pound), gram, weightconv.GtoP(gram))

	}
}
