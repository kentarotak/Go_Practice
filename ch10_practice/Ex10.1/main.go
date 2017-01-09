// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 287.

//!+main

// The jpeg command reads a PNG image from the standard input
// and writes it as a JPEG image to the standard output.
package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"io"
	"os"
)

func main() {
	var outformat string
	var inputfile string
	flag.StringVar(&outformat, "out", "blank", "string flag")
	flag.StringVar(&inputfile, "file", "blank", "string flag")
	flag.Parse()

	file, err := os.Open(inputfile)

	defer file.Close()

	if err != nil {
		fmt.Println("ファイルが存在しません\n")
		return
	}

	if err := toOtherFormat(file, os.Stdout, outformat); err != nil {
		fmt.Fprintf(os.Stderr, "err %v\n", err)
		os.Exit(1)
	}
}

func toOtherFormat(in io.Reader, out io.Writer, convform string) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	switch convform {
	case "png":
		err = png.Encode(out, img)
	case "jpeg":
		err = jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
	case "gif":
		err = gif.Encode(out, img, nil)
	}

	return err
}

//!-main

/*
//!+with
$ go build gopl.io/ch3/mandelbrot
$ go build gopl.io/ch10/jpeg
$ ./mandelbrot | ./jpeg >mandelbrot.jpg
Input format = png
//!-with

//!+without
$ go build gopl.io/ch10/jpeg
$ ./mandelbrot | ./jpeg >mandelbrot.jpg
jpeg: image: unknown format
//!-without
*/
