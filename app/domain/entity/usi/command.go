// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package usi

var (
	// CmdUsi USI の usi コマンド
	CmdUsi = []byte("usi")
	// CmdIsReady USI の isready コマンド
	CmdIsReady = []byte("isready")
	// CmdNewGame USI の usinewgame コマンド
	CmdNewGame = []byte("usinewgame")
	// CmdGoInf USI の go infinite コマンド
	CmdGoInf = []byte("go infinite")
	// CmdStop USI の stop コマンド
	CmdStop = []byte("stop")
	// CmdQuit USI の quit コマンド
	CmdQuit = []byte("quit")

	// ResOk USI の usiok コマンド
	ResOk = []byte("usiok")
	// ResReady USI の readyok コマンド
	ResReady = []byte("readyok")
)
