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
	"github.com/mdhender/medoly/pkg/hexes"
	"github.com/spf13/cobra"
	"log"
)

// cmdArea runs the area command
var cmdArea = &cobra.Command{
	Use:   "area",
	Short: "area things",
	Run: func(cmd *cobra.Command, args []string) {
		r := 0
		if argsArea.minHexes > 0 {
			for ; hexes.Area(r) < argsArea.minHexes; r++ {
				//
			}
		} else if argsArea.radius > 0 {
			r = argsArea.radius
		} else {
			r = 110
		}
		log.Printf("radius %4d: area %8d\n", r-1, hexes.Area(r-1))
		log.Printf("radius %4d: area %8d\n", r, hexes.Area(r))
		log.Printf("radius %4d: area %8d\n", r+1, hexes.Area(r+1))
	},
}

var argsArea struct {
	minHexes int
	radius   int
}

func init() {
	cmdRoot.AddCommand(cmdArea)

	cmdArea.Flags().IntVar(&argsArea.radius, "a", 0, "calculate area for radius")
	cmdArea.Flags().IntVar(&argsArea.minHexes, "n", 0, "calculate radius to enclose hexes")
}
