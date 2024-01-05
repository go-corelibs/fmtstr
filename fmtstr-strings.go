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

// TODO: move TrimVarPrefix to go-corelibs/strings

func TrimVarPrefix(tmplVar string) (trimmed string) {
	if len(tmplVar) > 0 {
		if tmplVar[0] == '$' || tmplVar[0] == '.' {
			trimmed = tmplVar[1:]
			return
		}
	}
	trimmed = tmplVar
	return
}
