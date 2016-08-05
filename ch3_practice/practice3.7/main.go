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
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			var ctype int
			if x < 0 && y < 0 {
				ctype = 1
			} else if x >= 0 && y < 0 {
				ctype = 2
			} else if x >= 0 && y >= 0 {
				ctype = 3
			} else if x < 0 && y >= 0 {
				ctype = 4
			}

			img.Set(px, py, newton(z, ctype))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func newton(z complex128, ctype int) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			if ctype == 1 {
				return color.RGBA{i, 255 * i, i, 255 - contrast*i}
			} else if ctype == 2 {
				return color.RGBA{255 * i, i, i, 255 - contrast*i}
			} else if ctype == 3 {
				return color.RGBA{i, i, 255 * i, 255 - contrast*i}
			} else if ctype == 4 {
				return color.RGBA{i, 255 * i, 255 * i, 255 - contrast*i}
			}
		}
	}
	return color.Black
}
