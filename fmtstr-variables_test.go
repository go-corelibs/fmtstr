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
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestVariables(t *testing.T) {
	Convey("String", t, func() {
		f := Variables{}
		So(f.String(), ShouldEqual, "")
		f = append(f, &Variable{
			Pos:  2,
			Verb: "s",
		})
		So(f.String(), ShouldEqual, "%[2]s")
		f = append(f, &Variable{
			Pos:  1,
			Verb: "s",
		})
		So(f.String(), ShouldEqual, "%[2]s %[1]s")
	})

	Convey("Process", t, func() {
		v := Variables{}
		replaced, labelled, variables, err := v.process("", nil)
		So(err, ShouldEqual, nil)
		So(replaced, ShouldEqual, "")
		So(labelled, ShouldEqual, "")
		So(len(variables), ShouldEqual, 0)

		v = Variables{
			{Type: "string", Label: "Key", Source: "%s", Pos: 1, Verb: "s"},
			{Type: "string", Label: "AnotherKey", Source: "%s", Pos: 2, Verb: "s"},
		}

		replaced, labelled, variables, err = v.process("Test %s %s", []string{"", ""})
		So(err, ShouldEqual, nil)
		So(replaced, ShouldEqual, "Test %[1]s %[2]s")
		So(labelled, ShouldEqual, "Test {Key} {AnotherKey}")
		So(len(variables), ShouldEqual, 2)

		v[1].Type = "int"
		v[1].Pos = 1
		v[1].Verb = "d"

		replaced, labelled, variables, err = v.process("Test %[1]s %[2]s", []string{"", ""})
		So(err, ShouldNotEqual, nil)
		So(replaced, ShouldEqual, "")
		So(labelled, ShouldEqual, "")
		So(len(variables), ShouldEqual, 0)
	})
}
