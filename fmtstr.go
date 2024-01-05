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
	built := make(map[int][]*cState)

	for i := 0; i <= last; i++ {
		r := rune(format[i])
		char := string(r)

		var ok bool
		if state, ok, err = checkContinue(currentPos, r, state); err != nil {
			return
		} else if !ok {
			continue
		}

		// update the source value
		state.source += char

		switch r {
		case 'b', 'c', 'd', 'e', 'E', 'f', 'F', 'g', 'G', 'o', 'O', 'p', 'q', 's', 't', 'T', 'U', 'v', 'x', 'X':
			// found a valid variable type which concludes this substitution
			// variable parameter

			state.verb = Verb(char)
			state.value += char

			subPos := calcVarPos(state, built)
			list = append(list, state.make(subPos, argv))

			built[state.pos] = append(built[state.pos], state)

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
			state.value += char
			state.decimal = true

		default:

			if opened {
				// positional brace opened
				if !closed && unicode.IsDigit(r) {
					state.value += char
					position += char
				}
				continue

			} else if unicode.IsDigit(r) {
				// no opened brace, is width or precision
				state.value += char
				if r == '0' && state.width == "" && state.precision == "" {
					state.zero = true
				} else if state.decimal {
					state.precision += char
				} else {
					state.width += char
				}

			} else {
				err = fmt.Errorf("invalid format at: %v", state.source)
				return
			}

		}

	}

	if replaced, labelled, variables, err = list.process(format); err == nil {
		if varc, argc := variables.Count(), len(argv); varc > argc {
			err = fmt.Errorf("format requires %d argument(s), received %d instead", varc, argc)
			replaced, labelled, variables = "", "", nil
		}
	}
	return
}

func checkContinue(currentPos int, char rune, state *cState) (parsed *cState, proceed bool, err error) {
	if parsed = state; state == nil {
		if char == '%' {
			// found new opening
			parsed = &cState{
				source: "%",
				pos:    currentPos,
			}
		}
	} else if char == '%' {
		if parsed.source == "%" {
			parsed = nil
			return
		}
		// found another opening
		err = fmt.Errorf("invalid format at: %v", parsed.source+"%")
	} else {
		// process this char and state
		proceed = true
	}
	return
}

func calcVarPos(state *cState, built map[int][]*cState) (subPos int) {
	subPos = -1
	if found, ok := built[state.pos]; ok {
		for idx, other := range found {
			if other.value == state.value {
				subPos = idx
				break
			}
		}
		if subPos == -1 {
			subPos = len(found)
		}
	}
	return
}
