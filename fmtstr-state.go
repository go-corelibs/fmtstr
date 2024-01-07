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
	"fmt"
	"strconv"

	"github.com/go-corelibs/convert"
	"github.com/iancoleman/strcase"
)

type cState struct {
	// pos is the current replacement variable position, not a slice index
	pos int
	// verb is the `s` in %s
	verb Verb

	plus    bool
	minus   bool
	hash    bool
	space   bool
	decimal bool
	zero    bool

	width     string
	precision string

	source string
}

func (s *cState) modifiers() (m Modifier) {
	if s.plus {
		m |= ModPlus
	}
	if s.minus {
		m |= ModMinus
	}
	if s.hash {
		m |= ModHash
	}
	if s.space {
		m |= ModSpace
	}
	if s.decimal {
		m |= ModDecimal
	}
	if s.zero {
		m |= ModZeroPad
	}
	return
}

func (s *cState) updatePMHS(r rune, char string) {
	switch r {
	case '+':
		s.plus = true

	case '-':
		s.minus = true

	case '#':
		s.hash = true

	case ' ':
		s.space = true
	}
}

func (s *cState) updateDigitFlag(r rune, char string) {
	if r == '0' && s.width == "" && s.precision == "" {
		s.zero = true
	} else if s.decimal {
		s.precision += char
	} else {
		s.width += char
	}
	return
}

func (s *cState) make(subPos int, argv []string) (variable *Variable) {
	var label, valueType string
	valueType = s.verb.Type()

	if len(argv) >= s.pos {
		label = strcase.ToCamel(TrimVarPrefix(argv[s.pos-1]))
		if label != "" && subPos > 0 {
			label += fmt.Sprintf("%d%s", s.pos, convert.ToLetters(subPos))
		}
	}

	if label == "" {
		label = fmt.Sprintf("%s%d", s.verb.Label(), s.pos)
	}

	width, precision := 0, 0
	if s.width != "" {
		width, _ = strconv.Atoi(s.width)
	}
	if s.precision != "" {
		precision, _ = strconv.Atoi(s.precision)
	}

	return &Variable{
		Type:      valueType,
		Label:     label,
		Source:    s.source,
		Pos:       s.pos,
		Verb:      s.verb,
		Width:     width,
		Precision: precision,
		Modifiers: s.modifiers(),
	}
}
