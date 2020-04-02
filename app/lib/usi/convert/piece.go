package convert

import (
	"errors"
	"fmt"

	"github.com/murosan/shogi-board-server/app/domain/entity/shogi"
	"github.com/murosan/shogi-board-server/app/domain/entity/usi"
)

// Piece converts shogi.Piece to usi.Piece.
func Piece(p shogi.Piece) (s usi.Piece, e error) {
	switch p {
	case shogi.Fu0:
		s = usi.Fu0
	case shogi.Fu1:
		s = usi.Fu1
	case shogi.Kyou0:
		s = usi.Kyou0
	case shogi.Kyou1:
		s = usi.Kyou1
	case shogi.Kei0:
		s = usi.Kei0
	case shogi.Kei1:
		s = usi.Kei1
	case shogi.Gin0:
		s = usi.Gin0
	case shogi.Gin1:
		s = usi.Gin1
	case shogi.Kin0:
		s = usi.Kin0
	case shogi.Kin1:
		s = usi.Kin1
	case shogi.Kaku0:
		s = usi.Kaku0
	case shogi.Kaku1:
		s = usi.Kaku1
	case shogi.Hisha0:
		s = usi.Hisha0
	case shogi.Hisha1:
		s = usi.Hisha1
	case shogi.Gyoku0:
		s = usi.Gyoku0
	case shogi.Gyoku1:
		s = usi.Gyoku1
	case shogi.To0:
		s = usi.To0
	case shogi.To1:
		s = usi.To1
	case shogi.NariKyou0:
		s = usi.NariKyou0
	case shogi.NariKyou1:
		s = usi.NariKyou1
	case shogi.NariKei0:
		s = usi.NariKei0
	case shogi.NariKei1:
		s = usi.NariKei1
	case shogi.NariGin0:
		s = usi.NariGin0
	case shogi.NariGin1:
		s = usi.NariGin1
	case shogi.Uma0:
		s = usi.Uma0
	case shogi.Uma1:
		s = usi.Uma1
	case shogi.Ryu0:
		s = usi.Ryu0
	case shogi.Ryu1:
		s = usi.Ryu1
	default:
		e = errors.New("PieceIDが不正です id=" + fmt.Sprint(p))
	}
	return
}
