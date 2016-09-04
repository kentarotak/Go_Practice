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
	"math"
	"math/big"
	"os"

	"github.com/barnex/fmath"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin

			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(x, y))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(x float64, y float64) color.Color {
	const iterations = 50
	const contrast = 15
	const prec = 2048
	re := big.NewFloat(x)
	im := big.NewFloat(y)

	vre := big.NewFloat(0)
	vim := big.NewFloat(0)
	temp1 := big.NewFloat(0)
	temp2 := big.NewFloat(0)

	abs := big.NewFloat(0)

	for n := uint8(0); n < iterations; n++ {
		temp1 = new(big.Float).SetPrec(prec).Mul(vre, vre)
		temp2 = new(big.Float).SetPrec(prec).Mul(vim, vim)
		vre = new(big.Float).Sub(temp1, temp2)

		temp1 = new(big.Float).SetPrec(prec).Mul(vre, vim)
		temp2 = new(big.Float).SetPrec(prec).Mul(vre, vim)
		vim = new(big.Float).SetPrec(prec).Add(temp1, temp2)

		vre = new(big.Float).SetPrec(prec).Add(vre, re)
		vim = new(big.Float).SetPrec(prec).Add(vim, im)

		// 絶対値の2乗.
		temp1 = new(big.Float).SetPrec(prec).Mul(vre, vre)
		temp2 = new(big.Float).SetPrec(prec).Mul(vim, vim)
		abs = new(big.Float).SetPrec(prec).Add(temp1, temp2)
		res, _ := abs.Float64()
		res = math.Sqrt(res)
		if res > 2 {
			return color.Gray{255 - contrast*n}
		}
		//v = v*v + z
		//if Abs(v) > 2 {
		//	return color.Gray{255 - contrast*n}
		//}
	}
	return color.Black
}

func Abs(v complex64) float32 {

	p := real(v)
	q := imag(v)

	if p < 0 {
		p = -p
	}
	if q < 0 {
		q = -q
	}
	if p < q {
		p, q = q, p
	}
	if p == 0 {
		return 0
	}
	q = q / p
	return p * fmath.Sqrtf(1+q*q)
}
