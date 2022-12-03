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

// abs is a helper function to get the absolute value of an integer
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// mod is a helper function to get the modulus of an integer
// (as opposed to %, which is the remainder operator)
func mod(a, b int) int {
	// you can check for b == 0 separately and do what you want
	if b < 0 {
		return -mod(-a, -b)
	}
	m := a % b
	if m < 0 {
		m += b
	}
	return m
}

// hex_direction converts 0...5 to a hex offset
func hex_direction(direction int) Hex {
	// direction = mod(direction, 6)
	direction = (6 + (direction % 6)) % 6
	return hex_directions[direction]
}
