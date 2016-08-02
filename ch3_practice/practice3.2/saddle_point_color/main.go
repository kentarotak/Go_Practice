// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.001      // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	// 高さから色を判断するために、sliceで値を保持しておく.
	var ax, ay, bx, by, cx, cy, dx, dy []float64
	var height []float64
	var tempx, tempy, tempheight, tempheight2 float64
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			tempx, tempy, tempheight = corner(i+1, j)
			ax = append(ax, tempx)
			ay = append(ay, tempy)

			tempx, tempy, tempheight2 = corner(i, j)
			if tempheight < tempheight2 {
				tempheight = tempheight2
			}
			bx = append(bx, tempx)
			by = append(by, tempy)

			tempx, tempy, tempheight2 = corner(i, j+1)
			if tempheight < tempheight2 {
				tempheight = tempheight2
			}
			cx = append(cx, tempx)
			cy = append(cy, tempy)

			tempx, tempy, tempheight2 = corner(i+1, j+1)
			if tempheight < tempheight2 {
				tempheight = tempheight2
			}
			dx = append(dx, tempx)
			dy = append(dy, tempy)

			height = append(height, tempheight)
		}
	}

	// 色の計算を実施する.
	color := calcColor(height)

	for i := 0; i < len(ax); i++ {
		fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' stroke='#%x'/>\n",
			ax[i], ay[i], bx[i], by[i], cx[i], cy[i], dx[i], dy[i], color[i])
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func f(x, y float64) float64 {
	r := math.Pow(x, 2) - math.Pow(y, 2)
	return r
}

func calcColor(height []float64) []int32 {
	maxval := max(height)
	minval := min(height)
	// 最大値,最小値をもとに値の正規化を実施.
	// 色の最大 FF0000, 色の最小 0000FF
	var color []int32
	for i := 0; i < len(height); i++ {
		temp := 0xFEFF01*((height[i]-minval)/(maxval-minval)) + 0xFF
		color = append(color, int32(temp))
	}

	return color

}

func max(a []float64) float64 {
	max := a[0]
	for _, i := range a {
		if i > max {
			max = i
		}
	}
	return max
}

func min(a []float64) float64 {
	min := a[0]
	for _, i := range a {
		if i < min {
			min = i
		}
	}
	return min
}

//!-
