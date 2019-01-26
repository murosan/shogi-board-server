// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package converter

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"github.com/murosan/shogi-board-server/app/domain/entity/engine/option"
	"github.com/murosan/shogi-board-server/app/domain/entity/engine/result"
	"github.com/murosan/shogi-board-server/app/domain/entity/shogi"
	"github.com/murosan/shogi-board-server/app/domain/exception"
	testhelper "github.com/murosan/shogi-board-server/app/lib/test_helper"
)

var emp = ""

func TestFromUSI_Piece(t *testing.T) {
	cases := []struct {
		in      string
		want    int
		isError bool
	}{
		{shogi.UsiFu0, shogi.Fu0, false},
		{shogi.UsiFu1, shogi.Fu1, false},
		{shogi.UsiKyou0, shogi.Kyou0, false},
		{shogi.UsiKyou1, shogi.Kyou1, false},
		{shogi.UsiKei0, shogi.Kei0, false},
		{shogi.UsiKei1, shogi.Kei1, false},
		{shogi.UsiGin0, shogi.Gin0, false},
		{shogi.UsiGin1, shogi.Gin1, false},
		{shogi.UsiKin0, shogi.Kin0, false},
		{shogi.UsiKin1, shogi.Kin1, false},
		{shogi.UsiKaku0, shogi.Kaku0, false},
		{shogi.UsiKaku1, shogi.Kaku1, false},
		{shogi.UsiHisha0, shogi.Hisha0, false},
		{shogi.UsiHisha1, shogi.Hisha1, false},
		{shogi.UsiGyoku0, shogi.Gyoku0, false},
		{shogi.UsiGyoku1, shogi.Gyoku1, false},
		{shogi.UsiTo0, shogi.To0, false},
		{shogi.UsiTo1, shogi.To1, false},
		{shogi.UsiNariKyou0, shogi.NariKyou0, false},
		{shogi.UsiNariKyou1, shogi.NariKyou1, false},
		{shogi.UsiNariKei0, shogi.NariKei0, false},
		{shogi.UsiNariKei1, shogi.NariKei1, false},
		{shogi.UsiNariGin0, shogi.NariGin0, false},
		{shogi.UsiNariGin1, shogi.NariGin1, false},
		{shogi.UsiUma0, shogi.Uma0, false},
		{shogi.UsiUma1, shogi.Uma1, false},
		{shogi.UsiRyu0, shogi.Ryu0, false},
		{shogi.UsiRyu1, shogi.Ryu1, false},
		{"none", 0, true},
	}

	for i, c := range cases {
		r, e := NewFromUSI().Piece(c.in)

		if c.isError && e == nil {
			t.Errorf(`[FromUsi Piece] Expected error, but got nil
Index: %d
Input: %s`, i, c.in)
		}

		if !c.isError && e != nil {
			t.Errorf(`[FromUsi Piece] Expected nil, but got error.
Index: %d
Input: %s
Got:   %v`, i, c.in, e)
		}

		if c.want != r {
			t.Errorf(`[FromUsi Piece] The value was not as expected.
Index:    %d
Input:    %s
Expected: %d
Actual:   %d`, i, c.in, c.want, r)
		}
	}
}

func TestFromUSI_EngineID(t *testing.T) {
	fu := NewFromUSI()
	cases := []struct {
		in, name, val string
		err           error
	}{
		{"id name computer_name", name, "computer_name", nil},
		{"id name ", name, emp, nil},
		{"id name", emp, emp, exception.UnknownOption},
		{"id neimu typo_key", emp, emp, exception.UnknownOption},
		{"id author computer_author", author, "computer_author", nil},
		{"id author", emp, emp, exception.UnknownOption},
		{"id auther typo_key", emp, emp, exception.UnknownOption},
	}

	for _, c := range cases {
		n, v, e := fu.EngineID(c.in)
		if e != c.err {
			t.Errorf(`
Returned error was not as expected.
Input:    %v
Expected: %v
Actual:   %v`, c.in, c.err, e)
		}
		if n != c.name {
			t.Errorf(`
name of OptionId was not as expected.
Input:    %v
Expected: %v
Actual:   %v`, c.in, c.name, n)
		}
		if v != c.val {
			t.Errorf(`
Value of OptionId was not as expected.
Input:    %v
Expected: %v
Actual:   %v`, c.in, c.val, v)
		}
	}
}

