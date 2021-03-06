// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"net/http"
	"log"
	"io"
	"net/url"
	"strconv"
	"fmt"
)

func main() {

	handler := func(w http.ResponseWriter, r *http.Request) {
		m, _ := url.ParseQuery(r.URL.RawQuery)
		display(w , m)
	}
	http.HandleFunc("/", handler)
	//!-http
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	return


}

func display(out io.Writer,prm map[string][]string) {

	a	:= 0
	b	:= 0

	var scale float64
	scale = 100


	if prm["x"] != nil {
		a, _ = strconv.Atoi(prm["x"][0])
	}
	if prm["y"] != nil {
		b, _ = strconv.Atoi(prm["y"][0])
	}
	if prm["scale"] != nil {
		temp, _ := strconv.Atoi(prm["scale"][0])
		scale = float64(temp)/100
		scale = (4/scale)/2
	}

	fmt.Printf("scale %f\n",scale)

	var xmin,ymin,xmax,ymax float64

	xmin = -float64(scale) + float64(a)
	ymin = -float64(scale) + float64(b)
	xmax = float64(scale)  +  float64(a)
	ymax = float64(scale)  +  float64(b)

	/*
	xmin =  xmin/float64(scale)
	ymin =  ymin/float64(scale)
	xmax =  xmax/float64(scale)
	ymax =  ymax/float64(scale)
	*/

	width, height          := 1024, 1024

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}

	png.Encode(out, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			//return color.Gray{255 - contrast*n}
			return color.RGBA{n, 255 * n, n, 255 - contrast*n}
		}
	}
	return color.Black
}

//!-

// Some other interesting functions:

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{128, blue, red}
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//    = z - (z^4 - 1) / (4 * z^3)
//    = z - (z - 1/z^3) / 4
func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}
