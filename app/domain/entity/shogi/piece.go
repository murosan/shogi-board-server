// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package shogi

import (
	"fmt"
)

// Piece represents shogi piece.
type Piece int

const (
	// Empty is not piece, is just a empty cell.
	Empty Piece = 0

	// Fu0 is a Fu owned by the first player.
	// Fu moves like a Pawn in chess.
	Fu0 Piece = 1

	// Kyou0 is a Kyousha owned by the first player.
	// Kyousha moves like a Rook in chess, but can only go straight forward.
	Kyou0 Piece = 2

	// Kei0 is a Keima owned by the first player.
	// Keima moves like a Knight in chess, but can only go forward.
	Kei0 Piece = 3

	// Gin0 is a Gin owned by the first player.
	// Gin moves like King in chess, but can not move behind and to the side.
	Gin0 Piece = 4

	// Kin0 is a Kin owned by the first player.
	// Kin moves like King in chess, but can not move diagonally backwards.
	Kin0 Piece = 5

	// Kaku0 is a Kaku owned by the first player.
	// Kaku moves as same as Bishop in chess.
	Kaku0 Piece = 6

	// Hisha0 is a Hisha owned by the first player.
	// Hisha moves as same as Rook in chess.
	Hisha0 Piece = 7

	// Gyoku0 is a Gyoku owned by the first player.
	// Gyoku moves like King in chess.
	Gyoku0 Piece = 8

	// To0 is a To owned by the first player.
	// To moves as same as Kin.
	To0 Piece = 11

	// NariKyou0 is a NariKyou owned by the first player.
	// NariKyou moves as same as Kin.
	NariKyou0 Piece = 12

	// NariKei0 is a NariKei owned by the first player.
	// NariKei moves as same as Kin.
	NariKei0 Piece = 13

	// NariGin0 is a NariGin owned by the first player.
	// NariGin moves as same as Kin.
	NariGin0 Piece = 14

	// Uma0 is a Uma owned by the first player.
	// Uma can move to the place where Gyoku and Kaku can move.
	Uma0 Piece = 16

	// Ryu0 is a Ryu owned by the first player.
	// Ryu can moves to the place where Gyoku and Hisha can move.
	Ryu0 Piece = 17

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

func (p Piece) ToInt() int { return int(p) }

func (p Piece) String() string {
	return fmt.Sprintf("shogi.Piece(%d)", int(p))
}
