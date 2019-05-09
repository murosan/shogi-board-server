// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package shogi

// Turn は手番を表す
type Turn int

const (
	// Sente は先手のこと
	Sente Turn = 1

	// Gote は後手のこと
	Gote Turn = -Sente
)
