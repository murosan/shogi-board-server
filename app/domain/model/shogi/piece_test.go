package shogi

import "testing"

func TestPieceToUSI(t *testing.T) {
	cases := []struct {
		in      int
		want    string
		isError bool
	}{
		{Fu0, UsiFu0, false},
		{Fu1, UsiFu1, false},
		{Kyou0, UsiKyou0, false},
		{Kyou1, UsiKyou1, false},
		{Kei0, UsiKei0, false},
		{Kei1, UsiKei1, false},
		{Gin0, UsiGin0, false},
		{Gin1, UsiGin1, false},
		{Kin0, UsiKin0, false},
		{Kin1, UsiKin1, false},
		{Kaku0, UsiKaku0, false},
		{Kaku1, UsiKaku1, false},
		{Hisha0, UsiHisha0, false},
		{Hisha1, UsiHisha1, false},
		{Gyoku0, UsiGyoku0, false},
		{Gyoku1, UsiGyoku1, false},
		{To0, UsiTo0, false},
		{To1, UsiTo1, false},
		{NariKyou0, UsiNariKyou0, false},
		{NariKyou1, UsiNariKyou1, false},
		{NariKei0, UsiNariKei0, false},
		{NariKei1, UsiNariKei1, false},
		{NariGin0, UsiNariGin0, false},
		{NariGin1, UsiNariGin1, false},
		{Uma0, UsiUma0, false},
		{Uma1, UsiUma1, false},
		{Ryu0, UsiRyu0, false},
		{Ryu1, UsiRyu1, false},
		{0, "", true},
	}

	for i, c := range cases {
		r, e := PieceToUSI(c.in)

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
