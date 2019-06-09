package usi

import (
	"github.com/pkg/errors"

	"github.com/murosan/shogi-board-server/app/domain/model/shogi"
)

// ParsePiece converts USI Piece string to piece id and returns it.
func ParsePiece(s string) (i int, e error) {
	switch s {
	case shogi.UsiFu0:
		i = shogi.Fu0
	case shogi.UsiFu1:
		i = shogi.Fu1
	case shogi.UsiKyou0:
		i = shogi.Kyou0
	case shogi.UsiKyou1:
		i = shogi.Kyou1
	case shogi.UsiKei0:
		i = shogi.Kei0
	case shogi.UsiKei1:
		i = shogi.Kei1
	case shogi.UsiGin0:
		i = shogi.Gin0
	case shogi.UsiGin1:
		i = shogi.Gin1
	case shogi.UsiKin0:
		i = shogi.Kin0
	case shogi.UsiKin1:
		i = shogi.Kin1
	case shogi.UsiKaku0:
		i = shogi.Kaku0
	case shogi.UsiKaku1:
		i = shogi.Kaku1
	case shogi.UsiHisha0:
		i = shogi.Hisha0
	case shogi.UsiHisha1:
		i = shogi.Hisha1
	case shogi.UsiGyoku0:
		i = shogi.Gyoku0
	case shogi.UsiGyoku1:
		i = shogi.Gyoku1
	case shogi.UsiTo0:
		i = shogi.To0
	case shogi.UsiTo1:
		i = shogi.To1
	case shogi.UsiNariKyou0:
		i = shogi.NariKyou0
	case shogi.UsiNariKyou1:
		i = shogi.NariKyou1
	case shogi.UsiNariKei0:
		i = shogi.NariKei0
	case shogi.UsiNariKei1:
		i = shogi.NariKei1
	case shogi.UsiNariGin0:
		i = shogi.NariGin0
	case shogi.UsiNariGin1:
		i = shogi.NariGin1
	case shogi.UsiUma0:
		i = shogi.Uma0
	case shogi.UsiUma1:
		i = shogi.Uma1
	case shogi.UsiRyu0:
		i = shogi.Ryu0
	case shogi.UsiRyu1:
		i = shogi.Ryu1
	default:
		e = errors.New("PieceIDが不正です id = " + string(s))
	}
	return
}
