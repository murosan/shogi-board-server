package parse

import (
	"testing"

	"github.com/murosan/shogi-board-server/app/domain/entity/shogi"
	"github.com/murosan/shogi-board-server/app/domain/entity/usi"
)

func TestPiece(t *testing.T) {
	cases := []struct {
		in      usi.Piece
		want    shogi.Piece
		isError bool
	}{
		{in: usi.Fu0, want: shogi.Fu0, isError: false},
		{in: usi.Fu1, want: shogi.Fu1, isError: false},
		{in: usi.Kyou0, want: shogi.Kyou0, isError: false},
		{in: usi.Kyou1, want: shogi.Kyou1, isError: false},
		{in: usi.Kei0, want: shogi.Kei0, isError: false},
		{in: usi.Kei1, want: shogi.Kei1, isError: false},
		{in: usi.Gin0, want: shogi.Gin0, isError: false},
		{in: usi.Gin1, want: shogi.Gin1, isError: false},
		{in: usi.Kin0, want: shogi.Kin0, isError: false},
		{in: usi.Kin1, want: shogi.Kin1, isError: false},
		{in: usi.Kaku0, want: shogi.Kaku0, isError: false},
		{in: usi.Kaku1, want: shogi.Kaku1, isError: false},
		{in: usi.Hisha0, want: shogi.Hisha0, isError: false},
		{in: usi.Hisha1, want: shogi.Hisha1, isError: false},
		{in: usi.Gyoku0, want: shogi.Gyoku0, isError: false},
		{in: usi.Gyoku1, want: shogi.Gyoku1, isError: false},
		{in: usi.To0, want: shogi.To0, isError: false},
		{in: usi.To1, want: shogi.To1, isError: false},
		{in: usi.NariKyou0, want: shogi.NariKyou0, isError: false},
		{in: usi.NariKyou1, want: shogi.NariKyou1, isError: false},
		{in: usi.NariKei0, want: shogi.NariKei0, isError: false},
		{in: usi.NariKei1, want: shogi.NariKei1, isError: false},
		{in: usi.NariGin0, want: shogi.NariGin0, isError: false},
		{in: usi.NariGin1, want: shogi.NariGin1, isError: false},
		{in: usi.Uma0, want: shogi.Uma0, isError: false},
		{in: usi.Uma1, want: shogi.Uma1, isError: false},
		{in: usi.Ryu0, want: shogi.Ryu0, isError: false},
		{in: usi.Ryu1, want: shogi.Ryu1, isError: false},
		{in: "none", want: 0, isError: true},
	}

	for i, c := range cases {
		r, e := Piece(c.in)

		if c.isError && e == nil {
			t.Errorf(`[Piece] Expected error, but got nil.
Index: %d
Input: %s`, i, c.in)
		}

		if !c.isError && e != nil {
			t.Errorf(`[Piece] Expected nil, but got error.
Index: %d
Input: %s
Got:   %v`, i, c.in, e)
		}

		if c.want != r {
			t.Errorf(`[Piece] The value was not as expected.
Index:    %d
Input:    %s
Expected: %d
Actual:   %d`, i, c.in, c.want, r)
		}
	}
}
