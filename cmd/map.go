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
	"fmt"
	"github.com/mdhender/medoly/pkg/board"
	"github.com/mdhender/medoly/pkg/terrain"
	"github.com/spf13/cobra"
	"log"
	"math/rand"
	"os"
)

// cmdMap runs the map command
var cmdMap = &cobra.Command{
	Use:   "map",
	Short: "map things",
	Run: func(cmd *cobra.Command, args []string) {
		mc, err := terrain.LoadCorpus("corpus.json")
		if err != nil {
			log.Fatal(err)
		}
		if mc == nil {
			log.Fatal(fmt.Errorf("assert(mc != nil)"))
		}

		cols, rows := argsMap.cols, argsMap.rows
		b := board.New(cols, rows)

		// initialize the map
		for y := 0; y < rows; y++ {
			for x := 0; x < cols; x++ {
				// northern and southern rows are impassable ice
				if y == 0 || y == rows-1 {
					b.SetTerrain(x, y, terrain.Ice)
					continue
				}
				// center of the map is the mountain of the gods
				if x == cols/2 && y == rows/2 {
					b.SetTerrain(x, y, terrain.Mountain)
					continue
				}
				b.SetTerrain(x, y, mc.Next(terrain.Clear, terrain.Terrain(rand.Intn(int(terrain.Swamp)))))
			}
		}

		if err := os.WriteFile("medoly.svg", b.AsSVG(), 0666); err != nil {
			log.Fatal(err)
		}
		log.Printf("created %q", "medoly.svg")
	},
}

var argsMap struct {
	rows int
	cols int
}

func init() {
	cmdRoot.AddCommand(cmdMap)

	cmdMap.Flags().IntVar(&argsMap.rows, "rows", 200, "number of rows to generate")
	cmdMap.Flags().IntVar(&argsMap.cols, "cols", 400, "number of columns to generate")
}
