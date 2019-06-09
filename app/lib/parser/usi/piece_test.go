package usi

import (
	"testing"

	"github.com/murosan/shogi-board-server/app/domain/model/shogi"
)

func TestParsePiece(t *testing.T) {
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
		r, e := ParsePiece(c.in)

		if c.isError && e == nil {
			t.Errorf(`[ParsePiece] Expected error, but got nil.
Index: %d
Input: %s`, i, c.in)
		}

		if !c.isError && e != nil {
			t.Errorf(`[ParsePiece] Expected nil, but got error.
Index: %d
Input: %s
Got:   %v`, i, c.in, e)
		}

		if c.want != r {
			t.Errorf(`[ParsePiece] The value was not as expected.
Index:    %d
Input:    %s
Expected: %d
Actual:   %d`, i, c.in, c.want, r)
		}
	}
}
