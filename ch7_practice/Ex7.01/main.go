package main

import (
	"bufio"
	"fmt"
	"strings"
)

type WordLineCounter struct {
	Line int
	Word int
}

func (c *WordLineCounter) ScanWord(data []byte, atEOF bool) (advance int, token []byte, err error) {
	advance, token, err = bufio.ScanWords(data, atEOF)

	if atEOF == true {
		c.Line++
	} else {
		c.Word++
	}
	return advance, token, err
}

func main() {
	const input = "1234 5678 1234567901234567890\n aaa bb cc"
	scanner := bufio.NewScanner(strings.NewReader(input))

	var c WordLineCounter

	scanner.Split(c.ScanWord)

	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}

	fmt.Printf("%d\n", c.Word)
	fmt.Printf("%d\n", c.Line)

}
