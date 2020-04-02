// Copyright 2020 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package usi

var (
	// Command is a set of USI commands.
	Command = struct {
		USI,
		IsReady,
		NewGame,
		GoInf,
		Stop,
		Quit []byte
	}{
		USI:     []byte("usi"),
		IsReady: []byte("isready"),
		NewGame: []byte("usinewgame"),
		GoInf:   []byte("go infinite"),
		Stop:    []byte("stop"),
		Quit:    []byte("quit"),
	}

	Response = struct {
		OK,
		ReadyOK []byte
	}{
		OK:      []byte("usiok"),
		ReadyOK: []byte("readyok"),
	}
)
