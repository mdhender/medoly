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
	"github.com/mdhender/medoly/pkg/td"
	"github.com/spf13/cobra"
	"log"
	"math/rand"
	"net/http"
	"sync"
)

// cmdTD runs the TD command
var cmdTD = &cobra.Command{
	Use:   "td",
	Short: "td things",
	Run: func(cmd *cobra.Command, args []string) {
		tdData := td.TD(nil)
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "image/svg+xml")
			_, _ = w.Write(tdData)
		})

		var mu sync.Mutex
		var cycles int
		var land [][]float64
		http.HandleFunc("/shatter", func(w http.ResponseWriter, r *http.Request) {
			mu.Lock()
			if land == nil {
				land = make([][]float64, 102, 102)
				for i := 0; i < len(land); i++ {
					land[i] = make([]float64, 102, 102)
				}
			}
			cycles++
			whack := 100 / cycles
			for i := 0; i < 102; i += whack {
				for j := 0; j < 102; j += whack {
					land[i][j] += float64(rand.Intn(10 - cycles))
				}
			}
			mu.Unlock()
			log.Printf("shatter: %8d cycles\n", cycles)
			w.Header().Set("Content-Type", "image/svg+xml")
			_, _ = w.Write(td.Shatter(func(x, y int) float64 {
				return land[x][y]
			}))
		})

		if err := http.ListenAndServe(":3000", nil); err != nil {
			log.Fatal(err)
		}
	},
}

var argsTD struct {
}

func init() {
	cmdRoot.AddCommand(cmdTD)
}

// https://lemire.me/blog/2022/12/06/fast-midpoint-between-two-integers-without-overflow/
// smallest value no smaller than (x+y)/2
func mids(x, y int) int {
	return (x | y) - ((x ^ y) >> 1)
}

// https://lemire.me/blog/2022/12/06/fast-midpoint-between-two-integers-without-overflow/
// largest value no larger than (x+y)/2
func midl(x, y int) int {
	return ((x ^ y) >> 1) + (x & y)
}
