// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package engine

import (
	"bufio"
	"sync"

	"github.com/murosan/shogi-board-server/app/domain/entity/engine/option"
	"github.com/murosan/shogi-board-server/app/domain/entity/engine/result"
	"github.com/murosan/shogi-board-server/app/domain/entity/engine/state"
	"github.com/murosan/shogi-board-server/app/domain/entity/usi"
	"github.com/murosan/shogi-board-server/app/domain/exception"
	engineModel "github.com/murosan/shogi-board-server/app/domain/infrastracture/engine"
	"github.com/murosan/shogi-board-server/app/domain/infrastracture/os/command"
	"github.com/murosan/shogi-board-server/app/domain/logger"
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

	// エンジンの思考結果を貯めておくところ
	result *result.Result

	// エンジンの出力を流し込む scanner
	// Singleton で持っておく
	sc *bufio.Scanner

	ch chan []byte

	mux sync.Mutex

	log logger.Log
}

// NewEngine 新しい Engine を返す
func NewEngine(c command.OsCmd, log logger.Log) engineModel.Engine {
	return &engine{
		cmd:     c,
		state:   state.NotConnected,
		options: *option.NewOptMap(),
		result:  result.NewResult(),
		sc:      bufio.NewScanner(*c.GetStdoutPipe()),
		ch:      make(chan []byte),
		log:     log,
	}
}

func (e *engine) GetName() string { return e.name }

func (e *engine) SetName(n string) { e.name = n }

func (e *engine) GetAuthor() string { return e.author }

func (e *engine) SetAuthor(a string) { e.author = a }

func (e *engine) SetOption(n string, opt option.Option) { e.options.Append(opt) }

func (e *engine) GetOptions() option.OptMap { return e.options }

func (e *engine) UpdateOption(v option.UpdateOptionValue) error {
	u, err := e.options.Update(v)
	if err != nil {
		e.log.Warn("EngineUpdateOption", zap.Error(exception.FailedToUpdateOption))
		return err
	}
	return e.Exec([]byte(u))
}

func (e *engine) SetState(s state.State) { e.state = s }

func (e *engine) GetState() state.State { return e.state }

func (e *engine) SetResult(i result.Info, key int) { e.result.Set(i, key) }

func (e *engine) GetResult() *result.Result { return e.result }

func (e *engine) FlushResult() { e.result.Flush() }

func (e *engine) Lock() { e.mux.Lock() }

func (e *engine) Unlock() { e.mux.Unlock() }

func (e *engine) Exec(b []byte) error {
	e.log.Info("StdinPipe", zap.ByteString("Exec", b))
	return e.cmd.Write(append(b, '\n'))
}

func (e *engine) Start() error { return e.cmd.Start() }

func (e *engine) Close() error {
	e.Exec(usi.CmdQuit)
	return e.cmd.Wait()
}

func (e *engine) GetScanner() *bufio.Scanner { return e.sc }

func (e *engine) GetChan() chan []byte { return e.ch }
