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

package terrain

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
)

// markov chain for random terrain generator

type MarkovCorpus struct {
	Clear     []MarkovWeight `json:"clear"`
	Delta     []MarkovWeight `json:"delta"`
	Desert    []MarkovWeight `json:"desert"`
	Forest    []MarkovWeight `json:"forest"`
	Gravel    []MarkovWeight `json:"gravel"`
	Ice       []MarkovWeight `json:"ice"`
	Mountain  []MarkovWeight `json:"mountain"`
	Ocean     []MarkovWeight `json:"ocean"`
	Plain     []MarkovWeight `json:"plain"`
	Rock      []MarkovWeight `json:"rock"`
	Rough     []MarkovWeight `json:"rough"`
	SaltMarsh []MarkovWeight `json:"saltMarsh"`
	Sea       []MarkovWeight `json:"sea"`
	Steppe    []MarkovWeight `json:"steppe"`
	Swamp     []MarkovWeight `json:"swamp"`
}

type MarkovWeight struct {
	Terrain Terrain `json:"terrain"`
	Weight  int     `json:"weight"`
}

func key(a, b Terrain) int {
	return int(a)*32 + int(b)
}

type MarkovChain map[int][]Terrain

func (mk MarkovChain) Next(a, b Terrain) Terrain {
	if choices, ok := mk[key(a, b)]; ok {
		return choices[rand.Intn(len(choices))]
	}
	return b
}

func LoadCorpus(name string) (MarkovChain, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, fmt.Errorf("LoadCorpus: %w", err)
	}
	var corpus MarkovCorpus
	if err := json.Unmarshal(data, &corpus); err != nil {
		return nil, fmt.Errorf("LoadCorpus: %w", err)
	}

	log.Printf("%+v\n", corpus)

	chain := make(map[int][]Terrain)

	for _, mw := range corpus.Clear {
		k := key(Clear, mw.Terrain)
		for i := 0; i < mw.Weight; i++ {
			chain[k] = append(chain[k], mw.Terrain)
		}
	}
	for _, mw := range corpus.Delta {
		k := key(Delta, mw.Terrain)
		for i := 0; i < mw.Weight; i++ {
			chain[k] = append(chain[k], mw.Terrain)
		}
	}
	for _, mw := range corpus.Desert {
		k := key(Desert, mw.Terrain)
		for i := 0; i < mw.Weight; i++ {
			chain[k] = append(chain[k], mw.Terrain)
		}
	}
	for _, mw := range corpus.Forest {
		k := key(Forest, mw.Terrain)
		for i := 0; i < mw.Weight; i++ {
			chain[k] = append(chain[k], mw.Terrain)
		}
	}
	for _, mw := range corpus.Gravel {
		k := key(Gravel, mw.Terrain)
		for i := 0; i < mw.Weight; i++ {
			chain[k] = append(chain[k], mw.Terrain)
		}
	}
	for _, mw := range corpus.Ice {
		k := key(Ice, mw.Terrain)
		for i := 0; i < mw.Weight; i++ {
			chain[k] = append(chain[k], mw.Terrain)
		}
	}
	for _, mw := range corpus.Mountain {
		k := key(Mountain, mw.Terrain)
		for i := 0; i < mw.Weight; i++ {
			chain[k] = append(chain[k], mw.Terrain)
		}
	}
	for _, mw := range corpus.Ocean {
		k := key(Ocean, mw.Terrain)
		for i := 0; i < mw.Weight; i++ {
			chain[k] = append(chain[k], mw.Terrain)
		}
	}
	for _, mw := range corpus.Plain {
		k := key(Plain, mw.Terrain)
		for i := 0; i < mw.Weight; i++ {
			chain[k] = append(chain[k], mw.Terrain)
		}
	}
	for _, mw := range corpus.Rock {
		k := key(Rock, mw.Terrain)
		for i := 0; i < mw.Weight; i++ {
			chain[k] = append(chain[k], mw.Terrain)
		}
	}
	for _, mw := range corpus.Rough {
		k := key(Rough, mw.Terrain)
		for i := 0; i < mw.Weight; i++ {
			chain[k] = append(chain[k], mw.Terrain)
		}
	}
	for _, mw := range corpus.SaltMarsh {
		k := key(SaltMarsh, mw.Terrain)
		for i := 0; i < mw.Weight; i++ {
			chain[k] = append(chain[k], mw.Terrain)
		}
	}
	for _, mw := range corpus.Sea {
		k := key(Sea, mw.Terrain)
		for i := 0; i < mw.Weight; i++ {
			chain[k] = append(chain[k], mw.Terrain)
		}
	}
	for _, mw := range corpus.Steppe {
		k := key(Steppe, mw.Terrain)
		for i := 0; i < mw.Weight; i++ {
			chain[k] = append(chain[k], mw.Terrain)
		}
	}
	for _, mw := range corpus.Swamp {
		k := key(Swamp, mw.Terrain)
		for i := 0; i < mw.Weight; i++ {
			chain[k] = append(chain[k], mw.Terrain)
		}
	}

	return chain, nil
}
