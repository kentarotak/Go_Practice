package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"net/url"
	"strconv"
)

var width, height = 600, 320               // canvas size in pixels
var cells = 100                            // number of grid cells
var xyrange = 30.0                         // axis ranges (-xyrange..+xyrange)
var xyscale = float64(width) / 2 / xyrange // pixels per x or y unit
var zscale = float64(height) * 0.4         // pixels per z unit
var angle = math.Pi / 6                    // angle of x, y axes (=30°)
var maxcolor, mincolor = 0xFF0000, 0x0000FF
var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		m, _ := url.ParseQuery(r.URL.RawQuery)
		display(w, m)
	}
	http.HandleFunc("/", handler)
	//!-http
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	return
}

func display(out io.Writer, prm map[string][]string) {

	// パラメータのセット.
	if prm["height"] != nil {
		height, _ = strconv.Atoi(prm["height"][0])
	}
	if prm["width"] != nil {
		width, _ = strconv.Atoi(prm["width"][0])
	}
	if prm["cells"] != nil {
		cells, _ = strconv.Atoi(prm["cells"][0])
	}
	if prm["maxcolor"] != nil && prm["mincolor"] != nil {
		max, _ := strconv.ParseInt(prm["maxcolor"][0], 0, 32)
		min, _ := strconv.ParseInt(prm["mincolor"][0], 0, 32)
		maxcolor, mincolor = int(max), int(min)
	}

	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	// 高さから色を判断するために、sliceで値を保持しておく.
	var ax, ay, bx, by, cx, cy, dx, dy []float64
	var atheight []float64
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

			atheight = append(atheight, tempheight)
		}
	}

	// 色の計算を実施する.
	color := calcColor(atheight)

	for i := 0; i < len(ax); i++ {
		fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' stroke='#%x'/>\n",
			ax[i], ay[i], bx[i], by[i], cx[i], cy[i], dx[i], dy[i], color[i])
	}
	fmt.Fprintf(out, "</svg>")
}

func corner(i, j int) (float64, float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/float64(cells) - 0.5)
	y := xyrange * (float64(j)/float64(cells) - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(width)/2 + (x-y)*cos30*xyscale
	sy := float64(height)/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

func calcColor(height []float64) []int32 {
	maxval := max(height)
	minval := min(height)
	// 最大値,最小値をもとに値の正規化を実施.
	// 色の最大 FF0000, 色の最小 0000FF
	var color []int32
	for i := 0; i < len(height); i++ {
		temp := (float64(maxcolor)-float64(mincolor))*((height[i]-minval)/(maxval-minval)) + float64(mincolor)
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
