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

var cube_direction_vectors = []Cube{
	Cube{Q: +1, R: 0, S: -1},
	Cube{Q: +1, R: -1, S: 0},
	Cube{Q: 0, R: -1, S: +1},
	Cube{Q: -1, R: 0, S: +1},
	Cube{Q: -1, R: +1, S: 0},
	Cube{Q: 0, R: +1, S: -1},
}

func cube_direction(direction int) Cube {
	direction = (6 + (direction % 6)) % 6 // mod(direction, 6)
	return cube_direction_vectors[direction]
}

func cube_add(c Cube, vec Cube) Cube {
	return Cube{c.Q + vec.Q, c.R + vec.R, c.S + vec.S}
}

func cube_neighbor(c Cube, direction int) Cube {
	return cube_add(c, cube_direction(direction))
}

var cube_diagonal_vectors = []Cube{
	Cube{Q: +2, R: -1, S: -1},
	Cube{Q: +1, R: -2, S: +1},
	Cube{Q: -1, R: -1, S: +2},
	Cube{Q: -2, R: +1, S: +1},
	Cube{Q: -1, R: +2, S: -1},
	Cube{Q: +1, R: +1, S: -2},
}

func cube_diagonal_neighbor(c Cube, direction int) Cube {
	return cube_add(c, cube_diagonal_vectors[direction])
}
