// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package from_usi

import (
	"strings"
	"testing"

	"github.com/murosan/shogi-proxy-server/app/domain/entity/engine/result"
	"github.com/murosan/shogi-proxy-server/app/domain/entity/shogi"
	"github.com/murosan/shogi-proxy-server/app/domain/exception"
	"github.com/murosan/shogi-proxy-server/app/lib/test_helper"
)

func TestFromUsi_Info(t *testing.T) {
	cases := []struct {
		in   string
		want *result.Info
		mpv  int
		err  error
	}{
		{"info time 1141 depth 3 seldepth 3 nodes 135125 score cp -1521 pv 3a3b L*4h 4c4d",
			&result.Info{
				Values: map[string]int{
					result.Time:     1141,
					result.Depth:    3,
					result.SelDepth: 3,
					result.Nodes:    135125,
				},
				Score: -1521,
				Moves: []shogi.Move{
					{[]int{2, 0}, []int{2, 1}, 0, false},
					{[]int{-1, -1}, []int{3, 7}, 2, false},
					{[]int{3, 2}, []int{3, 3}, 0, false},
				},
			}, 0, nil},
		{"info nodes 120000 nps 116391 hashfull 104",
			&result.Info{
				Values: map[string]int{
					result.Nodes:    120000,
					result.Nps:      116391,
					result.HashFull: 104,
				},
				Score: 0,
				Moves: []shogi.Move{},
			}, 0, nil},
		{"info score cp 156 multipv 1 pv P*5h 4g5g 5h5g 8b8f",
			&result.Info{
				Values: map[string]int{},
				Score:  156,
				Moves: []shogi.Move{
					{[]int{-1, -1}, []int{4, 7}, 1, false},
					{[]int{3, 6}, []int{4, 6}, 0, false},
					{[]int{4, 7}, []int{4, 6}, 0, false},
					{[]int{7, 1}, []int{7, 5}, 0, false},
				},
			}, 1, nil},
		{"info score cp -99 multipv 2 pv 2d4d 3c4e 8h5e N*7f",
			&result.Info{
				Values: map[string]int{},
				Score:  -99,
				Moves: []shogi.Move{
					{[]int{1, 3}, []int{3, 3}, 0, false},
					{[]int{2, 2}, []int{3, 4}, 0, false},
					{[]int{7, 7}, []int{4, 4}, 0, false},
					{[]int{-1, -1}, []int{6, 5}, 3, false},
				},
			}, 2, nil},
		{"info score cp -157 multipv 3 pv 5g5f 4g4f 4e3c+ 4c3c",
			&result.Info{
				Values: map[string]int{},
				Score:  -157,
				Moves: []shogi.Move{
					{[]int{4, 6}, []int{4, 5}, 0, false},
					{[]int{3, 6}, []int{3, 5}, 0, false},
					{[]int{3, 4}, []int{2, 2}, 0, true},
					{[]int{3, 2}, []int{2, 2}, 0, false},
				},
			}, 3, nil},
		{"info score cp -157 str multipv 3 lalala... pv 5g5f 4g4f 4e3c+ 4c3c",
			&result.Info{
				Values: map[string]int{},
				Score:  -157,
				Moves: []shogi.Move{
					{[]int{4, 6}, []int{4, 5}, 0, false},
					{[]int{3, 6}, []int{3, 5}, 0, false},
					{[]int{3, 4}, []int{2, 2}, 0, true},
					{[]int{3, 2}, []int{2, 2}, 0, false},
				},
			}, 3, nil},
		{"info score cp -225 multipv 4 pv 5g6h 8b8f P*8g 8f5f",
			&result.Info{
				Values: map[string]int{},
				Score:  -225,
				Moves: []shogi.Move{
					{[]int{4, 6}, []int{5, 7}, 0, false},
					{[]int{7, 1}, []int{7, 5}, 0, false},
					{[]int{-1, -1}, []int{7, 6}, 1, false},
					{[]int{7, 5}, []int{4, 5}, 0, false},
				},
			}, 4, nil},
		{"info score cp aaa multipv 4 pv 5g6h 8b8f P*8g 8f5f",
			nil, 0, exception.FailedToParseInfo},
		{"info score cp 4 multipv 4 pv 5g6h 8b8f P*8g 8f5z",
			nil, 0, exception.FailedToParseInfo},
	}

	for i, c := range cases {
		infoHelper(t, i, c.in, c.want, c.mpv, c.err)
	}
}

func infoHelper(t *testing.T, i int, in string, want *result.Info, mpv int, err error) {
	t.Helper()
	msg := ""
	res, mpvKey, e := NewFromUsi().Info(in)

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

	if !test_helper.InfoEquals(res, want) {
		msg = "The value was not as expected."
		infoErrorPrintHelper(t, i, msg, in, want, res)
	}
}

func infoErrorPrintHelper(t *testing.T, i int, msg, in string, expected, actual interface{}) {
	t.Helper()
	t.Errorf(`(From Usi: Parse Info) %s
Index:    %d
Input:    %s
Expected: %v
Actual:   %v
`, msg, i, in, expected, actual)
}
