// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package from_usi

import (
	"strings"
	"testing"

	"github.com/murosan/shogi-proxy-server/app/domain/entity/shogi"
	"github.com/murosan/shogi-proxy-server/app/domain/exception"
	"github.com/murosan/shogi-proxy-server/app/lib/intutil"
)

func TestFromUsi_Move(t *testing.T) {
	cases := []struct {
		in   string
		want *shogi.Move
		err  error
	}{
		{"7g7f", &shogi.Move{
			Source: []int{6, 6}, Dest: []int{6, 5}, PieceId: 0, Extra: shogi.None},
			nil},
		{"8h2b+", &shogi.Move{
			Source: []int{7, 7}, Dest: []int{1, 1}, PieceId: 0, Extra: shogi.Promote},
			nil},
		{"G*5b", &shogi.Move{
			Source: []int{-1, -1}, Dest: []int{4, 1}, PieceId: 5, Extra: shogi.FromCaptured},
			nil},
		{"s*5b", &shogi.Move{
			Source: []int{-1, -1}, Dest: []int{4, 1}, PieceId: -4, Extra: shogi.FromCaptured},
			nil},
		{"", nil, exception.UnknownCharacter},
		{"7g7z", nil, exception.UnknownCharacter},
		{"7g7z$", nil, exception.UnknownCharacter},
		{"0g7a", nil, exception.UnknownCharacter},
		{"1x7a", nil, exception.UnknownCharacter},
		{"G*vb", nil, exception.UnknownCharacter},
		{"G*4z", nil, exception.UnknownCharacter},
		{"A*7a", nil, exception.UnknownCharacter},
		{"7g7fh", nil, exception.UnknownCharacter},
	}

	for i, c := range cases {
		moveHelper(t, i, c.in, c.want, c.err)
	}
}

func moveHelper(t *testing.T, i int, in string, want *shogi.Move, err error) {
	t.Helper()

	res, e := NewFromUsi().Move(in)

	if (e == nil && err != nil) || (e != nil && err == nil) {
		t.Errorf(`(From Usi: Paese Move) Expected error, but was not as expected.
Index:    %d
Input:    %s
Expected: %v
Actual:   %v
`, i, in, err, e)
	}

	// 想定通りのエラー
	if e != nil && strings.Contains(string(e.Error()), string(err.Error())) {
		return
	}

	// エラーだったが、想定と違った。
	if e != nil && !strings.Contains(string(e.Error()), string(err.Error())) {
		t.Errorf(`(From Usi: Paese Move) Expected error, but was not as expected.
Index:    %d
Input:    %s
Expected: %v
Actual:   %v
`, i, in, err, e)
	}

	if !moveEquals(res, want) {
		t.Errorf(`(From Usi: Parse Move) The value was not as expected.
Index:    %d
Input:    %s
Expected: %v
Actual:   %v
`, i, in, want, res)
	}
}

func moveEquals(a, b *shogi.Move) bool {
	if a == nil || b == nil {
		return false
	}
	return intutil.SliceEquals(a.Source, b.Source) &&
		intutil.SliceEquals(a.Dest, b.Dest) &&
		a.PieceId == b.PieceId &&
		a.Extra == b.Extra
}
