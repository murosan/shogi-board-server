// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package engine

import "sync"

// State represents the state of shogi engine or connection.
type State struct {
	sync.Mutex
	current StateID
}

// NewState returns new state
func NewState() *State {
	return &State{current: NotConnected}
}

// Set changes current state id
func (s *State) Set(id StateID) {
	s.Lock()
	s.current = id
	s.Unlock()
}

// Get returns current state id
func (s *State) Get() StateID {
	s.Lock()
	defer s.Unlock()
	return s.current
}

// StateID is a alias of int type, and the values of this type
// represents the state of shogi engine.
type StateID int

const (
	// NotConnected is the state before connecting to a shogi engine.
	NotConnected StateID = 1

	// Connected is the state after connected to a shogi engine,
	// and before executing usinewgame(USI command).
	Connected StateID = 2

	// StandBy is the state after executing usinewgame(USI command),
	// and the shogi engine is not thinking.
	StandBy StateID = 3

	// Thinking is the state the connected shogi engine is thinking.
	Thinking StateID = 4
)