func TestFromUSI_Option(t *testing.T) {
	fu := NewFromUSI()
	cases := []struct {
		in   string
		want *option.Button
		err  error
	}{
		{"option name ResetLearning type button", option.NewButton("ResetLearning"), nil},
		{"option name <empty> type button", option.NewButton("<empty>"), nil}, // まぁいい
		{"option name ResetLearning type button sur", nil, exception.InvalidOptionSyntax},
		{"option name 1 type button", option.NewButton("1"), nil},
	}
	for _, c := range cases {
		o, err := fu.Option(c.in)
		basicOptionMatching(t, c.in, o, c.want, err, c.err)
	}
}

func TestFromUSI_Option2(t *testing.T) {
	fu := NewFromUSI()
	cases := []struct {
		in   string
		want *option.Check
		err  error
	}{
		{"option name UseBook type check default true", option.NewCheck("UseBook", true, true), nil},
		{"   option name UseBook type check default true   ", option.NewCheck("UseBook", true, true), nil},
		{"option name UseBook type check default ", nil, exception.InvalidOptionSyntax},
		{"option name UseBook type check default not_bool", nil, exception.InvalidOptionSyntax},
		{"option name UseBook type check dlft true", nil, exception.InvalidOptionSyntax},
	}

	for _, c := range cases {
		o, err := fu.Option(c.in)
		basicOptionMatching(t, c.in, o, c.want, err, c.err)
	}
}

func TestFromUSI_Option3(t *testing.T) {
	fu := NewFromUSI()
	cases := []struct {
		in   string
		want *option.Spin
		err  error
	}{
		{
			"option name Selectivity type spin default 2 min 0 max 4",
			option.NewSpin("Selectivity", 2, 2, 0, 4),
			nil,
		},
		{
			"option name Selectivity type spin default -100 min -123456 max 54321 ",
			option.NewSpin("Selectivity", -100, -100, -123456, 54321),
			nil,
		},
		{
			"option name Selectivity type spin min 0 max 4",
			nil,
			exception.InvalidOptionSyntax,
		},
		{
			"option name Selectivity type spin default 2",
			nil,
			exception.InvalidOptionSyntax,
		},
		{
			"option name Selectivity type spin min 0 max 4 default 2",
			nil,
			exception.InvalidOptionSyntax,
		},
		{
			"option name Selectivity type spin default two min 0 max 4",
			nil,
			exception.InvalidOptionSyntax,
		},
	}
	for _, c := range cases {
		o, err := fu.Option(c.in)
		basicOptionMatching(t, c.in, o, c.want, err, c.err)
	}
}

func TestFromUSI_Option4(t *testing.T) {
	fu := NewFromUSI()
	cases := []struct {
		in   string
		want *option.Select
		err  error
	}{
		{
			"option name Style type combo default Normal var Solid var Normal var Risky",
			option.NewSelect("Style", "Normal", "Normal", []string{"Solid", "Normal", "Risky"}),
			nil,
		},
		{
			"option name Style type combo default Nor mal var Sol id var Nor mal var R isky",
			option.NewSelect("Style", "Nor mal", "Nor mal", []string{"Sol id", "Nor mal", "R isky"}),
			nil,
		},
		{"option name Style type combo default None var Solid var Normal var Risky",
			nil,
			exception.InvalidOptionSyntax,
		},
		{"option name Style type combo var Solid var Normal var Risky",
			nil,
			exception.InvalidOptionSyntax,
		},
		{"option name Style type combo default Normal",
			nil,
			exception.InvalidOptionSyntax,
		},
	}
	for _, c := range cases {
		o, err := fu.Option(c.in)
		basicOptionMatching(t, c.in, o, c.want, err, c.err)
	}
}

func TestFromUSI_Option5(t *testing.T) {
	fu := NewFromUSI()
	cases := []struct {
		in   string
		want *option.String
		err  error
	}{
		{"option name BookFile type string default public.bin",
			option.NewString("BookFile", "public.bin", "public.bin"),
			nil,
		},
		{"option name BookFile type string default public.bin var a",
			option.NewString("BookFile", "public.bin var a", "public.bin var a"),
			nil,
		},
		{"option name BookFile type string",
			nil,
			exception.InvalidOptionSyntax,
		},
		{"option name BookFile type string public.bin",
			nil,
			exception.InvalidOptionSyntax,
		},
	}
	for _, c := range cases {
		o, err := fu.Option(c.in)
		basicOptionMatching(t, c.in, o, c.want, err, c.err)
	}
}

