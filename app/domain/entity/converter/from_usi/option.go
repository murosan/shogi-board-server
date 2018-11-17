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
	"github.com/murosan/shogi-proxy-server/app/lib/stringutil"
)

// TODO: byteじゃなくてstringにしてからsplitする

// id name <EngineName>
// id author <AuthorName> をEngineにセットする
// EngineNameやAuthorNameにスペースが入る場合もあるので最後にJoinしている
// TODO: 正規表現でやるか検討
func (fu *FromUsi) EngineID(b string) (string, string, error) {
	s := strings.Split(strings.TrimSpace(b), space)
	if len(s) < 3 || s[0] != id {
		return "", "", exception.InvalidIdSyntax
	}

	if s[1] == name {
		return name, strings.Join(s[2:], space), nil
	}

	if s[1] == author {
		return author, strings.Join(s[2:], space), nil
	}

	return "", "", exception.UnknownOption
}

// 一行受け取って、EngineのOptionMapにセットする
// パースできなかったらエラーを返す
func (fu *FromUsi) Option(b string) (option.Option, error) {
	s := strings.Split(strings.TrimSpace(b), space)
	if len(s) < 5 || s[0] != opt || s[1] != name || s[3] != tpe || len(s[4]) == 0 {
		return nil, exception.InvalidOptionSyntax
	}

	switch string(s[4]) {
	case check:
		return fu.parseCheck(s)
	case spin:
		return fu.parseSpin(s)
	case combo:
		return fu.parseSelect(s)
	case button:
		return fu.parseButton(s)
	case str:
		return fu.parseString(s)
	case fileName:
		return fu.parseFileName(s)
	default:
		return nil, exception.UnknownOptionType
	}
}

// check type を Egn の Options にセットする
// option name <string> type check default <bool>
// このフォーマット以外は許容しない
// default がなかったり、bool ではない時はエラー
func (fu *FromUsi) parseCheck(b []string) (*option.Check, error) {
	if len(b) != 7 || b[5] != deflt || len(b[2]) == 0 || len(b[6]) == 0 {
		return nil, exception.InvalidOptionSyntax.WithMsg("Received option type was 'check', but malformed. The format must be [option name <string> type check default <bool>]")
	}

	n, d := string(b[2]), string(b[6])
	if d != "true" && d != "false" {
		return nil, exception.InvalidOptionSyntax.WithMsg("initial want of 'check' type was not bool. Received: " + string(d))
	}

	boolVal := d == "true"
	return option.NewCheck(n, boolVal, boolVal), nil
}

// spin type を Egn の Options にセットする
// option name <string> type spin default <int> min <int> max <int>
// このフォーマット以外は許容しない
// 各値がなかったり、int ではない時、min > max の時はエラー
func (fu *FromUsi) parseSpin(b []string) (*option.Spin, error) {
	if len(b) != 11 || b[5] != deflt || b[7] != min || b[9] != max || len(b[2]) == 0 {
		return nil, exception.InvalidOptionSyntax.WithMsg("Received option type was 'spin', but malformed. The format must be [option name <string> type spin default <int> min <int> max <int>]")
	}

	df, err := strconv.Atoi(string(b[6]))
	if err != nil {
		return nil, exception.InvalidOptionSyntax.WithMsg("initial want of 'spin' type was not int. Received: " + string(min))
	}
	mi, err := strconv.Atoi(string(b[8]))
	if err != nil {
		return nil, exception.InvalidOptionSyntax.WithMsg("min want of 'spin' type was not int. Received: " + string(min))
	}
	ma, err := strconv.Atoi(string(b[10]))
	if err != nil {
		return nil, exception.InvalidOptionSyntax.WithMsg("max want of 'spin' type was not int. Received: " + string(min))
	}

	return option.NewSpin(string(b[2]), df, df, mi, ma), nil
}

// select type を Egn の Options にセットする
// option name <string> type combo default <string> rep(var <string>)
// このフォーマット以外は許容しない
// initial がない、var がない、default が var にない時はエラー
func (fu *FromUsi) parseSelect(b []string) (*option.Select, error) {
	if len(b) < 9 || len(b[2]) == 0 || len(b[6]) == 0 {
		return nil, exception.InvalidOptionSyntax.WithMsg("Received option type was 'combo', but malformed. The format must be [option name <string> type combo default <string> rep(var <string>)]")
	}

	var (
		name = string(b[2])
		init = string(b[6]) // デフォルト。TODO: もうちょい方法を考える。正規表現でやるかなぁ・・・
		vars []string
	)

	i := 8
	for i < len(b) && b[i-1] == selOpt {
		str := string(b[i])
		vars = append(vars, str)
		i += 2
	}

	if stringutil.IndexOf(vars, init) == -1 {
		return nil, exception.InvalidOptionSyntax.WithMsg(fmt.Sprintf("Default value of Select was not in vars. default: %s, vars: %v", init, vars))
	}

	return option.NewSelect(name, init, init, vars), nil
}

// button type を Egn の Options にセットする
// option name <string> type button
func (fu *FromUsi) parseButton(b []string) (*option.Button, error) {
	if len(b) != 5 || len(b[2]) == 0 {
		return nil, exception.InvalidOptionSyntax.WithMsg("Received option type was 'button', but malformed. The format must be [option name <string> type button]")
	}
	return option.NewButton(string(b[2])), nil
}

// string type を Egn の Options にセットする
// option name <string> type string default <string>
func (fu *FromUsi) parseString(b []string) (*option.String, error) {
	if len(b) != 7 || len(b[2]) == 0 || len(b[6]) == 0 {
		return nil, exception.InvalidOptionSyntax.WithMsg("Received option type was 'string', but malformed. The format must be [option name <string> type string default <string>]")
	}
	return option.NewString(string(b[2]), string(b[6]), string(b[6])), nil
}

// option name <string> type filename default <string>
func (fu *FromUsi) parseFileName(b []string) (*option.FileName, error) {
	if len(b) != 7 || len(b[2]) == 0 || len(b[6]) == 0 {
		return nil, exception.InvalidOptionSyntax.WithMsg("Received option type was 'filename', but malformed. The format must be [option name <string> type filename default <string>]")
	}
	return option.NewFileName(string(b[2]), string(b[6]), string(b[6])), nil
}
