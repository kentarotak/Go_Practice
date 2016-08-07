package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	file := os.Args[1:]

	f, err := os.Open(file[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "file not exsist\n")
		os.Exit(1)
	}
	wordfreq(f, counts)

}

func wordfreq(f *os.File, counts map[string]int) {

	input := bufio.NewScanner(f)
	input.Split(bufio.ScanWords)

	size := 0
	for input.Scan() {
		counts[input.Text()]++
		size++
	}

	fmt.Printf("単語数 %d\n", size)

	for line, n := range counts {
		freq := int((float64(n) / float64(size)) * 100)
		fmt.Printf("%s\t%d%%\n", line, freq)
	}

}
