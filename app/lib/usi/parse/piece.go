package parse

import (
	"fmt"

	"github.com/murosan/shogi-board-server/app/domain/entity/shogi"
	"github.com/murosan/shogi-board-server/app/domain/entity/usi"
)

func Piece(p usi.Piece) (s shogi.Piece, e error) {
	switch p {
	case usi.Fu0:
		s = shogi.Fu0
	case usi.Fu1:
		s = shogi.Fu1
	case usi.Kyou0:
		s = shogi.Kyou0
	case usi.Kyou1:
		s = shogi.Kyou1
	case usi.Kei0:
		s = shogi.Kei0
	case usi.Kei1:
		s = shogi.Kei1
	case usi.Gin0:
		s = shogi.Gin0
	case usi.Gin1:
		s = shogi.Gin1
	case usi.Kin0:
		s = shogi.Kin0
	case usi.Kin1:
		s = shogi.Kin1
	case usi.Kaku0:
		s = shogi.Kaku0
	case usi.Kaku1:
		s = shogi.Kaku1
	case usi.Hisha0:
		s = shogi.Hisha0
	case usi.Hisha1:
		s = shogi.Hisha1
	case usi.Gyoku0:
		s = shogi.Gyoku0
	case usi.Gyoku1:
		s = shogi.Gyoku1
	case usi.To0:
		s = shogi.To0
	case usi.To1:
		s = shogi.To1
	case usi.NariKyou0:
		s = shogi.NariKyou0
	case usi.NariKyou1:
		s = shogi.NariKyou1
	case usi.NariKei0:
		s = shogi.NariKei0
	case usi.NariKei1:
		s = shogi.NariKei1
	case usi.NariGin0:
		s = shogi.NariGin0
	case usi.NariGin1:
		s = shogi.NariGin1
	case usi.Uma0:
		s = shogi.Uma0
	case usi.Uma1:
		s = shogi.Uma1
	case usi.Ryu0:
		s = shogi.Ryu0
	case usi.Ryu1:
		s = shogi.Ryu1
	default:
		e = fmt.Errorf("PieceIDが不正です id = %d", int(s))
	}
	return
}
