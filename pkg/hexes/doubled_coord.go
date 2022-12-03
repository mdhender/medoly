// medokh - a clone of a much better game
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

type DoubledCoord struct {
	col, row int
}

func NewDoubledCoord(col, row int) DoubledCoord {
	return DoubledCoord{col: col, row: row}
}

func (a DoubledCoord) Equals(b DoubledCoord) bool {
	return a.col == b.col && a.row == b.row
}

func qdoubled_from_cube(h Hex) DoubledCoord {
	col := h.q
	row := 2*h.r + h.q

	return NewDoubledCoord(col, row)
}

func qdoubled_to_cube(h DoubledCoord) Hex {
	q := h.col
	r := (h.row - h.col) / 2
	s := -q - r

	return NewHex(q, r, s)
}

func rdoubled_from_cube(h Hex) DoubledCoord {
	col := 2*h.q + h.r
	row := h.r

	return NewDoubledCoord(col, row)
}

func rdoubled_to_cube(h DoubledCoord) Hex {
	q := (h.col - h.row) / 2
	r := h.row
	s := -q - r

	return NewHex(q, r, s)
}
