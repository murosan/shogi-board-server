// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package to_usi

import (
	"testing"

	"github.com/murosan/shogi-proxy-server/app/domain/entity/shogi"
)

func TestFromUsi_Piece(t *testing.T) {
	cases := []struct {
		in      int
		want    string
		isError bool
	}{
		{shogi.Fu0, shogi.UsiFu0, false},
		{shogi.Fu1, shogi.UsiFu1, false},
		{shogi.Kyou0, shogi.UsiKyou0, false},
		{shogi.Kyou1, shogi.UsiKyou1, false},
		{shogi.Kei0, shogi.UsiKei0, false},
		{shogi.Kei1, shogi.UsiKei1, false},
		{shogi.Gin0, shogi.UsiGin0, false},
		{shogi.Gin1, shogi.UsiGin1, false},
		{shogi.Kin0, shogi.UsiKin0, false},
		{shogi.Kin1, shogi.UsiKin1, false},
		{shogi.Kaku0, shogi.UsiKaku0, false},
		{shogi.Kaku1, shogi.UsiKaku1, false},
		{shogi.Hisha0, shogi.UsiHisha0, false},
		{shogi.Hisha1, shogi.UsiHisha1, false},
		{shogi.Gyoku0, shogi.UsiGyoku0, false},
		{shogi.Gyoku1, shogi.UsiGyoku1, false},
		{shogi.To0, shogi.UsiTo0, false},
		{shogi.To1, shogi.UsiTo1, false},
		{shogi.NariKyou0, shogi.UsiNariKyou0, false},
		{shogi.NariKyou1, shogi.UsiNariKyou1, false},
		{shogi.NariKei0, shogi.UsiNariKei0, false},
		{shogi.NariKei1, shogi.UsiNariKei1, false},
		{shogi.NariGin0, shogi.UsiNariGin0, false},
		{shogi.NariGin1, shogi.UsiNariGin1, false},
		{shogi.Uma0, shogi.UsiUma0, false},
		{shogi.Uma1, shogi.UsiUma1, false},
		{shogi.Ryu0, shogi.UsiRyu0, false},
		{shogi.Ryu1, shogi.UsiRyu1, false},
		{0, "", true},
	}

	for i, c := range cases {
		r, e := NewToUsi().Piece(c.in)

		if c.isError && e == nil {
			t.Errorf(`[ToUsi Piece] Expected error, but got nil
Index: %d
Input: %d`, i, c.in)
		}

		if !c.isError && e != nil {
			t.Errorf(`[ToUsi Piece] Expected nil, but got error.
Index: %d
Input: %d
Got:   %v`, i, c.in, e)
		}

		if c.want != r {
			t.Errorf(`[ToUsi Piece] The value was not as expected.
Index:    %d
Input:    %d
Expected: %s
Actual:   %s`, i, c.in, c.want, r)
		}
	}
}
