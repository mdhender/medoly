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

package main

import (
	"github.com/mdhender/medoly/cmd"
	"log"
	"os"
	"time"
)

func main() {
	started := time.Now()
	//rand.Seed(started.UnixNano())

	rv := 0
	if err := cmd.Execute(); err != nil {
		log.Printf("\n%+v\n", err)
		rv = 2
	}

	log.Printf("completed in %v\n", time.Now().Sub(started))

	os.Exit(rv)
}
