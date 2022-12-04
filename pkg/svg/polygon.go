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
	"github.com/mdhender/medoly/pkg/terrain"
)

// polygon is the actual hex on the board
type polygon struct {
	x, y    int
	terrain terrain.Terrain // terrain type of the hex

	cx, cy, radius float64 // center of the hex
	points         []point

	style struct {
		fill        string
		stroke      string
		strokeWidth string
	}

	addCircle bool
	text      []string
}

func (p *polygon) Bytes() []byte {
	buf := bytes.Buffer{}
	buf.WriteString(fmt.Sprintf(`<polygon style="fill: %s; stroke: %s; stroke-width: %s;"`, p.style.fill, p.style.stroke, p.style.strokeWidth))
	buf.WriteString(fmt.Sprintf(` points="`))
	for i, pt := range p.points {
		if i > 0 {
			buf.WriteByte(' ')
		}
		buf.Write(pt.Bytes())
	}
	buf.WriteString(`"></polygon>\n`)
	buf.WriteByte('\n')

	//fontSize := 14
	//s += fmt.Sprintf(`<text x="%f" y="%f" text-anchor="middle" fill="grey" font-size="%d" font-weight="bold">%s</text>`, p.cx, p.cy, fontSize, fmt.Sprintf("%02d%02d", p.x+1, p.y+1))

	return buf.Bytes()
}
