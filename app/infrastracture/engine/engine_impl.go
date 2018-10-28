// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package engine

import (
	"bufio"
	"sync"

	"github.com/murosan/shogi-proxy-server/app/domain/entity/engine/state"
	egn "github.com/murosan/shogi-proxy-server/app/domain/infrastracture/engine"
	"github.com/murosan/shogi-proxy-server/app/domain/infrastracture/os/command"
)

type engine struct {
	cmd command.OsCmd

	// 将棋エンジンの状態を管理
	state state.State

	// その他の情報
	name   []byte
	author []byte

	mux sync.Mutex
}

func NewEngine(c command.OsCmd) egn.Engine {
	return &engine{
		cmd:   c,
		state: state.NotConnected,
	}
}

func (e *engine) GetName() *[]byte { return &e.name }

func (e *engine) SetName(b *[]byte) { e.name = *b }

func (e *engine) GetAuthor() *[]byte { return &e.author }

func (e *engine) SetAuthor(b *[]byte) { e.author = *b }

func (e *engine) SetState(s state.State) { e.state = s }

func (e *engine) GetState() state.State { return e.state }

func (e *engine) Lock() { e.mux.Lock() }

func (e *engine) Unlock() { e.mux.Unlock() }

// USIコマンドの実行
func (e *engine) Exec(b *[]byte) error { return e.cmd.Write(append(*b, '\n')) }

func (e *engine) Start() error { return e.cmd.Start() }

func (e *engine) Close(c chan struct{}) {
	// TODO: quit をこっちから書き込むか検討
	// TODO: timeout とかのエラーを返す
	e.cmd.Wait()
	c <- struct{}{}
}

func (e *engine) GetScanner() *bufio.Scanner { return bufio.NewScanner(*e.cmd.GetStdoutPipe()) }
