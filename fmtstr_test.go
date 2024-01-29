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

func Test(t *testing.T) {
	Convey("Decompose", t, func() {

		replaced, labelled, variables, err := Decompose("")
		So(err, ShouldEqual, nil)
		So(replaced, ShouldEqual, "")
		So(labelled, ShouldEqual, "")
		So(len(variables), ShouldEqual, 0)

		replaced, labelled, variables, err = Decompose("No vars")
		So(err, ShouldEqual, nil)
		So(replaced, ShouldEqual, "No vars")
		So(labelled, ShouldEqual, "No vars")
		So(len(variables), ShouldEqual, 0)

		replaced, labelled, variables, err = Decompose("One var %d", ".Key")
		So(err, ShouldEqual, nil)
		So(replaced, ShouldEqual, "One var %[1]d")
		So(labelled, ShouldEqual, "One var {Key}")
		So(len(variables), ShouldEqual, 1)

		replaced, labelled, variables, err = Decompose("Two vars %[2]d %[1]s", ".str_var_name", ".intVarName")
		So(err, ShouldEqual, nil)
		So(replaced, ShouldEqual, "Two vars %[2]d %[1]s")
		So(labelled, ShouldEqual, "Two vars {IntVarName} {StrVarName}")
		So(len(variables), ShouldEqual, 2)

		replaced, labelled, variables, err = Decompose("Same vars %% %[1]s %[1]s", ".var_name")
		So(err, ShouldEqual, nil)
		So(replaced, ShouldEqual, "Same vars %% %[1]s %[1]s")
		So(labelled, ShouldEqual, "Same vars %% {VarName} {VarName}")
		So(len(variables), ShouldEqual, 1)

		replaced, labelled, variables, err = Decompose("Same vars %[1]%[1]s", ".var_name")
		So(err, ShouldNotEqual, nil)
		So(err.Error(), ShouldEqual, "invalid format at: %[1]%")
		So(replaced, ShouldEqual, "")
		So(labelled, ShouldEqual, "")
		So(len(variables), ShouldEqual, 0)

		replaced, labelled, variables, err = Decompose("Two vars %[1]d %[1]s", ".bad_var_name")
		So(err, ShouldNotEqual, nil)
		So(replaced, ShouldEqual, "")
		So(labelled, ShouldEqual, "")
		So(len(variables), ShouldEqual, 0)

		replaced, labelled, variables, err = Decompose("Two vars %10.2f %s", ".var_name", "$moar")
		So(err, ShouldEqual, nil)
		So(replaced, ShouldEqual, "Two vars %10.2[1]f %[2]s")
		So(labelled, ShouldEqual, "Two vars {VarName} {Moar}")
		So(len(variables), ShouldEqual, 2)

		replaced, labelled, variables, err = Decompose("Two vars %10!2f %s", ".var_name", "$moar")
		So(err, ShouldNotEqual, nil)
		So(err.Error(), ShouldEqual, "invalid format at: %10!")
		So(replaced, ShouldEqual, "")
		So(labelled, ShouldEqual, "")
		So(len(variables), ShouldEqual, 0)

		replaced, labelled, variables, err = Decompose("Two vars %[3]*.2f %s", ".var_name", "$moar")
		So(err, ShouldEqual, ErrPosArgNotImpl)
		So(replaced, ShouldEqual, "")
		So(labelled, ShouldEqual, "")
		So(len(variables), ShouldEqual, 0)

		replaced, labelled, variables, err = Decompose("One var %#+v", ".var_name")
		So(err, ShouldEqual, nil)
		So(replaced, ShouldEqual, "One var %#+[1]v")
		So(labelled, ShouldEqual, "One var {VarName}")
		So(len(variables), ShouldEqual, 1)

		replaced, labelled, variables, err = Decompose("One var %-02.f", ".var_name")
		So(err, ShouldEqual, nil)
		So(replaced, ShouldEqual, "One var %-02.[1]f")
		So(labelled, ShouldEqual, "One var {VarName}")
		So(len(variables), ShouldEqual, 1)

		replaced, labelled, variables, err = Decompose("One var %-02.f %f %v", ".var_name")
		So(err, ShouldEqual, nil)
		So(replaced, ShouldEqual, "One var %-02.[1]f %[2]f %[3]v")
		So(labelled, ShouldEqual, "One var {VarName} {Float} {Var}")
		So(len(variables), ShouldEqual, 3)

		replaced, labelled, variables, err = Decompose("One var %t", ".var_name")
		So(err, ShouldEqual, nil)
		So(replaced, ShouldEqual, "One var %[1]t")
		So(labelled, ShouldEqual, "One var {VarName}")
		So(len(variables), ShouldEqual, 1)

		replaced, labelled, variables, err = Decompose("One var % d", ".var_name")
		So(err, ShouldEqual, nil)
		So(replaced, ShouldEqual, "One var % [1]d")
		So(labelled, ShouldEqual, "One var {VarName}")
		So(len(variables), ShouldEqual, 1)

		replaced, labelled, variables, err = Decompose("One var %d %d", "", "")
		So(err, ShouldEqual, nil)
		So(replaced, ShouldEqual, "One var %[1]d %[2]d")
		So(labelled, ShouldEqual, "One var {Num} {Num1}")
		So(len(variables), ShouldEqual, 2)

		replaced, labelled, variables, err = Decompose("One var %d %[1]d %d", ".Var")
		So(err, ShouldEqual, nil)
		So(replaced, ShouldEqual, "One var %[1]d %[1]d %[2]d")
		So(labelled, ShouldEqual, "One var {Var} {Var} {Num}")
		So(len(variables), ShouldEqual, 2)

		replaced, labelled, variables, err = Decompose("One var %d %d", ".VarName", ".VarName")
		So(err, ShouldEqual, nil)
		So(replaced, ShouldEqual, "One var %[1]d %[2]d")
		So(labelled, ShouldEqual, "One var {VarName} {VarName1}")
		So(len(variables), ShouldEqual, 2)

		replaced, labelled, variables, err = Decompose("One var %[11]d", ".V0", ".V1", ".V2", ".V3", ".V4", ".V5", ".V6", ".V7", ".V8", ".V9", ".V10")
		So(err, ShouldEqual, nil)
		So(replaced, ShouldEqual, "One var %[11]d")
		So(labelled, ShouldEqual, "One var {V10}")
		So(len(variables), ShouldEqual, 1)

		replaced, labelled, variables, err = Decompose("One var %[2.5]d", ".var_name")
		So(err, ShouldNotEqual, nil)
		So(err.Error(), ShouldEqual, "invalid format at: %[2.")
		So(replaced, ShouldEqual, "")
		So(labelled, ShouldEqual, "")
		So(len(variables), ShouldEqual, 0)

	})
}
