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
	"net/http"
	"sync"
)

// cmdTD runs the TD command
var cmdTD = &cobra.Command{
	Use:   "td",
	Short: "td things",
	Run: func(cmd *cobra.Command, args []string) {
		tdData := td.TD()
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "image/svg+xml")
			_, _ = w.Write(tdData)
		})

		var mu sync.Mutex
		var cycles int
		http.HandleFunc("/shatter", func(w http.ResponseWriter, r *http.Request) {
			mu.Lock()
			cycles++
			mu.Unlock()
			log.Printf("shatter: %8d cycles\n", cycles)
			w.Header().Set("Content-Type", "image/svg+xml")
			_, _ = w.Write(td.Shatter(cycles))
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
