package usi

import "regexp"

const (
	// TypeButton is used for check the engine output is option type button
	TypeButton = " type button"

	// TypeCheck is used for check the engine output is option type check
	TypeCheck = " type check "

	// TypeRange is used for check the engine output is option type range
	TypeRange = " type spin "

	// TypeSelect is used for check the engine output is option type select
	TypeSelect = " type combo "

	// TypeString is used for check the engine output is option type string
	TypeString = " type string "

	// TypeFilename is used for check the engine output is option type string
	// shogi-board-servre treat the type string and filename as the same type.
	TypeFilename = " type filename "
)

var (
	// option regular expressions
	// We only accept if the output follows these formats.
	buttonRegex   = regexp.MustCompile(`^option name (.*) type button$`)
	checkRegex    = regexp.MustCompile(`^option name (.*) type check default (true|false)$`)
	spinRegex     = regexp.MustCompile(`^option name (.*) type spin default (-?[0-9]+) min (-?[0-9]+) max (-?[0-9]+)$`)
	selectRegex   = regexp.MustCompile(`^option name (.*) type combo default (.*?) (var .*)$`)
	stringRegex   = regexp.MustCompile(`^option name (.*) type string default (.*)$`)
	fileNameRegex = regexp.MustCompile(`^option name (.*) type filename default (.*)$`)

	usiOptionFormat = `
Formats
  Button:   option name <string> type button
  Check:    option name <string> type check default <bool>
  Spin:     option name <string> type spin default <int> min <int> max <int>
  Select:   option name <string> type combo default <string> rep(var <string>)
  Text:     option name <string> type string default <string>
            option name <string> type filename default <string>
`
)
