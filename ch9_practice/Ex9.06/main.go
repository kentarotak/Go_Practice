// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// GOMAXPROCS =1 経過時間 10.0661687s
// GOMAXPROCS = 2 経過時間 1.6282489s
// GOMAXPROCS = 3 経過時間 1.2523401s
// GOMAXPROCS = 4 経過時間 957.6622ms
// GOMAXPROCS = 5 経過時間 812.5233ms
// GOMAXPROCS = 6 経過時間 722.512ms
// GOMAXPROCS = 7 経過時間 772.5654ms

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
	"runtime"
	"sync"
	"time"
)

type SetImgVal struct {
	px    int
	py    int
	color color.Color
}

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
)

var wg sync.WaitGroup

var img = image.NewRGBA(image.Rect(0, 0, width, height))

func main() {

	fmt.Fprintf(os.Stderr, "NumCPU=%d\n", runtime.NumCPU())
	runtime.GOMAXPROCS(7)
	start := time.Now()
	ch := make(chan struct{})

	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			wg.Add(1)

			go calPallarel(px, py, y, ch)

		}
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	png.Encode(os.Stdout, img) // NOTE: ignoring errors

	fmt.Fprintf(os.Stderr, "経過時間 %s", time.Since(start))
}

func calPallarel(px int, py int, y float64, ch chan<- struct{}) {
	defer wg.Done()
	x := float64(px)/width*(xmax-xmin) + xmin
	z := complex(x, y)

	//fmt.Fprintf(os.Stderr, "px: %d", px)

	img.Set(px, py, mandelbrot(z))
	/*
		val.px = px
		val.py = py
		val.color = mandelbrot(z)
		ch <- val
	*/
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
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
