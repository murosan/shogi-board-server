// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package to_usi

import (
	"strconv"

	"github.com/murosan/shogi-proxy-server/app/domain/entity/exception"
	"github.com/murosan/shogi-proxy-server/app/domain/entity/shogi"
)

// TODO: USIの文字列もentityに移す。byteで返すかどうか
func (tu *ToUsi) Piece(i int) (s string, e error) {
	switch i {
	case shogi.Fu0:
		s = "P"
	case shogi.Fu1:
		s = "p"
	case shogi.Kyou0:
		s = "L"
	case shogi.Kyou1:
		s = "l"
	case shogi.Kei0:
		s = "N"
	case shogi.Kei1:
		s = "n"
	case shogi.Gin0:
		s = "S"
	case shogi.Gin1:
		s = "s"
	case shogi.Kin0:
		s = "G"
	case shogi.Kin1:
		s = "g"
	case shogi.Kaku0:
		s = "B"
	case shogi.Kaku1:
		s = "b"
	case shogi.Hisha0:
		s = "R"
	case shogi.Hisha1:
		s = "r"
	case shogi.Gyoku0:
		s = "K"
	case shogi.Gyoku1:
		s = "k"
	case shogi.To0:
		s = "+P"
	case shogi.To1:
		s = "+p"
	case shogi.NariKyou0:
		s = "+L"
	case shogi.NariKyou1:
		s = "+l"
	case shogi.NariKei0:
		s = "+N"
	case shogi.NariKei1:
		s = "+n"
	case shogi.NariGin0:
		s = "+S"
	case shogi.NariGin1:
		s = "+s"
	case shogi.Uma0:
		s = "+B"
	case shogi.Uma1:
		s = "+b"
	case shogi.Ryu0:
		s = "+R"
	case shogi.Ryu1:
		s = "+r"
	default:
		e = exception.InvalidPieceId.WithMsg("PieceIdが不正です id=" + strconv.Itoa(i))
	}
	return
}
