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
	"bytes"
	"fmt"
)

type Terrain int

const (
	Clear Terrain = iota
	Delta
	Desert
	Forest
	Gravel
	Ice
	Mountain
	Ocean
	Plain
	Rock
	Rough
	SaltMarsh
	Sea
	Steppe
	Swamp
)

// ToFill returns the background fill for a terrain type.
func (t Terrain) ToFill() string {
	switch t {
	case Clear:
		return "hsl(53, 100%, 94%)"
	case Delta:
		return "hsl(74, 48%, 76%)"
	case Desert:
		return "hsl(48, 81%, 66%)"
	case Forest:
		return "hsl(85, 56%, 71%)"
	case Gravel:
		return "hsl(49, 79%, 89%)"
	case Ice:
		return "White"
	case Mountain:
		return "hsl(47, 40%, 63%)"
	case Ocean:
		return "LightBlue"
	case Plain:
		return "hsl(85, 26%, 31%)"
	case Rock:
		return "hsl(49, 79%, 89%)"
	case Rough:
		return "hsl(49, 79%, 89%)"
	case SaltMarsh:
		return "hsl(65, 85%, 90%)"
	case Sea:
		return "hsl(197, 78%, 85%)"
	case Steppe:
		return "hsl(43, 43%, 77%)"
	case Swamp:
		return "hsl(68, 78%, 93%)"
	}
	panic(fmt.Sprintf("assert(t != %d)", t))
}

func (t Terrain) MarshalJSON() ([]byte, error) {
	switch t {
	case Clear:
		return []byte("clear"), nil
	case Delta:
		return []byte("delta"), nil
	case Desert:
		return []byte("desert"), nil
	case Forest:
		return []byte("forest"), nil
	case Gravel:
		return []byte("gravel"), nil
	case Ice:
		return []byte("ice"), nil
	case Mountain:
		return []byte("mountain"), nil
	case Ocean:
		return []byte("ocean"), nil
	case Plain:
		return []byte("plain"), nil
	case Rock:
		return []byte("rock"), nil
	case Rough:
		return []byte("rough"), nil
	case SaltMarsh:
		return []byte("salt-marsh"), nil
	case Sea:
		return []byte("sea"), nil
	case Steppe:
		return []byte("steppe"), nil
	case Swamp:
		return []byte("swamp"), nil
	}
	return nil, fmt.Errorf("invalid value")
}

func (t *Terrain) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, []byte(`"clear"`)) {
		*t = Clear
		return nil
	} else if bytes.Equal(b, []byte(`"delta"`)) {
		*t = Delta
		return nil
	} else if bytes.Equal(b, []byte(`"desert"`)) {
		*t = Desert
		return nil
	} else if bytes.Equal(b, []byte(`"forest"`)) {
		*t = Forest
		return nil
	} else if bytes.Equal(b, []byte(`"gravel"`)) {
		*t = Gravel
		return nil
	} else if bytes.Equal(b, []byte(`"ice"`)) {
		*t = Ice
		return nil
	} else if bytes.Equal(b, []byte(`"mountain"`)) {
		*t = Mountain
		return nil
	} else if bytes.Equal(b, []byte(`"ocean"`)) {
		*t = Ocean
		return nil
	} else if bytes.Equal(b, []byte(`"plain"`)) {
		*t = Plain
		return nil
	} else if bytes.Equal(b, []byte(`"rock"`)) {
		*t = Rock
		return nil
	} else if bytes.Equal(b, []byte(`"rough"`)) {
		*t = Rough
		return nil
	} else if bytes.Equal(b, []byte(`"salt-marsh"`)) {
		*t = SaltMarsh
		return nil
	} else if bytes.Equal(b, []byte(`"sea"`)) {
		*t = Sea
		return nil
	} else if bytes.Equal(b, []byte(`"steppe"`)) {
		*t = Steppe
		return nil
	} else if bytes.Equal(b, []byte(`"swamp"`)) {
		*t = Swamp
		return nil
	}
	*t = Clear
	return fmt.Errorf("invalid terrain: %q", string(b))
}

func (t Terrain) String() string {
	switch t {
	case Clear:
		return "clear"
	case Delta:
		return "delta"
	case Desert:
		return "desert"
	case Forest:
		return "forest"
	case Gravel:
		return "gravel"
	case Ice:
		return "ice"
	case Mountain:
		return "mountain"
	case Ocean:
		return "ocean"
	case Plain:
		return "plain"
	case Rock:
		return "rock"
	case Rough:
		return "rough"
	case SaltMarsh:
		return "salt-marsh"
	case Sea:
		return "sea"
	case Steppe:
		return "steppe"
	case Swamp:
		return "swamp"
	}
	return "clear"
}
