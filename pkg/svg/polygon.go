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

var addCoordinates = false

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

func (p *polygon) Bytes(id string) []byte {
	buf := bytes.Buffer{}
	//buf.WriteString(fmt.Sprintf(`<polygon style="fill: %s; stroke: %s; stroke-width: %s;"`, p.style.fill, p.style.stroke, p.style.strokeWidth))
	buf.WriteString(`<polygon`)
	if id != "" {
		buf.WriteString(fmt.Sprintf(` id="%s"`, id))
	}
	buf.WriteString(fmt.Sprintf(` fill="%s"`, p.style.fill))
	buf.WriteString(fmt.Sprintf(` stroke="%s"`, p.style.stroke))
	buf.WriteString(fmt.Sprintf(` stroke-width="%s"`, p.style.strokeWidth))
	buf.WriteString(fmt.Sprintf(` points="`))

	for i, pt := range p.points {
		if i > 0 {
			buf.WriteByte(' ')
		}
		buf.Write(pt.Bytes())
	}
	buf.WriteString(`"></polygon>`)
	buf.WriteByte('\n')

	if addCoordinates {
		fontSize := 8
		buf.WriteString(fmt.Sprintf(`<text x="%f" y="%f" text-anchor="middle" fill="grey" font-size="%d" font-weight="bold">%s</text>`, p.cx, p.cy, fontSize, fmt.Sprintf("%02d %02d", p.x, p.y)))
		buf.WriteByte('\n')
	}

	return buf.Bytes()
}

func (p *polygon) Use(ref *polygon, id string) []byte {
	buf := bytes.Buffer{}
	dx := p.cx - ref.cx
	dy := p.cy - ref.cy
	buf.WriteString(fmt.Sprintf(`<use href="#%s" x="%f" y="%f" />`, id, dx, dy))
	buf.WriteByte('\n')

	if addCoordinates {
		fontSize := 8
		buf.WriteString(fmt.Sprintf(`<text x="%f" y="%f" text-anchor="middle" fill="grey" font-size="%d" font-weight="bold">%s</text>`, p.cx, p.cy, fontSize, fmt.Sprintf("%02d %02d", p.x, p.y)))
		buf.WriteByte('\n')
	}

	return buf.Bytes()
}

/*
   <circle id="myCircle" cx="200" cy="200" r="4" stroke="blue" />
   <use href="#myCircle" x="10" fill="blue" />
   <use href="#myCircle" x="20" fill="white" stroke="red" />
*/
