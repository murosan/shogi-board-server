// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package engine

import (
	"bufio"

	"github.com/murosan/shogi-proxy-server/app/domain/entity/engine/option"
	"github.com/murosan/shogi-proxy-server/app/domain/entity/engine/result"
	"github.com/murosan/shogi-proxy-server/app/domain/entity/engine/state"
)

// TODO
type Engine interface {
	GetName() string
	SetName(string)

	GetAuthor() string
	SetAuthor(string)

	SetOption(string, option.Option)
	GetOptions() option.OptMap
	UpdateOption(option.UpdateOptionValue) error

	SetState(state.State)
	GetState() state.State

	SetResult(i *result.Info, key int)
	FlushResult()

	Lock()
	Unlock()

	Exec([]byte) error

	Start() error
	Close() error

	// 将棋エンジンが出力した値を読み取る Scanner を作る
	GetScanner() *bufio.Scanner
}
