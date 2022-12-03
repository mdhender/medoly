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

// Orientation stores the forward matrix (the fN variables) and backward matrix
// (the bN variables), plus the start angle. The start angle determines if we
// have a "flat top" (which is 0°) or "pointy top" (which is 60°) hex.
type Orientation struct {
	f0, f1, f2, f3 float64
	b0, b1, b2, b3 float64
	// The starting angle should be 0.0 for 0° (flat top) or 0.5 for 60° (pointy top).
	start_angle float64 // in multiples of 60°
}

// NewOrientation returns an initialized Orientation.
func NewOrientation(f0, f1, f2, f3, b0, b1, b2, b3 float64, hexLayout HEXLAYOUT) Orientation {
	switch hexLayout {
	case FLATHEX:
		return Orientation{
			f0: f0, f1: f1, f2: f2, f3: f3,
			b0: b0, b1: b1, b2: b2, b3: b3,
			start_angle: 0.0,
		}
	case POINTYHEX:
		return Orientation{
			f0: f0, f1: f1, f2: f2, f3: f3,
			b0: b0, b1: b1, b2: b2, b3: b3,
			start_angle: 0.5,
		}
	}
	panic("assert(layoutType in (FLAT, POINTY))")
}
