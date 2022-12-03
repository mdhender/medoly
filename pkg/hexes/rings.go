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

func cube_scale(c Cube, factor int) Cube {
	return Cube{Q: c.Q * factor, R: c.R * factor, S: c.S * factor}
}

// this code doesn't work for radius == 0; can you see why?
func cube_ring(center Cube, radius int) []Cube {
	var results []Cube
	var hex = cube_add(center, cube_scale(cube_direction(4), radius))
	for i := 0; i < 6; i++ {
		for j := 0; j < radius; j++ {
			results = append(results, hex)
			hex = cube_neighbor(hex, i)
		}
	}
	return results
}

// returns number of cubes in the radius
func Area(radius int) int {
	if radius < 0 {
		return 0
	}
	return 1 + 3*radius*(radius+1)
}
