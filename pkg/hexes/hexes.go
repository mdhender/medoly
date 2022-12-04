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

// Coordinates are x, y with
//
//	x increasing from left right
//	y increasing from top to bottom
type Coordinates struct {
	q int // x is col is q
	r int // y is row is r
}

func NewCoordinates(x, y int) Coordinates {
	return Coordinates{q: x, r: y}
}

// XYToHex assumes
func XYToHex(x, y int) Hex {
	return QOffsetToCube(x, y, ODD)
}

type HEXORIENTATION int

const (
	RADIUS = 30 // radius of a single hex on the board
)

// Axial coordinates are like Cube, but they don't store s.
// We can recover s since s = -q - r.
type Axial struct {
	Q int `json:"q"`
	R int `json:"r"`
}

// Cube coordinates have three axes, q, r, and s
// and the constraint that q + r + s = 0.
type Cube struct {
	Q int `json:"q"`
	R int `json:"r"`
	S int `json:"s"`
}

func (c Cube) ToAxial() Axial {
	var q = c.Q
	var r = c.R
	return Axial{Q: q, R: r}
}

// ToEvenQ shoves even rows to the right
func (c Cube) ToEvenQ() Offset {
	var col = c.Q
	var row = c.R + (c.Q+(c.Q&1))/2
	return Offset{Q: col, R: row}
}

// ToEvenR shoves even columns down
func (c Cube) ToEvenR() Offset {
	var col = c.Q + (c.R+(c.R&1))/2
	var row = c.R
	return Offset{Q: col, R: row}
}

// ToOddQ shoves odd rows to the right
func (c Cube) ToOddQ() Offset {
	var col = c.Q
	var row = c.R + (c.Q-(c.Q&1))/2
	return Offset{Q: col, R: row}
}

// ToOddR shoves odd columns down
func (c Cube) ToOddR() Offset {
	var col = c.Q + (c.R-(c.R&1))/2
	var row = c.R
	return Offset{Q: col, R: row}
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
