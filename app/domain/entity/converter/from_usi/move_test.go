// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package from_usi

import (
	"strings"
	"testing"

	"github.com/murosan/shogi-proxy-server/app/domain/entity/shogi"
	"github.com/murosan/shogi-proxy-server/app/domain/exception"
	"github.com/murosan/shogi-proxy-server/app/lib/test_helper"
)

func TestFromUsi_Move(t *testing.T) {
	cases := []struct {
		in   string
		want *shogi.Move
		err  error
	}{
		{"7g7f", &shogi.Move{
			Source: []int{6, 6}, Dest: []int{6, 5}, PieceId: 0, IsPromoted: false},
			nil},
		{"8h2b+", &shogi.Move{
			Source: []int{7, 7}, Dest: []int{1, 1}, PieceId: 0, IsPromoted: true},
			nil},
		{"G*5b", &shogi.Move{
			Source: []int{-1, -1}, Dest: []int{4, 1}, PieceId: 5, IsPromoted: false},
			nil},
		{"s*5b", &shogi.Move{
			Source: []int{-1, -1}, Dest: []int{4, 1}, PieceId: -4, IsPromoted: false},
			nil},
		{"", nil, exception.UnknownCharacter},
		{"7g7z", nil, exception.UnknownCharacter},
		{"7g7$", nil, exception.UnknownCharacter},
		{"0g7a", nil, exception.UnknownCharacter},
		{"1x7a", nil, exception.UnknownCharacter},
		{"G*vb", nil, exception.UnknownCharacter},
		{"G*4z", nil, exception.UnknownCharacter},
		{"A*7a", nil, exception.UnknownCharacter},
	}

	for i, c := range cases {
		moveHelper(t, i, c.in, c.want, c.err)
	}
}

func moveHelper(t *testing.T, i int, in string, want *shogi.Move, err error) {
	t.Helper()
	res, e := NewFromUsi().Move(in)
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

	if (res == nil || want == nil) && res != want {
		msg = "The value was not as expected."
		moveErrorPrintHelper(t, i, msg, in, want, res)
	}

	if !test_helper.MoveEquals(*res, *want) {
		msg = "The value was not as expected."
		moveErrorPrintHelper(t, i, msg, in, want, res)
	}
}

func moveErrorPrintHelper(t *testing.T, i int, msg, in string, expected, actual interface{}) {
	t.Helper()
	t.Errorf(`(From Usi: Parse Move) %s
Index:    %d
Input:    %s
Expected: %v
Actual:   %v
`, msg, i, in, expected, actual)
}
