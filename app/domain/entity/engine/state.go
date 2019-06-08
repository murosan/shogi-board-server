// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package engine

// State represents the state of shogi engine or connection.
type State int

const (
	// NotConnected is the state before connecting to a shogi engine.
	NotConnected State = 1

	// Connected is the state after connected to a shogi engine,
	// and before executing usinewgame(USI command).
	Connected State = 2

	// StandBy is the state after executing usinewgame(USI command),
	// and the shogi engine is not thinking.
	StandBy State = 3

	// Thinking is the state the connected shogi engine is thinking.
	Thinking State = 4
)
