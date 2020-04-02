package convert

import (
	"testing"

	"github.com/murosan/shogi-board-server/app/domain/entity/shogi"
	"github.com/murosan/shogi-board-server/app/domain/entity/usi"
)

func TestPieceToUSI(t *testing.T) {
	cases := []struct {
		in      shogi.Piece
		want    usi.Piece
		isError bool
	}{
		{in: shogi.Fu0, want: usi.Fu0, isError: false},
		{in: shogi.Fu1, want: usi.Fu1, isError: false},
		{in: shogi.Kyou0, want: usi.Kyou0, isError: false},
		{in: shogi.Kyou1, want: usi.Kyou1, isError: false},
		{in: shogi.Kei0, want: usi.Kei0, isError: false},
		{in: shogi.Kei1, want: usi.Kei1, isError: false},
		{in: shogi.Gin0, want: usi.Gin0, isError: false},
		{in: shogi.Gin1, want: usi.Gin1, isError: false},
		{in: shogi.Kin0, want: usi.Kin0, isError: false},
		{in: shogi.Kin1, want: usi.Kin1, isError: false},
		{in: shogi.Kaku0, want: usi.Kaku0, isError: false},
		{in: shogi.Kaku1, want: usi.Kaku1, isError: false},
		{in: shogi.Hisha0, want: usi.Hisha0, isError: false},
		{in: shogi.Hisha1, want: usi.Hisha1, isError: false},
		{in: shogi.Gyoku0, want: usi.Gyoku0, isError: false},
		{in: shogi.Gyoku1, want: usi.Gyoku1, isError: false},
		{in: shogi.To0, want: usi.To0, isError: false},
		{in: shogi.To1, want: usi.To1, isError: false},
		{in: shogi.NariKyou0, want: usi.NariKyou0, isError: false},
		{in: shogi.NariKyou1, want: usi.NariKyou1, isError: false},
		{in: shogi.NariKei0, want: usi.NariKei0, isError: false},
		{in: shogi.NariKei1, want: usi.NariKei1, isError: false},
		{in: shogi.NariGin0, want: usi.NariGin0, isError: false},
		{in: shogi.NariGin1, want: usi.NariGin1, isError: false},
		{in: shogi.Uma0, want: usi.Uma0, isError: false},
		{in: shogi.Uma1, want: usi.Uma1, isError: false},
		{in: shogi.Ryu0, want: usi.Ryu0, isError: false},
		{in: shogi.Ryu1, want: usi.Ryu1, isError: false},
		{in: 0, want: "", isError: true},
	}

	for i, c := range cases {
		r, e := Piece(c.in)

		if c.isError && e == nil {
			t.Errorf(`[PieceToUSI] Expected error, but got nil
Index: %d
Input: %d`, i, c.in)
		}

		if !c.isError && e != nil {
			t.Errorf(`[PieceToUSI] Expected nil, but got error.
Index: %d
Input: %d
Got:   %v`, i, c.in, e)
		}

		if c.want != r {
			t.Errorf(`[PieceToUSI] The value was not as expected.
Index:    %d
Input:    %d
Expected: %s
Actual:   %s`, i, c.in, c.want, r)
		}
	}
}
