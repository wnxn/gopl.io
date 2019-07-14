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
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	maxz, minz := math.Inf(-1), math.Inf(1)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			az, bz, cz, dz := math.Hypot(float64(i+1), float64(j)), math.Hypot(float64(i), float64(j)), math.Hypot(float64(i),
				float64(j+1)), math.Hypot(float64(i+1), float64(j+1))
			maxz = math.Max(math.Max(math.Max(az, bz), math.Max(cz, dz)), maxz)
			minz = math.Min(math.Min(math.Min(az, bz), math.Min(cz, dz)), minz)
		}
	}

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az := corner(i+1, j)
			bx, by, bz := corner(i, j)
			cx, cy, cz := corner(i, j+1)
			dx, dy, dz := corner(i+1, j+1)
			clr := int(math.Max(math.Max(az, bz), math.Max(cz, dz)) / (maxz - minz) * (0xFF0000 - 0x0000FF))
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill:#%x'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, clr)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)
	if z == math.Inf(1) || z == math.Inf(-1) {
		return math.NaN(), math.NaN(), math.NaN()
	}
	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

//!-
