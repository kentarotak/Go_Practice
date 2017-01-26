package main

import (
	"fmt"
	"os"

	"github.com/kentarotak/Go_Practice/ch10_practice/Ex10.2/archivelib"
)

func main() {
	files := os.Args[1:]
	dir, _ := os.Getwd()
	err := archivelib.Decompress(files[0], dir)

	fmt.Printf("%s\n", err)
}
