// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package shogi

import (
	"fmt"
	"github.com/pkg/errors"
)

const (
	// Empty is not piece, is just a empty cell.
	Empty = 0

	// Fu0 is a Fu owned by the first player.
	// Fu moves like a Pawn in chess.
	Fu0 = 1

	// Kyou0 is a Kyousha owned by the first player.
	// Kyousha moves like a Rook in chess, but can only go straight forward.
	Kyou0 = 2

	// Kei0 is a Keima owned by the first player.
	// Keima moves like a Knight in chess, but can only go forward.
	Kei0 = 3

	// Gin0 is a Gin owned by the first player.
	// Gin moves like King in chess, but can not move behind and to the side.
	Gin0 = 4

	// Kin0 is a Kin owned by the first player.
	// Kin moves like King in chess, but can not move diagonally backwards.
	Kin0 = 5

	// Kaku0 is a Kaku owned by the first player.
	// Kaku moves as same as Bishop in chess.
	Kaku0 = 6

	// Hisha0 is a Hisha owned by the first player.
	// Hisha moves as same as Rook in chess.
	Hisha0 = 7

	// Gyoku0 is a Gyoku owned by the first player.
	// Gyoku moves like King in chess.
	Gyoku0 = 8

	// To0 is a To owned by the first player.
	// To moves as same as Kin.
	To0 = 11

	// NariKyou0 is a NariKyou owned by the first player.
	// NariKyou moves as same as Kin.
	NariKyou0 = 12

	// NariKei0 is a NariKei owned by the first player.
	// NariKei moves as same as Kin.
	NariKei0 = 13

	// NariGin0 is a NariGin owned by the first player.
	// NariGin moves as same as Kin.
	NariGin0 = 14

	// Uma0 is a Uma owned by the first player.
	// Uma can move to the place where Gyoku and Kaku can move.
	Uma0 = 16

	// Ryu0 is a Ryu owned by the first player.
	// Ryu can moves to the place where Gyoku and Hisha can move.
	Ryu0 = 17

	// Fu1 is a Fu owned by the second player.
	Fu1 = -Fu0

	// Kyou1 is a Kyou owned by the second player.
	Kyou1 = -Kyou0

	// Kei1 is a Kei owned by the second player.
	Kei1 = -Kei0

	// Gin1 is a Gin owned by the second player.
	Gin1 = -Gin0

	// Kin1 is a Kin owned by the second player.
	Kin1 = -Kin0

	// Kaku1 is a Kaku owned by the second player.
	Kaku1 = -Kaku0

	// Hisha1 is a Hisha owned by the second player.
	Hisha1 = -Hisha0

	// Gyoku1 is a Gyoku owned by the second player.
	Gyoku1 = -Gyoku0

	// To1 is a To owned by the second player.
	To1 = -To0

	// NariKyou1 is a NariKyou owned by the second player.
	NariKyou1 = -NariKyou0

	// NariKei1 is a NariKei owned by the second player.
	NariKei1 = -NariKei0

	// NariGin1 is a NariGin owned by the second player.
	NariGin1 = -NariGin0

	// Uma1 is a Uma owned by the second player.
	Uma1 = -Uma0

	// Ryu1 is a Ryu owned by the second player.
	Ryu1 = -Ryu0
)

// PieceToUSI converts piece id to USIPiece string and returns it.
func PieceToUSI(i int) (s string, e error) {
	switch i {
	case Fu0:
		s = UsiFu0
	case Fu1:
		s = UsiFu1
	case Kyou0:
		s = UsiKyou0
	case Kyou1:
		s = UsiKyou1
	case Kei0:
		s = UsiKei0
	case Kei1:
		s = UsiKei1
	case Gin0:
		s = UsiGin0
	case Gin1:
		s = UsiGin1
	case Kin0:
		s = UsiKin0
	case Kin1:
		s = UsiKin1
	case Kaku0:
		s = UsiKaku0
	case Kaku1:
		s = UsiKaku1
	case Hisha0:
		s = UsiHisha0
	case Hisha1:
		s = UsiHisha1
	case Gyoku0:
		s = UsiGyoku0
	case Gyoku1:
		s = UsiGyoku1
	case To0:
		s = UsiTo0
	case To1:
		s = UsiTo1
	case NariKyou0:
		s = UsiNariKyou0
	case NariKyou1:
		s = UsiNariKyou1
	case NariKei0:
		s = UsiNariKei0
	case NariKei1:
		s = UsiNariKei1
	case NariGin0:
		s = UsiNariGin0
	case NariGin1:
		s = UsiNariGin1
	case Uma0:
		s = UsiUma0
	case Uma1:
		s = UsiUma1
	case Ryu0:
		s = UsiRyu0
	case Ryu1:
		s = UsiRyu1
	default:
		e = errors.New("PieceIDが不正です id=" + fmt.Sprint(i))
	}
	return
}
