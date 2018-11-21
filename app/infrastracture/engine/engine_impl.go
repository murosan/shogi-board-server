// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package engine

import (
	"bufio"
	"sync"

	"github.com/murosan/shogi-proxy-server/app/domain/entity/engine/option"
	"github.com/murosan/shogi-proxy-server/app/domain/entity/engine/state"
	"github.com/murosan/shogi-proxy-server/app/domain/entity/usi"
	engineModel "github.com/murosan/shogi-proxy-server/app/domain/infrastracture/engine"
	"github.com/murosan/shogi-proxy-server/app/domain/infrastracture/os/command"
	"github.com/murosan/shogi-proxy-server/app/service/logger"
	"go.uber.org/zap"
)

type engine struct {
	cmd command.OsCmd

	// 将棋エンジンの状態を管理
	state state.State

	// その他の情報
	name    string
	author  string
	options option.OptMap

	// エンジンの出力を流し込む scanner
	// Singleton で持っておく
	sc *bufio.Scanner

	mux sync.Mutex
}

func NewEngine(c command.OsCmd) engineModel.Engine {
	return &engine{
		cmd:     c,
		state:   state.NotConnected,
		options: *option.EmptyOptMap(),
		sc:      bufio.NewScanner(*c.GetStdoutPipe()),
	}
}

func (e *engine) GetName() string { return e.name }

func (e *engine) SetName(n string) { e.name = n }

func (e *engine) GetAuthor() string { return e.author }

func (e *engine) SetAuthor(a string) { e.author = a }

func (e *engine) SetOption(n string, opt option.Option) { e.options.Push(opt) }

func (e *engine) GetOptions() option.OptMap { return e.options }

func (e *engine) SetState(s state.State) { e.state = s }

func (e *engine) GetState() state.State { return e.state }

func (e *engine) Lock() { e.mux.Lock() }

func (e *engine) Unlock() { e.mux.Unlock() }

// USIコマンドの実行
func (e *engine) Exec(b []byte) error {
	logger.Use().Info("StdinPipe", zap.ByteString("Exec", b))
	return e.cmd.Write(append(b, '\n'))
}

func (e *engine) Start() error { return e.cmd.Start() }

func (e *engine) Close() error {
	e.Exec(usi.CmdQuit)
	return e.cmd.Wait()
}

func (e *engine) GetScanner() *bufio.Scanner { return e.sc }