func TestFromUSI_Option6(t *testing.T) {
	fu := NewFromUSI()
	cases := []struct {
		in   string
		want *option.FileName
		err  error
	}{
		{
			"option name LearningFile type filename default <empty>",
			option.NewFileName("LearningFile", "<empty>", "<empty>"),
			nil,
		},
		{"option name LearningFile type filename default <empty> var a",
			option.NewFileName("LearningFile", "<empty> var a", "<empty> var a"),
			nil,
		},
		{"option name LearningFile type filename",
			nil,
			exception.InvalidOptionSyntax,
		},
		{"option name LearningFile type filename <empty>",
			nil,
			exception.InvalidOptionSyntax,
		},
	}
	for _, c := range cases {
		o, err := fu.Option(c.in)
		basicOptionMatching(t, c.in, o, c.want, err, c.err)
	}
}

// in: input
// o1: Returned Option
// o2: Expected Option
// e1: Returned Error
// e2: Expected Error
func basicOptionMatching(t *testing.T, in string, o1, o2 option.Option, e1, e2 error) {
	t.Helper()
	if (e1 == nil && e2 != nil) || (e1 != nil && e2 == nil) {
		t.Errorf(`
Returned error was not as expected.
Input:    %v
Expected: %v
Actual:   %v`, in, e2, e1)
	}

	// 予想通りのエラーが返った
	if e1 != nil && strings.Contains(string(e1.Error()), string(e2.Error())) {
		return
	}

	// エラーは返ったが、想定と違った
	if e1 != nil && !strings.Contains(string(e1.Error()), string(e2.Error())) {
		t.Errorf(`
Returned error was not as expected.
Input:    %v
Expected: %v
Actual:   %v`, in, e2, e1)
	}

	// USIコマンドが想定通りかどうか
	usi1 := o1.Usi()
	usi2 := o2.Usi()
	if usi1 != usi2 {
		t.Errorf(`
Update value was not as expected.
Input:    %v
Expected: %v
Actual:   %v`, in, usi2, usi1)
	}

	// json化した値が同等かどうか
	json1, _ := json.MarshalIndent(o1, "", "  ")
	json2, _ := json.MarshalIndent(o2, "", "  ")
	if !bytes.Equal(json1, json2) {
		t.Errorf(`
Marshaled value (json bytes) was not as expected.
Input:    %s
Expected: %s
Actual:   %s
`, in, string(json2), string(json1))
	}
}

func TestFromUSI_Move(t *testing.T) {
	cases := []struct {
		in   string
		want shogi.Move
		err  error
	}{
		{"7g7f",
			shogi.Move{
				Source:     shogi.Point{Row: 6, Column: 6},
				Dest:       shogi.Point{Row: 5, Column: 6},
				PieceID:    0,
				IsPromoted: false,
			},
			nil,
		},
		{"8h2b+",
			shogi.Move{
				Source:     shogi.Point{Row: 7, Column: 7},
				Dest:       shogi.Point{Row: 1, Column: 1},
				PieceID:    0,
				IsPromoted: true},
			nil},
		{"G*5b",
			shogi.Move{
				Source:  shogi.Point{Row: -1, Column: -1},
				Dest:    shogi.Point{Row: 1, Column: 4},
				PieceID: 5, IsPromoted: false,
			},
			nil,
		},
		{"s*5b",
			shogi.Move{
				Source:     shogi.Point{Row: -1, Column: -1},
				Dest:       shogi.Point{Row: 1, Column: 4},
				PieceID:    -4,
				IsPromoted: false,
			},
			nil,
		},
		{"", shogi.Move{}, exception.UnknownCharacter},
		{"7g7z", shogi.Move{}, exception.UnknownCharacter},
		{"7g7$", shogi.Move{}, exception.UnknownCharacter},
		{"0g7a", shogi.Move{}, exception.UnknownCharacter},
		{"1x7a", shogi.Move{}, exception.UnknownCharacter},
		{"G*vb", shogi.Move{}, exception.UnknownCharacter},
		{"G*4z", shogi.Move{}, exception.UnknownCharacter},
		{"A*7a", shogi.Move{}, exception.UnknownCharacter},
	}

	for i, c := range cases {
		moveHelper(t, i, c.in, c.want, c.err)
	}
}

func moveHelper(t *testing.T, i int, in string, want shogi.Move, err error) {
	t.Helper()
	res, e := NewFromUSI().Move(in)
	msg := ""

	if (e == nil && err != nil) || (e != nil && err == nil) {
		msg = "Expected error, but was not as expected."
		moveErrorPrintHelper(t, i, msg, in, err, e)
	}

	// 想定通りのエラー
	if e != nil && strings.Contains(string(e.Error()), string(err.Error())) {
		return
	}

	// エラーだったが、想定と違った。
	if e != nil && !strings.Contains(string(e.Error()), string(err.Error())) {
		msg = "Expected error, but was not as expected."
		moveErrorPrintHelper(t, i, msg, in, err, e)
	}

	if !testhelper.MoveEquals(res, want) {
		msg = "The value was not as expected."
		moveErrorPrintHelper(t, i, msg, in, want, res)
	}
}

