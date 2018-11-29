// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package to_usi

import (
	"strconv"

	"github.com/murosan/shogi-proxy-server/app/domain/entity/shogi"
	"github.com/murosan/shogi-proxy-server/app/domain/exception"
)

// TODO: USIの文字列もentityに移す。byteで返すかどうか
func (tu *ToUsi) Piece(i int) (s string, e error) {
	switch i {
	case shogi.Fu0:
		s = shogi.UsiFu0
	case shogi.Fu1:
		s = shogi.UsiFu1
	case shogi.Kyou0:
		s = shogi.UsiKyou0
	case shogi.Kyou1:
		s = shogi.UsiKyou1
	case shogi.Kei0:
		s = shogi.UsiKei0
	case shogi.Kei1:
		s = shogi.UsiKei1
	case shogi.Gin0:
		s = shogi.UsiGin0
	case shogi.Gin1:
		s = shogi.UsiGin1
	case shogi.Kin0:
		s = shogi.UsiKin0
	case shogi.Kin1:
		s = shogi.UsiKin1
	case shogi.Kaku0:
		s = shogi.UsiKaku0
	case shogi.Kaku1:
		s = shogi.UsiKaku1
	case shogi.Hisha0:
		s = shogi.UsiHisha0
	case shogi.Hisha1:
		s = shogi.UsiHisha1
	case shogi.Gyoku0:
		s = shogi.UsiGyoku0
	case shogi.Gyoku1:
		s = shogi.UsiGyoku1
	case shogi.To0:
		s = shogi.UsiTo0
	case shogi.To1:
		s = shogi.UsiTo1
	case shogi.NariKyou0:
		s = shogi.UsiNariKyou0
	case shogi.NariKyou1:
		s = shogi.UsiNariKyou1
	case shogi.NariKei0:
		s = shogi.UsiNariKei0
	case shogi.NariKei1:
		s = shogi.UsiNariKei1
	case shogi.NariGin0:
		s = shogi.UsiNariGin0
	case shogi.NariGin1:
		s = shogi.UsiNariGin1
	case shogi.Uma0:
		s = shogi.UsiUma0
	case shogi.Uma1:
		s = shogi.UsiUma1
	case shogi.Ryu0:
		s = shogi.UsiRyu0
	case shogi.Ryu1:
		s = shogi.UsiRyu1
	default:
		e = exception.InvalidPieceId.WithMsg("PieceIdが不正です id=" + strconv.Itoa(i))
	}
	return
}
