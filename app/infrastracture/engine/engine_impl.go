// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package engine

import (
	"bufio"
	"sync"

	eg "github.com/murosan/shogi-board-server/app/domain/entity/engine"
	"github.com/murosan/shogi-board-server/app/domain/entity/usi"
	"github.com/murosan/shogi-board-server/app/domain/exception"
	engineModel "github.com/murosan/shogi-board-server/app/domain/infrastracture/engine"
	"github.com/murosan/shogi-board-server/app/domain/infrastracture/os/command"
	"github.com/murosan/shogi-board-server/app/domain/logger"
	pb "github.com/murosan/shogi-board-server/app/proto"

	"go.uber.org/zap"
)

type engine struct {
	cmd command.OsCmd

	// 将棋エンジンの状態を管理
	state eg.State

	// その他の情報
	name    string
	author  string
	options *pb.Options

	// エンジンの思考結果を貯めておくところ
	result *pb.Result

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
		state:   eg.NotConnected,
		options: pb.NewOptions(),
		result:  pb.NewResult(),
		sc:      bufio.NewScanner(*c.GetStdoutPipe()),
		ch:      make(chan []byte),
		log:     log,
	}
}

func (e *engine) GetName() string { return e.name }

func (e *engine) SetName(n string) { e.name = n }

func (e *engine) GetAuthor() string { return e.author }

func (e *engine) SetAuthor(a string) { e.author = a }

func (e *engine) AddOption(i interface{}) { eg.AppendOption(e.options, i) }

func (e *engine) GetOptions() *pb.Options { return e.options }

func (e *engine) SetState(s eg.State) { e.state = s }

func (e *engine) GetState() eg.State { return e.state }

func (e *engine) SetResult(i *pb.Info, key int) { e.result.Result[int32(key)] = i }

func (e *engine) GetResult() *pb.Result { return e.result }

func (e *engine) FlushResult() { e.result = pb.NewResult() }

func (e *engine) Lock() { e.mux.Lock() }

func (e *engine) Unlock() { e.mux.Unlock() }

func (e *engine) Exec(b []byte) error {
	e.log.Info("StdinPipe", zap.ByteString("Exec", b))
	return e.cmd.Write(append(b, '\n'))
}

func (e *engine) Start() error { return e.cmd.Start() }

func (e *engine) Close() error {
	if err := e.Exec(usi.CmdQuit); err != nil {
		e.log.Error("ExecError", zap.Error(err))
		return exception.FailedToExecUSI.WithMsg(err.Error())
	}
	return e.cmd.Wait()
}

func (e *engine) GetScanner() *bufio.Scanner { return e.sc }

func (e *engine) GetChan() chan []byte { return e.ch }
