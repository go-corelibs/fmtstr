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
	"fmt"
	"sort"
	"strings"
)

type Variables []*Variable

func (v Variables) String() (values string) {
	for idx, s := range v {
		if idx > 0 {
			values += " "
		}
		values += s.String()
	}
	return
}

func (v Variables) Sort() (sorted Variables) {
	sorted = append(sorted, v...)
	sort.Slice(sorted, func(i, j int) (less bool) {
		less = sorted[i].Pos < sorted[j].Pos
		return
	})
	return
}

func (v Variables) Count() (argc int) {
	unique := make(map[int]struct{})
	for _, variable := range v {
		unique[variable.Pos] = struct{}{}
	}
	argc = len(unique)
	return
}

func (v Variables) process(format string, argv []string) (replaced, labelled string, variables Variables, err error) {
	replaced = format[:]
	labelled = format[:]

	unique := map[int]*Variable{}
	for _, variable := range v {

		if orig, present := unique[variable.Pos]; present {
			if !orig.Verb.Equal(variable.Verb) {
				if err = fmt.Errorf(`conflicting substitution types: %v != %v`, orig, variable); err != nil {
					replaced, labelled, variables = "", "", nil
					return
				}
			}
		} else {
			unique[variable.Pos] = variable
			variables = append(variables, variable)
		}

		replaced = strings.Replace(replaced, variable.Source, variable.String(), 1)
		labelled = strings.Replace(labelled, variable.Source, "{"+variable.Label+"}", 1)
	}

	if varc, argc := variables.Count(), len(argv); varc > argc {
		err = fmt.Errorf("format requires %d argument(s), received %d instead", varc, argc)
		replaced, labelled, variables = "", "", nil
		return
	}

	variables = variables.Sort()
	return
}
