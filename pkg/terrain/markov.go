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
	"math/rand"
	"os"
)

// markov chain for random terrain generator

type MarkovCorpusEntry struct {
	Terrain []Terrain      `json:"terrain"`
	Weights []MarkovWeight `json:"weights"`
}

type MarkovWeight struct {
	Terrain Terrain `json:"terrain"`
	Weight  int     `json:"weight"`
}

func key(a, b Terrain) int {
	if a > b {
		a, b = b, a
	}
	return int(a)*32 + int(b)
}

type MarkovChain map[int][]Terrain

func (mk MarkovChain) Next(a, b Terrain) Terrain {
	if choices, ok := mk[key(a, b)]; ok {
		return choices[rand.Intn(len(choices))]
	}
	return b
}

func SampleCorpus(name string) error {
	var corpus []MarkovCorpusEntry
	t1 := Ice
	for _, t2 := range []Terrain{
		Ice,
		Mountain,
		Ocean,
		Steppe,
	} {
		corpus = append(corpus, MarkovCorpusEntry{
			Terrain: []Terrain{t1, t2},
			Weights: []MarkovWeight{{Terrain: t1}},
		})
	}

	for _, t1 := range []Terrain{
		Delta,
		Desert,
		Forest,
		Gravel,
		Mountain,
		Ocean,
		Plain,
		Rock,
		Rough,
		SaltMarsh,
		Sea,
		Steppe,
		Swamp,
	} {
		for _, t2 := range []Terrain{
			Delta,
			Desert,
			Forest,
			Gravel,
			Mountain,
			Ocean,
			Plain,
			Rock,
			Rough,
			SaltMarsh,
			Sea,
			Steppe,
			Swamp,
		} {
			mce := MarkovCorpusEntry{
				Terrain: []Terrain{t1, t2},
			}
			for _, t3 := range []Terrain{
				Delta,
				Desert,
				Forest,
				Gravel,
				Mountain,
				Ocean,
				Plain,
				Rock,
				Rough,
				SaltMarsh,
				Sea,
				Steppe,
				Swamp,
			} {
				mce.Weights = append(mce.Weights, MarkovWeight{
					Terrain: t3,
					Weight:  0,
				})
			}
			corpus = append(corpus, mce)
		}
	}

	if buf, err := json.MarshalIndent(corpus, "", "  "); err != nil {
		return err
	} else if err = os.WriteFile(name, buf, 0666); err != nil {
		return err
	}

	return nil
}

func LoadCorpus(name string) (MarkovChain, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, fmt.Errorf("LoadCorpus: %w", err)
	}
	var corpus []MarkovCorpusEntry
	if err := json.Unmarshal(data, &corpus); err != nil {
		return nil, fmt.Errorf("LoadCorpus: %w", err)
	}

	chain := make(map[int][]Terrain)

	for _, e := range corpus {
		id := key(e.Terrain[0], e.Terrain[1])
		for _, weight := range e.Weights {
			for i := 0; i < weight.Weight; i++ {
				chain[id] = append(chain[id], weight.Terrain)
			}
		}
	}

	return chain, nil
}
