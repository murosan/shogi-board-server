// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package engine

import (
	"bufio"
	"go.uber.org/zap"
	"sync"

	"github.com/murosan/shogi-proxy-server/app/domain/entity/engine/option"
	"github.com/murosan/shogi-proxy-server/app/domain/entity/engine/state"
	"github.com/murosan/shogi-proxy-server/app/domain/entity/usi"
	engineModel "github.com/murosan/shogi-proxy-server/app/domain/infrastracture/engine"
	"github.com/murosan/shogi-proxy-server/app/domain/infrastracture/os/command"
	"github.com/murosan/shogi-proxy-server/app/service/logger"
)

type engine struct {
	cmd command.OsCmd

	// 将棋エンジンの状態を管理
	state state.State

	// その他の情報
	name    []byte
	author  []byte
	options map[string]option.Option

	// エンジンの出力を流し込む scanner
	// Singleton で持っておく
	sc *bufio.Scanner

	mux sync.Mutex
}

func NewEngine(c command.OsCmd) engineModel.Engine {
	return &engine{
		cmd:     c,
		state:   state.NotConnected,
		options: make(map[string]option.Option),
		sc:      bufio.NewScanner(*c.GetStdoutPipe()),
	}
}

func (e *engine) GetName() *[]byte { return &e.name }

func (e *engine) SetName(b *[]byte) { e.name = *b }

func (e *engine) GetAuthor() *[]byte { return &e.author }

func (e *engine) SetAuthor(b *[]byte) { e.author = *b }

func (e *engine) SetOption(n string, o option.Option) { e.options[n] = o }

func (e *engine) GetOptions() map[string]option.Option { return e.options }

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
