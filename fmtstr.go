// Copyright (c) 2024  The Go-Curses Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fmtstr

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

var (
	ErrPosArgNotImpl = errors.New("positional argument support for width and precision is not implemented yet")
)

func Parse(format string, argv ...string) (replaced, labelled string, variables Variables, err error) {

	var opened, closed bool
	var position string
	var list Variables
	var state *cState

	currentPos := 1 // currentPos is the positional parameter index, not a string index
	last := len(format) - 1

	for i := 0; i <= last; i++ {
		r := rune(format[i])
		char := string(r)

		var ok bool
		if state, ok, err = checkContinue(currentPos, r, state); err != nil {
			return
		} else if !ok {
			continue
		}

		// if opened, record position until closed
		if opened {

			switch r {
			case ']':
			default:
				// positional brace opened
				if !closed {
					if !unicode.IsDigit(r) {
						err = fmt.Errorf("invalid format at: %v", state.source)
						return
					}
					position += char
					continue
				}
			}

		}

		switch r {
		case 'b', 'c', 'd', 'e', 'E', 'f', 'F', 'g', 'G', 'o', 'O', 'p', 'q', 's', 't', 'T', 'U', 'v', 'x', 'X':
			// found a valid variable type which concludes this substitution
			// variable parameter

			state.verb = Verb(char)

			list = append(list, state.make(argv))

			state = nil
			opened = false
			closed = false
			position = ""
			currentPos += 1

		case '+', '-', '#', ' ':
			state.updatePMHS(r, char)

		case '*':
			// this is supposed to be replaced with the value of the
			// preceding positional indicator's corresponding argv index,
			// see godoc example:
			//  fmt.Sprintf("%[3]*.[2]*[1]f", 12.0, 2, 6)
			//  is equivalent to:
			//  fmt.Sprintf("%6.2f", 12.0)
			//
			// currently fmtstr does not support this advanced usage
			err = ErrPosArgNotImpl
			return

		case '[':
			opened = true

		case ']':
			closed = true
			v, _ := strconv.Atoi(position)
			state.pos = v
			currentPos = v

		case '.':
			state.decimal = true

		default:

			if unicode.IsDigit(r) {
				// no opened brace, is width or precision
				state.updateDigitFlag(r, char)
				continue
			}

			// not a digit and not a flag
			err = fmt.Errorf("invalid format at: %v", state.source)
			return
		}

	}

	replaced, labelled, variables, err = list.process(format, argv)
	return
}

func checkContinue(currentPos int, r rune, state *cState) (parsed *cState, proceed bool, err error) {
	if parsed = state; parsed == nil {

		// state is nil
		if r == '%' {
			// found new opening, start a new state
			parsed = &cState{
				source: "%",
				pos:    currentPos,
			}
		}
		return

	} else if r == '%' {

		// state exists, already processing things
		// this may be a literal percent substitution
		if parsed.source == "%" {
			// this is the second half of the literal percent
			parsed = nil
			return
		}
		// found another opening
		err = fmt.Errorf("invalid format at: %v", parsed.source+"%")
		return

	}

	// process this rune and state
	proceed = true
	state.source += string(r)
	return
}
