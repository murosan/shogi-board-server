// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package shogi

// Turn represents turn of the game.
type Turn int

const (
	// Sente is the first player in shogi.
	// It is said `black` in chess.
	Sente Turn = 1

	// Gote is the second player in shogi.
	// It is said `white` in chess.
	Gote Turn = -Sente
)
