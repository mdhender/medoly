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

package hexes

// Cube coordinates have three axes, q, r, and s
// and the constraint that q + r + s = 0.
type Cube struct {
	Q int `json:"q"`
	R int `json:"r"`
	S int `json:"s"`
}

// Axial coordinates are like Cube, but they don't store s.
// We can recover s since s = -q - r.
type Axial struct {
	Q int `json:"q"`
	R int `json:"r"`
}

// Offset coordinates use col and row.
// In an "odd" layout, odd rows are shoved to the right or down.
// In an "even" layout, even rows are shoved to the right or down.
// Which depends on if the layout is horizontal or vertical.
// The "horizontal" layout shoves rows to the right,
// while the "vertical" layout shoves them down.
type Offset struct {
	Q int `json:"q"` // colum
	R int `json:"r"` // row
}

// Doubled coordinates use col and row.
// They have the constraint that (col + row) % 2 = 0.
type Doubled struct {
	Q int `json:"q"` // colum
	R int `json:"r"` // row
}

func cube_to_axial(c Cube) Axial {
	var q = c.Q
	var r = c.R
	return Axial{Q: q, R: r}
}

func axial_to_cube(a Axial) Cube {
	var q = a.Q
	var r = a.R
	var s = -q - r
	return Cube{Q: q, R: r, S: s}
}

func axial_to_oddr(a Axial) Offset {
	var col = a.Q + (a.R-(a.R&1))/2
	var row = a.R
	return Offset{Q: col, R: row}
}

func oddr_to_axial(o Offset) Axial {
	var q = o.Q - (o.R-(o.R&1))/2
	var r = o.R
	return Axial{Q: q, R: r}
}

func axial_to_evenr(a Axial) Offset {
	var col = a.Q + (a.R+(a.R&1))/2
	var row = a.R
	return Offset{Q: col, R: row}
}

func evenr_to_axial(o Offset) Axial {
	var q = o.Q - (o.R+(o.R&1))/2
	var r = o.R
	return Axial{Q: q, R: r}
}

func axial_to_oddq(a Axial) Offset {
	var col = a.Q
	var row = a.R + (a.Q-(a.Q&1))/2
	return Offset{Q: col, R: row}
}

func oddq_to_axial(o Offset) Axial {
	var q = o.Q
	var r = o.R - (o.Q-(o.Q&1))/2
	return Axial{Q: q, R: r}
}

func axial_to_evenq(a Axial) Offset {
	var col = a.Q
	var row = a.R + (a.Q+(a.Q&1))/2
	return Offset{Q: col, R: row}
}

func evenq_to_axial(o Offset) Axial {
	var q = o.Q
	var r = o.R - (o.Q+(o.Q&1))/2
	return Axial{Q: q, R: r}
}

func cube_to_oddr(c Cube) Offset {
	var col = c.Q + (c.R-(c.R&1))/2
	var row = c.R
	return Offset{Q: col, R: row}
}

func oddr_to_cube(o Offset) Cube {
	var q = o.Q - (o.R-(o.R&1))/2
	var r = o.R
	return Cube{Q: q, R: r, S: -q - r}
}

func cube_to_evenr(c Cube) Offset {
	var col = c.Q + (c.R+(c.R&1))/2
	var row = c.R
	return Offset{Q: col, R: row}
}

func evenr_to_cube(o Offset) Cube {
	var q = o.Q - (o.R+(o.R&1))/2
	var r = o.R
	return Cube{Q: q, R: r, S: -q - r}
}

func cube_to_oddq(c Cube) Offset {
	var col = c.Q
	var row = c.R + (c.Q-(c.Q&1))/2
	return Offset{Q: col, R: row}
}

func oddq_to_cube(o Offset) Cube {
	var q = o.Q
	var r = o.R - (o.Q-(o.Q&1))/2
	return Cube{Q: q, R: r, S: -q - r}
}

func cube_to_evenq(c Cube) Offset {
	var col = c.Q
	var row = c.R + (c.Q+(c.Q&1))/2
	return Offset{Q: col, R: row}
}

func evenq_to_cube(o Offset) Cube {
	var q = o.Q
	var r = o.R - (o.Q+(o.Q&1))/2
	return Cube{Q: q, R: r, S: -q - r}
}

func doubleheight_to_axial(d Doubled) Axial {
	var q = d.Q
	var r = (d.R - d.Q) / 2
	return Axial{Q: q, R: r}
}

func axial_to_doubleheight(a Axial) Doubled {
	var col = a.Q
	var row = 2*a.R + a.Q
	return Doubled{Q: col, R: row}
}

func doublewidth_to_axial(d Doubled) Axial {
	var q = (d.Q - d.R) / 2
	var r = d.R
	return Axial{Q: q, R: r}
}

func axial_to_doublewidth(a Axial) Doubled {
	var col = 2*a.Q + a.R
	var row = a.R
	return Doubled{Q: col, R: row}
}

func doubleheight_to_cube(d Doubled) Cube {
	var q = d.Q
	var r = (d.R - d.Q) / 2
	return Cube{Q: q, R: r, S: -q - r}
}

func cube_to_doubleheight(c Cube) Doubled {
	var col = c.Q
	var row = 2*c.R + c.Q
	return Doubled{Q: col, R: row}
}

func doublewidth_to_cube(d Doubled) Cube {
	var q = (d.Q - d.R) / 2
	var r = d.R
	return Cube{Q: q, R: r, S: -q - r}
}

func cube_to_doublewidth(c Cube) Doubled {
	var col = 2*c.Q + c.R
	var row = c.R
	return Doubled{Q: col, R: row}
}
