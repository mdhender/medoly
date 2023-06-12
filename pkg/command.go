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

package pkg

// Command interface
type Command interface {
	Execute() error
	Undo() error
}

// NoOp command does nothing
type NoOp struct{}

// Execute does nothing
func (c NoOp) Execute() error {
	return nil
}

// Undo undoes nothing
func (c NoOp) Undo() error {
	return nil
}
