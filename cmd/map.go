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

package cmd

import (
	"bytes"
	"fmt"
	"github.com/mdhender/medoly/pkg/board"
	"github.com/mdhender/medoly/pkg/hexes"
	"github.com/mdhender/medoly/pkg/terrain"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// cmdMap runs the map command
var cmdMap = &cobra.Command{
	Use:   "map",
	Short: "map things",
	Run: func(cmd *cobra.Command, args []string) {
		cols, rows := argsMap.cols, argsMap.rows
		log.Printf("map: generating %4d x %4d map\n", cols, rows)

		//if err := terrain.SampleCorpus("sample-corpus.json"); err != nil {
		//	log.Fatal(err)
		//}

		mc, err := terrain.LoadCorpus("corpus.json")
		if err != nil {
			log.Fatal(err)
		}
		if mc == nil {
			log.Fatal(fmt.Errorf("assert(mc != nil)"))
		}

		b := board.New(cols, rows)

		// seed the map
		// northern and southern rows are impassable ice
		for _, y := range []int{0, rows - 1} {
			for x := 0; x < cols; x++ {
				if b.IsSet(x, y) {
					continue
				}
				b.SetTerrain(x, y, terrain.Ice)
			}
		}
		// western and eastern columns are impassable oceans
		for _, x := range []int{0, cols - 1} {
			for y := 0; y < rows; y++ {
				if b.IsSet(x, y) {
					continue
				}
				b.SetTerrain(x, y, terrain.Ocean)
			}
		}
		// center of the map is the mountain of the gods
		mx, my := (cols+(cols&1))/2, (rows+(rows&1))/2
		b.SetTerrain(mx, my, terrain.SacredMountain)
		for _, dir := range []int{0, 1, 2, 3, 4, 5} {
			x, y := hexes.Neighbor(mx, my, dir)
			if b.IsSet(x, y) {
				continue
			}
			b.SetTerrain(x, y, terrain.Mountain)
		}
		// randomize the remainder of the map
		for y := 0; y < rows; y++ {
			for x := 0; x < cols; x++ {
				if b.IsSet(x, y) {
					continue
				}
				var t1, t2 terrain.Terrain
				for dir := 0; dir < 6; dir++ {
					if nx, ny := hexes.Neighbor(x, y, dir); b.IsSet(nx, ny) {
						if t1 == terrain.Clear {
							t1 = b.GetTerrain(nx, ny)
						} else if t2 == terrain.Clear {
							t2 = b.GetTerrain(nx, ny)
							break
						}
					}
				}
				if t1 == terrain.Clear {
					t1 = terrain.Rock
				}
				if t2 == terrain.Clear {
					t2 = terrain.Rough
				}
				b.SetTerrain(x, y, mc.Next(t1, t2))
			}
		}

		svg := b.AsSVG(argsMap.addCoordinates)
		//if err := os.WriteFile("medoly.svg", svg, 0666); err != nil {
		//	log.Fatal(err)
		//}
		//log.Printf("created %q", "medoly.svg")
		buf := &bytes.Buffer{}
		buf.WriteString(`<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title>Medoly World Map</title></head><body>`)
		buf.WriteByte('\n')
		buf.Write(svg)
		buf.WriteByte('\n')
		buf.WriteString(`</body></html>`)
		if err := os.WriteFile("medoly.html", buf.Bytes(), 0666); err != nil {
			log.Fatal(err)
		}
		log.Printf("created %q", "medoly.html")
	},
}

var argsMap struct {
	rows           int
	cols           int
	addCoordinates bool
}

func init() {
	cmdRoot.AddCommand(cmdMap)

	cmdMap.Flags().IntVar(&argsMap.cols, "cols", 40, "number of columns to generate")
	cmdMap.Flags().IntVar(&argsMap.rows, "rows", 40, "number of rows to generate")
	cmdMap.Flags().BoolVar(&argsMap.addCoordinates, "add-coordinates", false, "add coordinates to the map")
}