func moveErrorPrintHelper(
	t *testing.T,
	i int,
	msg, in string,
	expected, actual interface{},
) {
	t.Helper()
	t.Errorf(`(From Usi: Parse Move) %s
Index:    %d
Input:    %s
Expected: %v
Actual:   %v
`, msg, i, in, expected, actual)
}

func TestFromUSI_Info(t *testing.T) {
	cases := []struct {
		in   string
		want result.Info
		mpv  int
		err  error
	}{
		{
			"info time 1141 depth 3 seldepth 3 nodes 135125 score cp -1521 pv 3a3b L*4h 4c4d",
			result.Info{
				Values: map[string]int{
					result.Time:     1141,
					result.Depth:    3,
					result.SelDepth: 3,
					result.Nodes:    135125,
				},
				Score: -1521,
				Moves: []shogi.Move{
					{
						Source:     shogi.Point{Column: 2, Row: 0},
						Dest:       shogi.Point{Column: 2, Row: 1},
						PieceID:    0,
						IsPromoted: false,
					},
					{
						Source:     shogi.Point{Column: -1, Row: -1},
						Dest:       shogi.Point{Column: 3, Row: 7},
						PieceID:    2,
						IsPromoted: false,
					},
					{
						Source:     shogi.Point{Column: 3, Row: 2},
						Dest:       shogi.Point{Column: 3, Row: 3},
						PieceID:    0,
						IsPromoted: false,
					},
				},
			}, 0, nil},
		{
			"info nodes 120000 nps 116391 hashfull 104",
			result.Info{
				Values: map[string]int{
					result.Nodes:    120000,
					result.Nps:      116391,
					result.HashFull: 104,
				},
				Score: 0,
				Moves: []shogi.Move{},
			}, 0, nil},
		{
			"info score cp 156 multipv 1 pv P*5h 4g5g 5h5g 8b8f",
			result.Info{
				Values: map[string]int{},
				Score:  156,
				Moves: []shogi.Move{
					{
						Source:     shogi.Point{Column: -1, Row: -1},
						Dest:       shogi.Point{Column: 4, Row: 7},
						PieceID:    1,
						IsPromoted: false,
					},
					{
						Source:     shogi.Point{Column: 3, Row: 6},
						Dest:       shogi.Point{Column: 4, Row: 6},
						PieceID:    0,
						IsPromoted: false,
					},
					{
						Source:     shogi.Point{Column: 4, Row: 7},
						Dest:       shogi.Point{Column: 4, Row: 6},
						PieceID:    0,
						IsPromoted: false,
					},
					{
						Source:     shogi.Point{Column: 7, Row: 1},
						Dest:       shogi.Point{Column: 7, Row: 5},
						PieceID:    0,
						IsPromoted: false,
					},
				},
			}, 1, nil},
		{
			"info score cp -99 multipv 2 pv 2d4d 3c4e 8h5e N*7f",
			result.Info{
				Values: map[string]int{},
				Score:  -99,
				Moves: []shogi.Move{
					{
						Source: shogi.Point{Column: 1, Row: 3},
						Dest:   shogi.Point{Column: 3, Row: 3}, PieceID: 0,
						IsPromoted: false,
					},
					{
						Source: shogi.Point{Column: 2, Row: 2},
						Dest:   shogi.Point{Column: 3, Row: 4}, PieceID: 0,
						IsPromoted: false,
					},
					{
						Source: shogi.Point{Column: 7, Row: 7},
						Dest:   shogi.Point{Column: 4, Row: 4}, PieceID: 0,
						IsPromoted: false,
					},
					{
						Source: shogi.Point{Column: -1, Row: -1},
						Dest:   shogi.Point{Column: 6, Row: 5}, PieceID: 3,
						IsPromoted: false,
					},
				},
			}, 2, nil},
		{
			"info score cp -157 multipv 3 pv 5g5f 4g4f 4e3c+ 4c3c",
			result.Info{
				Values: map[string]int{},
				Score:  -157,
				Moves: []shogi.Move{
					{
						Source:     shogi.Point{Column: 4, Row: 6},
						Dest:       shogi.Point{Column: 4, Row: 5},
						PieceID:    0,
						IsPromoted: false,
					},
					{
						Source:     shogi.Point{Column: 3, Row: 6},
						Dest:       shogi.Point{Column: 3, Row: 5},
						PieceID:    0,
						IsPromoted: false,
					},
					{
						Source:     shogi.Point{Column: 3, Row: 4},
						Dest:       shogi.Point{Column: 2, Row: 2},
						PieceID:    0,
						IsPromoted: true,
					},

					{
						Source:     shogi.Point{Column: 3, Row: 2},
						Dest:       shogi.Point{Column: 2, Row: 2},
						PieceID:    0,
						IsPromoted: false,
					},
				},
			}, 3, nil},
		{
			"info score cp -157 str multipv 3 lalala... pv 5g5f 4g4f 4e3c+ 4c3c",
			result.Info{
				Values: map[string]int{},
				Score:  -157,
				Moves: []shogi.Move{
					{
						Source:     shogi.Point{Column: 4, Row: 6},
						Dest:       shogi.Point{Column: 4, Row: 5},
						PieceID:    0,
						IsPromoted: false,
					},
					{
						Source:     shogi.Point{Column: 3, Row: 6},
						Dest:       shogi.Point{Column: 3, Row: 5},
						PieceID:    0,
						IsPromoted: false,
					},
					{
						Source:     shogi.Point{Column: 3, Row: 4},
						Dest:       shogi.Point{Column: 2, Row: 2},
						PieceID:    0,
						IsPromoted: true,
					},

					{
						Source:     shogi.Point{Column: 3, Row: 2},
						Dest:       shogi.Point{Column: 2, Row: 2},
						PieceID:    0,
						IsPromoted: false,
					},
				},
			},
			3, nil},
		{
			"info score cp -225 multipv 4 pv 5g6h 8b8f P*8g 8f5f",
			result.Info{
				Values: map[string]int{},
				Score:  -225,
				Moves: []shogi.Move{
					{
						Source:     shogi.Point{Column: 4, Row: 6},
						Dest:       shogi.Point{Column: 5, Row: 7},
						PieceID:    0,
						IsPromoted: false,
					},
					{
						Source:     shogi.Point{Column: 7, Row: 1},
						Dest:       shogi.Point{Column: 7, Row: 5},
						PieceID:    0,
						IsPromoted: false,
					},
					{
						Source:     shogi.Point{Column: -1, Row: -1},
						Dest:       shogi.Point{Column: 7, Row: 6},
						PieceID:    1,
						IsPromoted: false,
					},
					{
						Source:     shogi.Point{Column: 7, Row: 5},
						Dest:       shogi.Point{Column: 4, Row: 5},
						PieceID:    0,
						IsPromoted: false,
					},
				},
			}, 4, nil},
		{
			"info score cp aaa multipv 4 pv 5g6h 8b8f P*8g 8f5f",
			result.Info{},
			0,
			exception.FailedToParseInfo,
		},
		{
			"info score cp 4 multipv 4 pv 5g6h 8b8f P*8g 8f5z",
			result.Info{},
			0,
			exception.FailedToParseInfo,
		},
	}

	for i, c := range cases {
		infoHelper(t, i, c.in, c.want, c.mpv, c.err)
	}
}

