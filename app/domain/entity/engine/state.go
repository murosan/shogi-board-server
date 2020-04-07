// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package engine

// State represents the state of shogi engine.
type State int

const (
	_ State = iota // ignore zero

	// NotConnected is the state before connecting to a shogi engine.
	NotConnected

	// Connected is the state after connected to a shogi engine,
	// and before executing 'usi' (USI command).
	// It also means the engine has never thought yet.
	Connected

	// StandBy is the state the engine is not thinking.
	// It means the engine has thought over once, and is stopped thinking now.
	StandBy

	// Thinking is the state the connected shogi engine is thinking.
	Thinking
)

func (s State) String() string {
	switch s {
	case NotConnected:
		return "State(NotConnected)"
	case Connected:
		return "State(Connected)"
	case StandBy:
		return "State(StandBy)"
	case Thinking:
		return "State(Thinking)"
	default:
		return "State(Unknown)"
	}
}

func (s State) isValid() bool {
	return NotConnected <= s && s <= Thinking
}
