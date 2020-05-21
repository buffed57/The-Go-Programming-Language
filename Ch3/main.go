// This package contains the solutions to Exercise 3.2
package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

type f func(x, y float64) float64

// Main Creates a web server that displays the SVG
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		svg(w)
	})
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func svg(out io.Writer) {
	var t f
	if len(os.Args) != 1 {
		switch os.Args[1] {
		case "eggBox":
			t = eggBox
		case "saddle":
			t = saddle
		case "normal":
			t = normal
		}
	} else {
		return
	}

	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	fmt.Fprintln(out,"")
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az := corner(i+1, j, t)
			bx, by, bz := corner(i, j, t)
			cx, cy, cz := corner(i, j+1, t)
			dx, dy, dz := corner(i+1, j+1, t)
			averageY := (az + bz + cz + dz) / 4
			fill := "#0000ff"
			if averageY > 0 {
				fill = "#ff0000"
			}

			if math.IsInf(ay, 0) != true && math.IsInf(by, 0) != true && math.IsInf(cy, 0) != true && math.IsInf(dy, 0) != true {
				fmt.Fprintf(out, "<polygon fill=\"%v\" points='%g,%g %g,%g %g,%g %g,%g'/>\n",
					fill,ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
	}
	fmt.Fprintln(out, "</svg>")
}

func corner(i, j int, t f) (float64, float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := t(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func normal(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

func eggBox(x, y float64) float64 {
	return 0.2 * (math.Cos(x) + math.Cos(y))
}

func saddle(x, y float64) float64 {
	a := 25.0
	b := 17.0
	a2 := a * a
	b2 := b * b
	return (y*y/a2 - x*x/b2)
}
