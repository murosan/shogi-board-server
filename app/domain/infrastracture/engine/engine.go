// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package engine

import (
	"bufio"

	"github.com/murosan/shogi-proxy-server/app/domain/entity/engine/state"
)

// TODO
type Engine interface {
	GetName() *[]byte
	SetName(*[]byte)

	GetAuthor() *[]byte
	SetAuthor(*[]byte)

	SetState(state.State)
	GetState() state.State

	Lock()
	Unlock()

	Exec(*[]byte) error

	Start() error
	Close(chan struct{})

	// 将棋エンジンが出力した値を読み取る Scanner を作る
	GetScanner() *bufio.Scanner
}
