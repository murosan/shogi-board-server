// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package converter

import "regexp"

var (
	name   = "name"
	author = "author"
	selVar = "var"

	nameRegex   = regexp.MustCompile(`^id name (.*)$`)
	authorRegex = regexp.MustCompile(`^id author (.*)$`)

	// name に空白とか \t とかを許容したくないなーw
	// ただし、↓以外のフォーマットは許容しない。min と max が入れ替わっているとかもだめ
	buttonRegex   = regexp.MustCompile(`^option name (.*) type button$`)
	checkRegex    = regexp.MustCompile(`^option name (.*) type check default (true|false)$`)
	spinRegex     = regexp.MustCompile(`^option name (.*) type spin default (-?[0-9]+) min (-?[0-9]+) max (-?[0-9]+)$`)
	selectRegex   = regexp.MustCompile(`^option name (.*) type combo default (.*?) (var .*)$`)
	stringRegex   = regexp.MustCompile(`^option name (.*) type string default (.*)$`)
	fileNameRegex = regexp.MustCompile(`^option name (.*) type filename default (.*)$`)

	parseErrorButton   = "Failed to parse 'button' type."
	parseErrorCheck    = "Failed to parse 'check' type."
	parseErrorSpin     = "Failed to parse 'spin' type."
	parseErrorSelect   = "Failed to parse 'combo' type."
	parseErrorString   = "Failed to parse 'string' type."
	parseErrorFileName = "Failed to parse 'filename' type."

	optionTypeDescription = `
Formats
  Button:   option name <string> type button
  Check:    option name <string> type check default <bool>
  Spin:     option name <string> type spin default <int> min <int> max <int>
  Combo:    option name <string> type combo default <string> rep(var <string>)
  String:   option name <string> type string default <string>
  FileName: option name <string> type filename default <string>
`
)