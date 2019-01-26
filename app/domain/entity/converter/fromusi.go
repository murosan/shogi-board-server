// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package converter

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/murosan/shogi-board-server/app/domain/entity/shogi"
	"github.com/murosan/shogi-board-server/app/domain/exception"
	pb "github.com/murosan/shogi-board-server/app/proto"
)

const (
	depth    = "depth"
	selDepth = "seldepth"
	time     = "time"
	nodes    = "nodes"
	hashFull = "hashfull"
	nps      = "nps"
	score    = "score"
	pv       = "pv"
	multiPv  = "multipv"
)

// FromUSI USI から Go に変換するインターフェース
type FromUSI interface {
	// Piece USIの駒を int に変換する
	// TODO: Piece 型を作る
	Piece(string) (int32, error)

	// EngineID USI の name と author を変換する
	EngineID(string) (string, string, error)

	// Option USI の option を変換する
	// any で返しているのでよろしくない・・
	Option(string) (interface{}, error)

	// Move USI の pv を変換する
	Move(string) (*pb.Move, error)

	// Info USI の info を変換する
	Info(string) (*pb.Info, int, error)
}

type fromUSI struct{}

// NewFromUSI 新しい FromUSI を返す
func NewFromUSI() FromUSI {
	return fromUSI{}
}

func (fu fromUSI) Piece(s string) (i int32, e error) {
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
		e = exception.InvalidPieceID.WithMsg("PieceIDが不正です id=" + s)
	}
	return
}

// EngineID
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

// Option エンジンのオプションをパースする
// s string 一行のオプションのUSI文字列を受け取り、
// パースした結果の Option を返す。
// interface だが、 Button|Check|Select|Spin|String|FileName
// パースに失敗した場合はエラー。成功時 error は nil
func (fu fromUSI) Option(s string) (interface{}, error) {
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
func (fu fromUSI) parseButton(s string) (*pb.Button, error) {
	res := buttonRegex.FindAllStringSubmatch(s, -1)
	if len(res) == 0 || len(res[0]) < 2 {
		return nil, invalidSyntax(s, parseErrorButton)
	}

	return &pb.Button{Name: res[0][1]}, nil
}

// check type を Egn の Options にセットする
// option name <string> type check default <bool>
// このフォーマット以外は許容しない。default が無くてもてもだめ
func (fu fromUSI) parseCheck(s string) (*pb.Check, error) {
	res := checkRegex.FindAllStringSubmatch(s, -1)
	if len(res) == 0 || len(res[0]) < 3 {
		return nil, invalidSyntax(s, parseErrorCheck)
	}

	b := res[0][2] == "true"
	return &pb.Check{Name: res[0][1], Val: b, Default: b}, nil
}

// spin type を Egn の Options にセットする
// option name <string> type spin default <int> min <int> max <int>
// このフォーマット以外は許容しない
// 各値がなかったり、int ではない時、min > max の時はエラー
func (fu fromUSI) parseSpin(s string) (*pb.Spin, error) {
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

	init32, min32, max32 := int32(init), int32(min), int32(max)
	return &pb.Spin{Name: res[0][1], Val: init32, Default: init32, Min: min32, Max: max32}, nil
}

// select type を Egn の Options にセットする
// option name <string> type combo default <string> rep(var <string>)
// このフォーマット以外は許容しない
// initial がない、var がない、default が var にない時はエラー
func (fu fromUSI) parseSelect(s string) (*pb.Select, error) {
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

	return &pb.Select{Name: res[0][1], Val: init, Default: init, Vars: vars}, nil
}

// string type を Egn の Options にセットする
// option name <string> type string default <string>
func (fu fromUSI) parseString(s string) (*pb.String, error) {
	res := stringRegex.FindAllStringSubmatch(s, -1)
	if len(res) == 0 || len(res[0]) < 3 {
		return nil, invalidSyntax(s, parseErrorString)
	}

	return &pb.String{Name: res[0][1], Val: res[0][2], Default: res[0][2]}, nil
}

// option name <string> type filename default <string>
func (fu fromUSI) parseFileName(s string) (*pb.Filename, error) {
	res := fileNameRegex.FindAllStringSubmatch(s, -1)
	if len(res) == 0 || len(res[0]) < 3 {
		return nil, invalidSyntax(s, parseErrorFileName)
	}

	return &pb.Filename{Name: res[0][1], Val: res[0][2], Default: res[0][2]}, nil
}

func invalidSyntax(input, msg string) error {
	return exception.InvalidOptionSyntax.WithMsg(msg + "\nInput: " + input + "\n")
}

func (fu fromUSI) Move(s string) (m *pb.Move, err error) {
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

		return &pb.Move{
			Source:  &pb.Point{Row: -1, Column: -1},
			Dest:    &pb.Point{Row: fu.row(a[3]), Column: fu.column(a[2])},
			PieceID: piece,
		}, nil
	}

	return &pb.Move{
		Source:     &pb.Point{Row: fu.row(a[1]), Column: fu.column(a[0])},
		Dest:       &pb.Point{Row: fu.row(a[3]), Column: fu.column(a[2])},
		IsPromoted: len(a) == 5 && a[4] == "+",
	}, nil
}

func (fu fromUSI) column(s string) int32 {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	if i < 1 || i > 9 {
		panic(exception.InvalidColumnNumber)
	}
	return int32(i - 1) // 0-8 にする
}

func (fu fromUSI) row(s string) int32 {
	r := []rune(s)[0]
	if r < 97 || r > 105 {
		panic(exception.InvalidRowNumber)
	}
	return int32(r - 97)
}

// Info をパース(info string は渡さない)
// return
//   r *result.Info パースした結果。失敗したら nil
//   mpv int multipvならその値。multipvじゃなければ 0
//   err error エラー
func (fu fromUSI) Info(s string) (r *pb.Info, mpv int, err error) {
	a := strings.Split(s, " ")
	r = &pb.Info{Values: make(map[string]int32)}

	// panic をリカバーしてエラーをセット
	defer func() {
		if rec := recover(); rec != nil {
			err = exception.FailedToParseInfo.WithMsg(fmt.Sprintf("%v", rec))
		}
	}()

	i := 0
	for i < len(a) {
		switch strings.TrimSpace(a[i]) {
		case depth:
			i++
			r.Values[depth] = toInt32(a[i])
		case selDepth:
			i++
			r.Values[selDepth] = toInt32(a[i])
		case time:
			i++
			r.Values[time] = toInt32(a[i])
		case nodes:
			i++
			r.Values[nodes] = toInt32(a[i])
		case hashFull:
			i++
			r.Values[hashFull] = toInt32(a[i])
		case nps:
			i++
			r.Values[nps] = toInt32(a[i])
		case score:
			if a[i+1] == "cp" || a[i+1] == "mate" {
				r.Score = toInt32(a[i+2])
			}
			i += 2
		case multiPv:
			i++
			mpv = int(toInt32(a[i]))
		case pv:
			fu.setMoves(r, a[i+1:])
			i += len(a) // pv は 末尾
		}
		i++
	}

	return
}

func (fu fromUSI) setMoves(r *pb.Info, a []string) {
	m := make([]*pb.Move, len(a))
	for i, v := range a {
		mv, err := fu.Move(v)
		if err != nil {
			panic(err)
		}
		m[i] = mv
	}

	r.Moves = m
}

func toInt32(s string) int32 {
	i, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		panic(err)
	}
	return int32(i)
}
