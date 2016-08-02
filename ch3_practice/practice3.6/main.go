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
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y1 := float64(py)/height*(ymax-ymin) + ymin
		y2 := (float64(py)+0.5)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x1 := float64(px)/width*(xmax-xmin) + xmin
			x2 := (float64(px)+0.5)/width*(xmax-xmin) + xmin

			z1 := complex(x1, y1)
			z2 := complex(x1, y2)
			z3 := complex(x2, y1)
			z4 := complex(x2, y2)

			r1, g1, b1, a1 := mandelbrot(z1)
			r2, g2, b2, a2 := mandelbrot(z2)
			r3, g3, b3, a3 := mandelbrot(z3)
			r4, g4, b4, a4 := mandelbrot(z4)

			rave := (uint32(r1) + uint32(r2) + uint32(r3) + uint32(r4)) / 4
			gave := (uint32(g1) + uint32(g2) + uint32(g3) + uint32(g4)) / 4
			bave := (uint32(b1) + uint32(b2) + uint32(b3) + uint32(b4)) / 4
			aave := (uint32(a1) + uint32(a2) + uint32(a3) + uint32(a4)) / 4

			// Image point (px, py) represents complex value z.
			img.Set(px, py, color.RGBA{uint8(rave), uint8(gave), uint8(bave), uint8(aave)})
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) (uint8, uint8, uint8, uint8) {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			//return color.Gray{255 - contrast*n}
			return n, 255 * n, n, 255 - contrast*n
		}
	}
	return 0, 0, 0, 0xFF
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
