[![godoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/go-corelibs/fmtstr)
[![codecov](https://codecov.io/gh/go-corelibs/fmtstr/graph/badge.svg?token=qkY3ffmjjC)](https://codecov.io/gh/go-corelibs/fmtstr)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-corelibs/fmtstr)](https://goreportcard.com/report/github.com/go-corelibs/fmtstr)

# fmtstr - format string utility

fmtstr is a package for parsing fmt package format strings and doing
interesting things with the results. This package is primarily built
to support the custom gotext parsing of text and html template files
within the [Go-Enjin] theming system.

# Installation

``` shell
> go get github.com/go-corelibs/fmtstr@latest
```

# Examples

## Parse

``` go
func main() {
    replaced, labelled, variables, err := Parse("Testing: %d %T things", ".Count", ".Data")
    // err == nil in this case
    // replaced == "Testing: %[1]d %[2]T things"
    // labelled == "Testing: {Count} {Data} things"
    // variables contains all the specifics of each format variable used and
    // if you're not building a language translation system, or otherwise doing
    // strange and different things with format strings, this is likely too
    // much information and not the actual Go package you're looking for
}
```

# Go-CoreLibs

[Go-CoreLibs] is a repository of shared code between the [Go-Curses] and
[Go-Enjin] projects.

# License

```
Copyright 2023 The Go-CoreLibs Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use file except in compliance with the License.
You may obtain a copy of the license at

 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```

[Go-CoreLibs]: https://github.com/go-corelibs
[Go-Curses]: https://github.com/go-curses
[Go-Enjin]: https://github.com/go-enjin
