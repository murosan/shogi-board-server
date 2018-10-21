// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package usi2go

import (
	"bytes"
	"strconv"

	"github.com/murosan/shogi-proxy-server/pkg/converter/models"
	"github.com/murosan/shogi-proxy-server/pkg/lib"
	"github.com/murosan/shogi-proxy-server/pkg/msg"
)

var (
	space  = []byte(" ")
	id     = []byte("id")
	opt    = []byte("option")
	name   = []byte("name")
	author = []byte("author")
	tpe    = []byte("type")
	deflt  = []byte("default")
	min    = []byte("min")
	max    = []byte("max")
	selOpt = []byte("var")
)

// id name <EngineName>
// id author <AuthorName> をEngineにセットする
// EngineNameやAuthorNameにスペースが入る場合もあるので最後にJoinしている
// TODO: 正規表現でやるか検討
func ParseId(b []byte) ([]byte, []byte, error) {
	s := bytes.Split(bytes.TrimSpace(b), space)
	if len(s) < 3 || !bytes.Equal(s[0], id) {
		return nil, nil, msg.InvalidIdSyntax
	}

	if bytes.Equal(s[1], name) {
		return name, bytes.Join(s[2:], space), nil
	}

	if bytes.Equal(s[1], author) {
		return author, bytes.Join(s[2:], space), nil
	}

	return nil, nil, msg.UnknownOption
}

// 一行受け取って、EngineのOptionMapにセットする
// パースできなかったらエラーを返す
func ParseOpt(b []byte) (models.Option, error) {
	s := bytes.Split(bytes.TrimSpace(b), space)
	if len(s) < 5 || !bytes.Equal(s[0], opt) || !bytes.Equal(s[1], name) || !bytes.Equal(s[3], tpe) || len(s[4]) == 0 {
		return nil, msg.InvalidOptionSyntax
	}

	switch string(s[4]) {
	case "check":
		return parseCheck(s)
	case "spin":
		return parseSpin(s)
	case "combo":
		return parseSelect(s)
	case "button":
		return parseButton(s)
	case "string":
		return parseString(s)
	case "filename":
		return parseFileName(s)
	default:
		return nil, msg.UnknownOptionType
	}
}

// check type を Egn の Options にセットする
// option name <string> type check default <bool>
// このフォーマット以外は許容しない
// default がなかったり、bool ではない時はエラー
func parseCheck(b [][]byte) (models.Option, error) {
	if len(b) != 7 || !bytes.Equal(b[5], deflt) || len(b[2]) == 0 || len(b[6]) == 0 {
		return nil, msg.InvalidOptionSyntax.WithMsg("Received option type was 'check', but malformed. The format must be [option name <string> type check default <bool>]")
	}

	n, d := b[2], b[6]
	if bytes.Equal(d, []byte("true")) {
		return models.Check{Name: n, Val: true, Default: true}, nil
	}
	if bytes.Equal(d, []byte("false")) {
		return models.Check{Name: n, Val: false, Default: false}, nil
	}
	return nil, msg.InvalidOptionSyntax.WithMsg("Default want of 'check' type was not bool. Received: " + string(d))
}

// spin type を Egn の Options にセットする
// option name <string> type spin default <int> min <int> max <int>
// このフォーマット以外は許容しない
// 各値がなかったり、int ではない時、min > max の時はエラー
func parseSpin(b [][]byte) (models.Spin, error) {
	if len(b) != 11 || !bytes.Equal(b[5], deflt) || !bytes.Equal(b[7], min) || !bytes.Equal(b[9], max) || len(b[2]) == 0 {
		return models.Spin{}, msg.InvalidOptionSyntax.WithMsg("Received option type was 'spin', but malformed. The format must be [option name <string> type spin default <int> min <int> max <int>]")
	}

	df, err := strconv.Atoi(string(b[6]))
	if err != nil {
		return models.Spin{}, msg.InvalidOptionSyntax.WithMsg("Default want of 'spin' type was not int. Received: " + string(min))
	}
	mi, err := strconv.Atoi(string(b[8]))
	if err != nil {
		return models.Spin{}, msg.InvalidOptionSyntax.WithMsg("Min want of 'spin' type was not int. Received: " + string(min))
	}
	ma, err := strconv.Atoi(string(b[10]))
	if err != nil {
		return models.Spin{}, msg.InvalidOptionSyntax.WithMsg("Max want of 'spin' type was not int. Received: " + string(min))
	}

	return models.Spin{Name: b[2], Val: df, Default: df, Min: mi, Max: ma}, nil
}

// select type を Egn の Options にセットする
// option name <string> type combo default <string> rep(var <string>)
// このフォーマット以外は許容しない
// Default がない、var がない、default が var にない時はエラー
func parseSelect(b [][]byte) (models.Select, error) {
	if len(b) < 9 || len(b[2]) == 0 || len(b[6]) == 0 {
		return models.Select{}, msg.InvalidOptionSyntax.WithMsg("Received option type was 'combo', but malformed. The format must be [option name <string> type combo default <string> rep(var <string>)]")
	}

	s := models.Select{Name: b[2]}

	i := 8
	for i < len(b) && bytes.Equal(b[i-1], selOpt) {
		s.Vars = append(s.Vars, b[i])
		i += 2
	}

	s.Index = lib.IndexOfBytes(s.Vars, b[6])
	if s.Index == -1 {
		return models.Select{}, msg.InvalidOptionSyntax.WithMsg("Default want of 'combo' type was not found in vars.")
	}

	return s, nil
}

// button type を Egn の Options にセットする
// option name <string> type button
func parseButton(b [][]byte) (models.Button, error) {
	if len(b) != 5 || len(b[2]) == 0 {
		return models.Button{}, msg.InvalidOptionSyntax.WithMsg("Received option type was 'button', but malformed. The format must be [option name <string> type button]")
	}
	return models.Button{Name: b[2]}, nil
}

// string type を Egn の Options にセットする
// option name <string> type string default <string>
func parseString(b [][]byte) (models.String, error) {
	if len(b) != 7 || len(b[2]) == 0 || len(b[6]) == 0 {
		return models.String{}, msg.InvalidOptionSyntax.WithMsg("Received option type was 'string', but malformed. The format must be [option name <string> type string default <string>]")
	}
	return models.String{Name: b[2], Val: b[6], Default: b[6]}, nil
}

// option name <string> type filename default <string>
func parseFileName(b [][]byte) (models.FileName, error) {
	if len(b) != 7 || len(b[2]) == 0 || len(b[6]) == 0 {
		return models.FileName{}, msg.InvalidOptionSyntax.WithMsg("Received option type was 'filename', but malformed. The format must be [option name <string> type filename default <string>]")
	}
	return models.FileName{Name: b[2], Val: b[6], Default: b[6]}, nil
}
