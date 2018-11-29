// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package from_json

import (
	"testing"

	"github.com/murosan/shogi-proxy-server/app/domain/entity/shogi"
	"github.com/murosan/shogi-proxy-server/app/domain/exception"
)

func TestFromJson_Position(t *testing.T) {
	fj := NewFromJson()
	cases := []struct {
		in   []byte
		want shogi.Position
		err  *exception.Error
	}{
		{[]byte(`
{
  "pos": [
    [-2, -3, -4, -5, -8, 0, -4, -3, -2],
    [0, 0, 0, 0, 0, 0, -5, -6, 0],
    [-1, 0, -1, -1, -1, -1, 0, 0, -1],
    [0, 0, 0, 0, 0, 0, 7, 0, 0],
    [0, 0, 0, 0, 0, 0, 0, 0, 0],
    [0, -7, 1, 0, 0, 0, 0, 0, 0],
    [1, 0, 0, 1, 1, 1, 1, 0, 1],
    [0, 6, 5, 0, 0, 0, 0, 0, 0],
    [2, 3, 4, 0, 8, 5, 4, 3, 2]
  ],
  "cap_0": [3, 0, 0, 0, 0, 0, 0],
  "cap_1": [2, 0, 0, 0, 0, 0, 0],
  "turn": 1,
  "move_count": 100
}`),
			shogi.Position{
				Pos: [9][9]int{
					{-2, -3, -4, -5, -8, 0, -4, -3, -2},
					{0, 0, 0, 0, 0, 0, -5, -6, 0},
					{-1, 0, -1, -1, -1, -1, 0, 0, -1},
					{0, 0, 0, 0, 0, 0, 7, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, -7, 1, 0, 0, 0, 0, 0, 0},
					{1, 0, 0, 1, 1, 1, 1, 0, 1},
					{0, 6, 5, 0, 0, 0, 0, 0, 0},
					{2, 3, 4, 0, 8, 5, 4, 3, 2},
				},
				Cap0:      [7]int{3, 0, 0, 0, 0, 0, 0},
				Cap1:      [7]int{2, 0, 0, 0, 0, 0, 0},
				Turn:      1,
				MoveCount: 100,
			},
			nil,
		},
		{[]byte(`
{
  "pos": [
    [-2, -3, -4, -5, -8, 0, -4, -3, -2],
    [0, 0, 0, 0, 0, 0, -5, -6, 0],
    [-1, 0, -1, -1, -1, -1, 0, 0, -1],
    [0, 0, 0, 0, 0, 0, 7, 0, 0],
    [0, 0, 0, 0, 0, 0, 0, 0, 0],
    [0, -7, 1, 0, 0, 0, 0, 0, 0],
    [1, 0, 0, 1, 1, 1, 1, 0, 1],
    [0, 6, 5, 0, 0, 0, 0, 0, 0],
    [2, 3, 4, 0, 8, 5, 4, 3, 2]
  ],
  "cap_0": [3, 0, 0, 0, 0, 0, 0],
  "cap_1": [2, 0, 0, 0, 0, 0, 0],
  "turn": 1,
}`),
			shogi.Position{},
			exception.FailedToParseJson,
		},
	}

	for i, c := range cases {
		p, e := fj.Position(c.in)
		if e != nil && c.err.Code == c.err.Code {
			continue
		}
		if e != nil {
			t.Errorf("Error: %v\nIndex: %d", e.Error(), i)
		}
		if c.want.MoveCount != p.MoveCount {
			t.Errorf("MoveCount does not match.\nIndex: %d\nExpected: %v\nActual: %v", i, c.want.MoveCount, p.MoveCount)
		}
		if c.want.Turn != p.Turn {
			t.Errorf("Turn does not match.\nIndex: %d\nExpected: %v\nActual: %v", i, c.want.Turn, p.Turn)
		}
		if c.want.Cap0 != p.Cap0 {
			t.Errorf("Cap0 does not match.\nIndex: %d\nExpected: %v\nActual: %v", i, c.want.Cap0, p.Cap0)
		}
		if c.want.Cap1 != p.Cap1 {
			t.Errorf("Cap1 does not match.\nIndex: %d\nExpected: %v\nActual: %v", i, c.want.Cap1, p.Cap1)
		}
		if c.want.Pos != p.Pos {
			t.Errorf("Pos does not match.\nIndex: %d\nExpected: %v\nActual: %v", i, c.want.Pos, p.Pos)
		}
	}
}
