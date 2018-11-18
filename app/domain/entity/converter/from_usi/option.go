// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package from_usi

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/murosan/shogi-proxy-server/app/domain/entity/engine/option"
	"github.com/murosan/shogi-proxy-server/app/domain/entity/exception"
)

// id name <EngineName>
// id author <AuthorName> をパースする
func (fu *FromUsi) EngineID(s string) (string, string, error) {
	if res := nameRegex.FindAllStringSubmatch(s, -1); len(res) != 0 {
		return name, res[0][1], nil
	}

	if res := authorRegex.FindAllStringSubmatch(s, -1); len(res) != 0 {
		return author, res[0][1], nil
	}

	return "", "", exception.UnknownOption
}

// エンジンのオプションをパースする
// @param s string 一行のオプションのUSI文字列
// @returns Option オプション。Button|Check|Select|Spin|String|FileName
//          error パースに失敗した場合はエラー。成功時は nil
func (fu *FromUsi) Option(s string) (option.Option, error) {
	t := strings.TrimSpace(s)

	if checkRegex.MatchString(t) {
		return fu.parseCheck(t)
	}
	if spinRegex.MatchString(t) {
		return fu.parseSpin(t)
	}
	if selectRegex.MatchString(t) {
		return fu.parseSelect(t)
	}
	if buttonRegex.MatchString(t) {
		return fu.parseButton(t)
	}
	if stringRegex.MatchString(t) {
		return fu.parseString(t)
	}
	if fileNameRegex.MatchString(t) {
		return fu.parseFileName(t)
	}

	return nil, invalidSyntax(s, optionTypeDescription)
}

// check type を Egn の Options にセットする
// option name <string> type check default <bool>
// このフォーマット以外は許容しない。default が無くてもてもだめ
func (fu *FromUsi) parseCheck(s string) (*option.Check, error) {
	res := checkRegex.FindAllStringSubmatch(s, -1)
	if len(res) == 0 || len(res[0]) < 3 {
		return nil, invalidSyntax(s, parseErrorCheck)
	}

	b := res[0][2] == "true"
	return option.NewCheck(res[0][1], b, b), nil
}

// spin type を Egn の Options にセットする
// option name <string> type spin default <int> min <int> max <int>
// このフォーマット以外は許容しない
// 各値がなかったり、int ではない時、min > max の時はエラー
func (fu *FromUsi) parseSpin(s string) (*option.Spin, error) {
	res := spinRegex.FindAllStringSubmatch(s, -1)
	if len(res) == 0 || len(res[0]) < 5 {
		return nil, invalidSyntax(s, parseErrorSpin)
	}

	init, err := strconv.Atoi(res[0][2])
	if err != nil {
		return nil, invalidSyntax(s, parseErrorSpin+" Default was not int. Value: "+res[0][2])
	}

	min, err := strconv.Atoi(res[0][3])
	if err != nil {
		return nil, invalidSyntax(s, parseErrorSpin+" Min was not int. Value: "+res[0][3])
	}

	max, err := strconv.Atoi(res[0][4])
	if err != nil {
		return nil, invalidSyntax(s, parseErrorSpin+" Max was not int. Value: "+res[0][4])
	}

	if min > max {
		return nil, invalidSyntax(s, fmt.Sprintf("%s Min value is not lesser than or equal to Max. Min: %d, Max: %d", s, min, max))
	}

	return option.NewSpin(res[0][1], init, init, min, max), nil
}

// select type を Egn の Options にセットする
// option name <string> type combo default <string> rep(var <string>)
// このフォーマット以外は許容しない
// initial がない、var がない、default が var にない時はエラー
func (fu *FromUsi) parseSelect(s string) (*option.Select, error) {
	res := selectRegex.FindAllStringSubmatch(s, -1)
	if len(res) == 0 || len(res[0]) < 4 {
		return nil, invalidSyntax(s, parseErrorSelect)
	}

	init := res[0][2]
	vars := strings.Split(res[0][3], selVar)

	if len(vars) < 2 {
		return nil, invalidSyntax(s, parseErrorSelect+" Number of vars must be at least one.")
	}

	vars = vars[1:] // 先頭は空白なので削る
	valid := false  // デフォルト値がvarsに含まれているかどうか

	for i, v := range vars {
		vars[i] = strings.TrimSpace(v)
		valid = valid || vars[i] == init
	}

	// vars にデフォルト値がない場合はエラー
	if !valid {
		return nil, invalidSyntax(s, fmt.Sprintf("%s Default value of Select was not in vars. Default: %s, Vars: %v", parseErrorSelect, init, vars))
	}

	return option.NewSelect(res[0][1], init, init, vars), nil
}

// button type を Egn の Options にセットする
// option name <string> type button
func (fu *FromUsi) parseButton(s string) (*option.Button, error) {
	res := buttonRegex.FindAllStringSubmatch(s, -1)
	if len(res) == 0 || len(res[0]) < 2 {
		return nil, invalidSyntax(s, parseErrorButton)
	}

	return option.NewButton(res[0][1]), nil
}

// string type を Egn の Options にセットする
// option name <string> type string default <string>
func (fu *FromUsi) parseString(s string) (*option.String, error) {
	res := stringRegex.FindAllStringSubmatch(s, -1)
	if len(res) == 0 || len(res[0]) < 3 {
		return nil, invalidSyntax(s, parseErrorString)
	}

	return option.NewString(res[0][1], res[0][2], res[0][2]), nil
}

// option name <string> type filename default <string>
func (fu *FromUsi) parseFileName(s string) (*option.FileName, error) {
	res := fileNameRegex.FindAllStringSubmatch(s, -1)
	if len(res) == 0 || len(res[0]) < 3 {
		return nil, invalidSyntax(s, parseErrorFileName)
	}

	return option.NewFileName(res[0][1], res[0][2], res[0][2]), nil
}

func invalidSyntax(input, msg string) error {
	return exception.InvalidOptionSyntax.WithMsg(msg + "\nInput: " + input + "\n")
}
