// medoly - a clone of much better games
// Copyright (c) 2022 Michael D Henderson
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package td

import (
	"bytes"
	"fmt"
	"math"
	"math/rand"
)

// From The Go Programming Language book

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange .. +xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit (the 0.4 is arbitrary)
	angle         = math.Pi / 6         // angle of x, y axes (=30degrees)
)

var (
	sin30 = math.Sin(angle) // sin of 30degrees
	cos30 = math.Cos(angle) // cos of 30degrees
)

func Shatter(cycles int) []byte {
	var land [cells + 2][cells + 2]float64 // land[x][y] = z, which is height

	buf := &bytes.Buffer{}
	//for w, h := cells/2, cells/2; cycles > 0 && w > 0 && h > 0; cycles-- {
	//	w, h = w/2, h/2
	//	land[w][h] = float64(rand.Intn(cycles))
	//}

	for cycles > 0 {
		for i := 0; i < cells; i++ {
			for j := 0; j < cells; j++ {
				if rand.Intn(2) == 1 {
					land[i][j] = float64(rand.Intn(cycles + 2))
				}
			}
		}
		cycles--
	}

	buf.WriteString(fmt.Sprintf("<svg style='stroke:grey; fill:none; stroke-width:0.7' width='%d' height='%d' xmlns='http://www.w3.org/2000/svg'>\n", width, height))
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			// ax, ay := float64(i+1), float64(j)
			// bx, by := float64(i), float64(j)
			// cx, cy := float64(i), float64(j+1)
			// dx, dy := float64(i+1), float64(j+1)

			ax, ay := project(float64(i+1), float64(j), land[i+1][j])
			bx, by := project(float64(i), float64(j), land[i][j])
			cx, cy := project(float64(i), float64(j+1), land[i][j+1])
			dx, dy := project(float64(i+1), float64(j+1), land[i+1][j+1])

			buf.WriteString(fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n", ax, ay, bx, by, cx, cy, dx, dy))
		}
	}
	buf.WriteString("</svg>")

	return buf.Bytes()
}

func TD() []byte {
	buf := &bytes.Buffer{}

	var coords [4][2]float64

	//buf.WriteString(fmt.Sprintf("<svg style='stroke:grey; fill:white; stroke-width:0.7' width='%d' height='%d' xmlns='http://www.w3.org/2000/svg'>\n", width, height))
	buf.WriteString("<svg style='stroke:grey; fill:white; stroke-width:0.7' xmlns='http://www.w3.org/2000/svg'>\n")
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			coords[0][0], coords[0][1] = corner(i+1, j)
			if math.IsInf(coords[0][0], 0) || math.IsNaN(coords[0][0]) {
				continue
			} else if math.IsInf(coords[0][1], 0) || math.IsNaN(coords[0][1]) {
				continue
			}

			coords[1][0], coords[1][1] = corner(i, j)
			if math.IsInf(coords[1][1], 0) || math.IsNaN(coords[1][1]) {
				continue
			} else if math.IsInf(coords[1][0], 0) || math.IsNaN(coords[1][0]) {
				continue
			}

			coords[2][0], coords[2][1] = corner(i, j+1)
			if math.IsInf(coords[2][0], 0) || math.IsNaN(coords[2][0]) {
				continue
			} else if math.IsInf(coords[2][1], 0) || math.IsNaN(coords[2][1]) {
				continue
			}

			coords[3][0], coords[3][1] = corner(i+1, j+1)
			if math.IsInf(coords[3][0], 0) || math.IsNaN(coords[3][0]) {
				continue
			} else if math.IsInf(coords[3][1], 0) || math.IsNaN(coords[3][1]) {
				continue
			}

			buf.WriteString("<polygon points='")
			for _, coord := range coords {
				buf.WriteString(fmt.Sprintf("%g,%g ", coord[0], coord[1]))
				if coord[0] > 1800 || coord[1] > 1800 {
					panic(fmt.Sprintf("assert(%g, %g)", coord[0], coord[1]))
				}
			}
			buf.WriteString("'/>\n")
		}
	}
	buf.WriteString("</svg>")
	return buf.Bytes()
}

func corner(i, j int) (sx, sy float64) {
	// find point (x,y) at corner of cell (i,j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// compute surface height z
	z := f(x, y)

	// project (x,y,z) isometrically onto 2-D SVG canvas (sx, sy)
	sx = width/2 + (x-y)*cos30*xyscale
	sy = height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

// f returns the z for a given x and y
func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from the origin
	return math.Sin(r) / r
}

// find point (x,y) at corner of cell (i,j)
func ijToXy(i, j int) (x, y float64) {
	x = xyrange * (float64(i)/cells - 0.5)
	y = xyrange * (float64(j)/cells - 0.5)
	return x, y
}

func project(x, y, z float64) (sx, sy float64) {
	z *= 0.2
	// project (x,y,z) isometrically onto 2-D SVG canvas (sx, sy)
	sx = width/2 + (x-y)*cos30*xyscale
	sy = height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}
