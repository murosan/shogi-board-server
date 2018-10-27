// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package engine

import (
	"bufio"
	"io"
	"log"
	"os/exec"
	"sync"

	"github.com/murosan/shogi-proxy-server/app/domain/entity/engine/state"
	"github.com/murosan/shogi-proxy-server/app/domain/entity/exception"
	egn "github.com/murosan/shogi-proxy-server/app/domain/infrastracture/engine"
)

type engine struct {
	// 将棋エンジンの実行コマンド
	cmd *exec.Cmd

	// 将棋エンジンの状態 state.go を参照
	state state.State

	// 将棋エンジンへ入力を渡すパイプ
	stdin io.WriteCloser
	// 将棋エンジンの出力を受け取るパイプ
	stdout io.ReadCloser

	// 将棋エンジンが出力した値を読み取る Scanner
	sc *bufio.Scanner

	// その他の情報
	name   []byte
	author []byte

	// 将棋エンジンが終了したかどうか
	done chan struct{}
	mux  sync.Mutex
}

// p: EngineCommandPath
func NewEngine(p string) egn.Engine {
	log.Println("Initializing Engine...")
	cmd := exec.Command(p)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatalln("connect stdin: " + err.Error())
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalln("connect stdout: " + err.Error())
	}

	return &engine{
		cmd:   cmd,
		state: state.NotConnected,

		stdin:  stdin,
		stdout: stdout,

		sc: bufio.NewScanner(stdout),

		done: make(chan struct{}),
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
func (e *engine) Exec(b *[]byte) error {
	_, err := e.stdin.Write(append(*b, '\n'))
	if err != nil {
		log.Println(exception.FailedToExecCommand)
		return err
	}
	return nil
}

func (e *engine) Start() error {
	return e.cmd.Start()
}

func (e *engine) Close(c chan struct{}) {
	e.cmd.Wait()
	c <- struct{}{}
}

func (e *engine) GetScanner() *bufio.Scanner {
	return bufio.NewScanner(e.stdout)
}
