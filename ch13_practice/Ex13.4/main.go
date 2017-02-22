package main

import (
	"io"
	"os"

	"github.com/kentarotak/Go_Practice/ch13_practice/Ex13.4/bzipcmd"
)

func main() {
	w := bzip2exe.NewWriter(os.Stdout)

	io.Copy(w, os.Stdin)
}
