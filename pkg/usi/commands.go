// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package usi

var (
	CmdUsi     = []byte("usi")
	CmdIsReady = []byte("isready")
	CmdNewGame = []byte("usinewgame")
	CmdGoInf   = []byte("go infinite")
	CmdStop    = []byte("stop")
	CmdQuit    = []byte("quit")

	ResOk    = []byte("usiok")
	ResReady = []byte("readyok")
)