func infoHelper(
	t *testing.T,
	i int,
	in string,
	want result.Info,
	mpv int,
	err error,
) {
	t.Helper()
	msg := ""
	res, mpvKey, e := NewFromUSI().Info(in)

	if (e == nil && err != nil) || (e != nil && err == nil) {
		msg = "Expected error, but was not as expected."
		infoErrorPrintHelper(t, i, msg, in, err, e)
	}

	// 想定通りのエラー
	if e != nil && err != nil && strings.Contains(string(e.Error()), string(err.Error())) {
		return
	}

	// エラーだったが、想定と違った。
	if e != nil && err != nil && !strings.Contains(string(e.Error()), string(err.Error())) {
		msg = "Expected error, but was not as expected."
		infoErrorPrintHelper(t, i, msg, in, err, e)
	}

	if mpvKey != mpv {
		msg = "The multipv index value was not as expected."
		infoErrorPrintHelper(t, i, msg, in, mpv, mpvKey)
	}

	if !testhelper.InfoEquals(res, want) {
		msg = "The value was not as expected."
		infoErrorPrintHelper(t, i, msg, in, want, res)
	}
}

func infoErrorPrintHelper(t *testing.T,
	i int,
	msg, in string,
	expected, actual interface{},
) {
	t.Helper()
	t.Errorf(`(From Usi: Parse Info) %s
Index:    %d
Input:    %s
Expected: %v
Actual:   %v
`, msg, i, in, expected, actual)
}