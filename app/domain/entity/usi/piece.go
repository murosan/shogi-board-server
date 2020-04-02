// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package usi

// Piece represents usi piece.
type Piece string

const (
	// Fu0 is a Fu in USI expression owned by the first player.
	Fu0 Piece = "P"

	// Kyou0 is a Kyou in USI expression owned by the first player.
	Kyou0 Piece = "L"

	// Kei0 is a Kei in USI expression owned by the first player.
	Kei0 Piece = "N"

	// Gin0 is a Gin in USI expression owned by the first player.
	Gin0 Piece = "S"

	// Kin0 is a Kin in USI expression owned by the first player.
	Kin0 Piece = "G"

	// Kaku0 is a Kaku in USI expression owned by the first player.
	Kaku0 Piece = "B"

	// Hisha0 is a Hisha in USI expression owned by the first player.
	Hisha0 Piece = "R"

	// Gyoku0 is a Gyoku in USI expression owned by the first player.
	Gyoku0 Piece = "K"

	// To0 is a To in USI expression owned by the first player.
	To0 Piece = "+P"

	// NariKyou0 is a NariKyou in USI expression owned by the first player.
	NariKyou0 Piece = "+L"

	// NariKei0 is a NariKei in USI expression owned by the first player.
	NariKei0 Piece = "+N"

	// NariGin0 is a NariGin in USI expression owned by the first player.
	NariGin0 Piece = "+S"

	// Uma0 is a Uma in USI expression owned by the first player.
	Uma0 Piece = "+B"

	// Ryu0 is a Ryu in USI expression owned by the first player.
	Ryu0 Piece = "+R"

	// Fu1 is a Fu in USI expression owned by the second player.
	Fu1 Piece = "p"

	// Kyou1 is a Kyou in USI expression owned by the second player.
	Kyou1 Piece = "l"

	// Kei1 is a Kei in USI expression owned by the second player.
	Kei1 Piece = "n"

	// Gin1 is a Gin in USI expression owned by the second player.
	Gin1 Piece = "s"

	// Kin1 is a Kin in USI expression owned by the second player.
	Kin1 Piece = "g"

	// Kaku1 is a Kaku in USI expression owned by the second player.
	Kaku1 Piece = "b"

	// Hisha1 is a Hisha in USI expression owned by the second player.
	Hisha1 Piece = "r"

	// Gyoku1 is a Gyoku in USI expression owned by the second player.
	Gyoku1 Piece = "k"

	// To1 is a To in USI expression owned by the second player.
	To1 Piece = "+p"

	// NariKyou1 is a NariKyou in USI expression owned by the second player.
	NariKyou1 Piece = "+l"

	// NariKei1 is a NariKei in USI expression owned by the second player.
	NariKei1 Piece = "+n"

	// NariGin1 is a NariGin in USI expression owned by the second player.
	NariGin1 Piece = "+s"

	// Uma1 is a Uma in USI expression owned by the second player.
	Uma1 Piece = "+b"

	// Ryu1 is a Ryu in USI expression owned by the second player.
	Ryu1 Piece = "+r"
)

func (p Piece) String() string { return string(p) }
