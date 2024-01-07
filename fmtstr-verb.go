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
	"github.com/iancoleman/strcase"
)

type Verb string

func (v Verb) String() string {
	return string(v)
}

func (v Verb) Label() string {
	if v == "-" || v == "q" || v == "v" {
		return "Var"
	}
	return strcase.ToCamel(v.Type())
}

func (v Verb) Type() string {
	switch v {
	case "b", "c", "d", "o", "O", "p", "U", "x", "X":
		return "num"
	case "e", "E", "f", "F", "g", "G":
		return "float"
	case "s":
		return "text"
	case "t":
		return "bool"
	case "q", "T", "v":
	}
	return "any"
}

func (v Verb) Equal(o Verb) (equal bool) {
	self := v.Type()
	other := o.Type()
	equal = self == "any" || other == "any" || self == other
	return
}
