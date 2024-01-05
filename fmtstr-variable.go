// Copyright (c) 2023  The Go-Curses Authors
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
	"strconv"
)

type Modifier uint8

const (
	NoModifiers Modifier = 0
	ModPlus     Modifier = 1 << iota
	ModMinus
	ModHash
	ModSpace
	ModDecimal
	ModZeroPad
)

type Variable struct {
	Type      string
	Label     string
	Source    string
	Pos       int
	Verb      Verb
	Value     string
	Width     int
	Precision int
	Modifiers Modifier
}

func (v *Variable) String() (value string) {
	value = "%"
	if v.Has(ModHash) {
		value += "#"
	}
	if v.Has(ModPlus) {
		value += "+"
	}
	if v.Has(ModMinus) {
		value += "-"
	}
	if v.Has(ModSpace) {
		value += " "
	}
	if v.Has(ModZeroPad) {
		value += "0"
	}
	if v.Width > 0 {
		value += strconv.Itoa(v.Width)
	}
	if v.Has(ModDecimal) {
		value += "."
		if v.Precision > 0 {
			value += strconv.Itoa(v.Precision)
		}
	}
	value += "[" + strconv.Itoa(v.Pos) + "]"
	value += v.Verb.String()
	return
}

func (v *Variable) Has(m Modifier) (present bool) {
	present = v.Modifiers&m == m
	return
}
