// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package converter

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/murosan/shogi-board-server/app/domain/entity/engine/option"
	"github.com/murosan/shogi-board-server/app/domain/entity/engine/result"
	"github.com/murosan/shogi-board-server/app/domain/entity/shogi"
	"github.com/murosan/shogi-board-server/app/domain/exception"
)

// FromUSI USI から Go に変換するインターフェース
type FromUSI interface {
	// Piece USIの駒を int に変換する
	// TODO: Piece 型を作る
	Piece(string) (int, error)

	// EngineID USI の name と author を変換する
	EngineID(string) (string, string, error)

	// Option USI の option を変換する
	Option(string) (option.Option, error)

	// Move USI の pv を変換する
	Move(string) (shogi.Move, error)

	// Info USI の info を変換する
	Info(string) (result.Info, int, error)
}

type fromUSI struct{}

// NewFromUSI 新しい FromUSI を返す
func NewFromUSI() FromUSI {
	return fromUSI{}
}

func (fu fromUSI) Piece(s string) (i int, e error) {
	switch s {
	case shogi.UsiFu0:
		i = shogi.Fu0
	case shogi.UsiFu1:
		i = shogi.Fu1
	case shogi.UsiKyou0:
		i = shogi.Kyou0
	case shogi.UsiKyou1:
		i = shogi.Kyou1
	case shogi.UsiKei0:
		i = shogi.Kei0
	case shogi.UsiKei1:
		i = shogi.Kei1
	case shogi.UsiGin0:
		i = shogi.Gin0
	case shogi.UsiGin1:
		i = shogi.Gin1
	case shogi.UsiKin0:
		i = shogi.Kin0
	case shogi.UsiKin1:
		i = shogi.Kin1
	case shogi.UsiKaku0:
		i = shogi.Kaku0
	case shogi.UsiKaku1:
		i = shogi.Kaku1
	case shogi.UsiHisha0:
		i = shogi.Hisha0
	case shogi.UsiHisha1:
		i = shogi.Hisha1
	case shogi.UsiGyoku0:
		i = shogi.Gyoku0
	case shogi.UsiGyoku1:
		i = shogi.Gyoku1
	case shogi.UsiTo0:
		i = shogi.To0
	case shogi.UsiTo1:
		i = shogi.To1
	case shogi.UsiNariKyou0:
		i = shogi.NariKyou0
	case shogi.UsiNariKyou1:
		i = shogi.NariKyou1
	case shogi.UsiNariKei0:
		i = shogi.NariKei0
	case shogi.UsiNariKei1:
		i = shogi.NariKei1
	case shogi.UsiNariGin0:
		i = shogi.NariGin0
	case shogi.UsiNariGin1:
		i = shogi.NariGin1
	case shogi.UsiUma0:
		i = shogi.Uma0
	case shogi.UsiUma1:
		i = shogi.Uma1
	case shogi.UsiRyu0:
		i = shogi.Ryu0
	case shogi.UsiRyu1:
		i = shogi.Ryu1
	default:
		e = exception.InvalidPieceID.WithMsg("PieceIdが不正です id=" + s)
	}
	return
}

