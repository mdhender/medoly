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

package svg

import (
	"bytes"
	"fmt"
	"github.com/mdhender/medoly/pkg/hexes"
	"github.com/mdhender/medoly/pkg/terrain"
	"math"
)

const (
	EDGE   = 20
	EDGES  = EDGE * 2
	RADIUS = 30
)

type SVG struct {
	id       string
	width    float64
	height   float64
	viewBox  viewBox
	layout   hexes.Layout
	polygons []*polygon
}

type Hex struct {
	cube    hexes.Cube
	terrain terrain.Terrain
}

func New(cols, rows int) *SVG {
	offset := (math.Sqrt(3) * RADIUS) / 2
	// cols -> width  -> x
	maxX := EDGES + offset*float64(cols*2)
	// rows -> height -> y
	maxY := EDGES + offset*float64(rows)*math.Sqrt(3)

	s := &SVG{
		id:     "s",
		width:  2.0 * RADIUS,
		height: math.Sqrt(3.0) * RADIUS,
		viewBox: viewBox{
			minX:   0,
			minY:   0,
			width:  int(maxX),
			height: int(maxY),
		},
		//polygons: make(map[int]*polygon),
	}

	// assumes flat with even-q layout
	s.layout = hexes.NewFlatLayout(hexes.NewPoint(RADIUS, RADIUS), hexes.NewPoint(s.height, s.width))

	return s
}

func (s *SVG) AddHex(x, y int, t terrain.Terrain) {
	poly := &polygon{
		x:       x,
		y:       y,
		radius:  s.height / 2.0,
		terrain: t,
	}
	h := hexes.XYToHex(x, y)
	poly.cx, poly.cy = s.layout.CenterPoint(h).Coords()

	poly.style.stroke = "Grey"
	poly.style.fill = t.ToFill()
	if poly.style.fill == poly.style.stroke {
		poly.style.stroke = "Black"
		if poly.style.fill == poly.style.stroke {
			poly.style.stroke = "White"
		}
	}
	poly.style.strokeWidth = "2px"

	for _, p := range s.layout.PolygonCorners(h) {
		px, py := p.Coords()
		if width := int(px); width+EDGES > s.viewBox.width {
			s.viewBox.width = width + EDGES
		}
		if height := int(py); height+EDGES > s.viewBox.height {
			s.viewBox.height = height + EDGES
		}
		poly.points = append(poly.points, point{x: px, y: py})
	}

	s.polygons = append(s.polygons, poly)
}

func (s *SVG) Bytes() []byte {
	buf := bytes.Buffer{}
	buf.WriteString("<svg")
	if s.id != "" {
		buf.WriteString(fmt.Sprintf(" id=%q", s.id))
	}
	buf.WriteString(fmt.Sprintf(` width="%d" height="%d"`, s.viewBox.width+EDGES, s.viewBox.height+EDGES))
	buf.Write(s.viewBox.Bytes())
	buf.Write([]byte(` xmlns="http://www.w3.org/2000/svg">`))
	buf.WriteByte('\n')
	buf.WriteString(fmt.Sprintf(`<rect height="%d" width="%d" style="fill: Grey; stroke: Black; stroke-width: 2px;" />`,
		s.viewBox.height, s.viewBox.width))
	buf.WriteByte('\n')
	buf.WriteString(fmt.Sprintf(`<rect x="%d" y="%d" height="%d" width="%d" style="fill: %s; stroke: %s; stroke-width: 2px;" />`,
		EDGE, EDGE, s.viewBox.height-EDGES, s.viewBox.width-EDGES, terrain.Ocean.ToFill(), terrain.Ocean.ToFill()))
	buf.WriteByte('\n')

	for _, poly := range s.polygons {
		//log.Printf("poyl y %3d x %3d\n", poly.y, poly.x)
		buf.Write(poly.Bytes())
		buf.WriteByte('\n')
	}

	buf.Write([]byte("</svg>"))

	return buf.Bytes()
}

type point struct {
	x, y float64
}

func (p point) Bytes() []byte {
	return []byte(fmt.Sprintf("%f,%f", p.x, p.y))
}

type viewBox struct {
	minX, minY    int
	width, height int
}

func (v viewBox) Bytes() []byte {
	return []byte(fmt.Sprintf(` viewBox="%d %d %d %d"`, v.minX, v.minY, v.width+EDGES, v.height+EDGES))
}
