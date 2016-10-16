package main

import "fmt"

func main() {
	fmt.Printf("%d\n", returnValueNonReturnSentence())
}

func returnValueNonReturnSentence() (value int) {
	defer func() {
		recover()
		value = 5
	}()
	panic("Panic!")
}