// id name <EngineName>
// id author <AuthorName> をパースする
func (fu fromUSI) EngineID(s string) (string, string, error) {
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
func (fu fromUSI) Option(s string) (option.Option, error) {
	t := strings.TrimSpace(s)

	if buttonRegex.MatchString(t) {
		return fu.parseButton(t)
	}
	if checkRegex.MatchString(t) {
		return fu.parseCheck(t)
	}
	if spinRegex.MatchString(t) {
		return fu.parseSpin(t)
	}
	if selectRegex.MatchString(t) {
		return fu.parseSelect(t)
	}
	if stringRegex.MatchString(t) {
		return fu.parseString(t)
	}
	if fileNameRegex.MatchString(t) {
		return fu.parseFileName(t)
	}

	return nil, invalidSyntax(s, optionTypeDescription)
}

// button type を Egn の Options にセットする
// option name <string> type button
func (fu fromUSI) parseButton(s string) (*option.Button, error) {
	res := buttonRegex.FindAllStringSubmatch(s, -1)
	if len(res) == 0 || len(res[0]) < 2 {
		return nil, invalidSyntax(s, parseErrorButton)
	}

	return option.NewButton(res[0][1]), nil
}

// check type を Egn の Options にセットする
// option name <string> type check default <bool>
// このフォーマット以外は許容しない。default が無くてもてもだめ
func (fu fromUSI) parseCheck(s string) (*option.Check, error) {
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
func (fu fromUSI) parseSpin(s string) (*option.Spin, error) {
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
func (fu fromUSI) parseSelect(s string) (*option.Select, error) {
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

// string type を Egn の Options にセットする
// option name <string> type string default <string>
func (fu fromUSI) parseString(s string) (*option.String, error) {
	res := stringRegex.FindAllStringSubmatch(s, -1)
	if len(res) == 0 || len(res[0]) < 3 {
		return nil, invalidSyntax(s, parseErrorString)
	}

	return option.NewString(res[0][1], res[0][2], res[0][2]), nil
}

// option name <string> type filename default <string>
func (fu fromUSI) parseFileName(s string) (*option.FileName, error) {
	res := fileNameRegex.FindAllStringSubmatch(s, -1)
	if len(res) == 0 || len(res[0]) < 3 {
		return nil, invalidSyntax(s, parseErrorFileName)
	}

	return option.NewFileName(res[0][1], res[0][2], res[0][2]), nil
}

func invalidSyntax(input, msg string) error {
	return exception.InvalidOptionSyntax.WithMsg(msg + "\nInput: " + input + "\n")
}

func (fu fromUSI) Move(s string) (m shogi.Move, err error) {
	defer func() {
		if rec := recover(); rec != nil {
			err = exception.UnknownCharacter.WithMsg(fmt.Sprintf("%v", rec))
		}
	}()

	a := strings.Split(strings.TrimSpace(s), "")

	if strings.Contains(s, "*") {
		// 持ち駒から
		piece, er := fu.Piece(a[0])
		if er != nil {
			err = exception.UnknownCharacter
			return
		}

		return shogi.Move{
			Source:  shogi.Point{Row: -1, Column: -1},
			Dest:    shogi.Point{Row: fu.row(a[3]), Column: fu.column(a[2])},
			PieceID: piece,
		}, nil
	}

	return shogi.Move{
		Source:     shogi.Point{Row: fu.row(a[1]), Column: fu.column(a[0])},
		Dest:       shogi.Point{Row: fu.row(a[3]), Column: fu.column(a[2])},
		IsPromoted: len(a) == 5 && a[4] == "+",
	}, nil
}

func (fu fromUSI) column(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	if i < 1 || i > 9 {
		panic(exception.InvalidColumnNumber)
	}
	return i - 1 // 0-8 にする
}

func (fu fromUSI) row(s string) int {
	r := []rune(s)[0]
	if r < 97 || r > 105 {
		panic(exception.InvalidRowNumber)
	}
	return int(r - 97)
}

// Info をパース(info string は渡さない)
// return
//   r *result.Info パースした結果。失敗したら nil
//   mpv int multipvならその値。multipvじゃなければ 0
//   err error エラー
func (fu fromUSI) Info(s string) (r result.Info, mpv int, err error) {
	a := strings.Split(s, " ")
	r = result.NewInfo()

	// panic をリカバーしてエラーをセット
	defer func() {
		if rec := recover(); rec != nil {
			err = exception.FailedToParseInfo.WithMsg(fmt.Sprintf("%v", rec))
		}
	}()

	i := 0
	for i < len(a) {
		switch strings.TrimSpace(a[i]) {
		case result.Depth:
			i++
			r.Values[result.Depth] = toInt(a[i])
		case result.SelDepth:
			i++
			r.Values[result.SelDepth] = toInt(a[i])
		case result.Time:
			i++
			r.Values[result.Time] = toInt(a[i])
		case result.Nodes:
			i++
			r.Values[result.Nodes] = toInt(a[i])
		case result.HashFull:
			i++
			r.Values[result.HashFull] = toInt(a[i])
		case result.Nps:
			i++
			r.Values[result.Nps] = toInt(a[i])
		case result.Score:
			if a[i+1] == "cp" || a[i+1] == "mate" {
				r.Score = toInt(a[i+2])
			}
			i += 2
		case result.MultiPv:
			i++
			mpv = toInt(a[i])
		case result.Pv:
			fu.setMoves(&r, a[i+1:])
			i += len(a) // pv は 末尾
		}
		i++
	}

	return
}

func (fu fromUSI) setMoves(r *result.Info, a []string) {
	m := make([]shogi.Move, len(a))
	for i, v := range a {
		mv, err := fu.Move(v)
		if err != nil {
			panic(err)
		}
		m[i] = mv
	}

	r.Moves = m
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
