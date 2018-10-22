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

	"github.com/murosan/shogi-proxy-server/pkg/msg"
)

type Engine struct {
	// 将棋エンジンの実行コマンド
	Cmd *exec.Cmd

	// 将棋エンジンの状態 state.go を参照
	State struct{}

	// 将棋エンジンへ入力を渡すパイプ
	Stdin io.WriteCloser
	// 将棋エンジンの出力を受け取るパイプ
	Stdout io.ReadCloser

	// 将棋エンジンが出力した値を読み取る Scanner
	Sc *bufio.Scanner
	// 将棋エンジンの出力を渡すチャネル
	EngineOut chan []byte

	// その他の情報
	Name    []byte
	Author  []byte

	// 将棋エンジンが終了したかどうか
	Done chan struct{}
	Mux  sync.Mutex
}

// p: EngineCommandPath
func NewEngine(p string) *Engine {
	log.Println("NewEngine to engine.")
	cmd := exec.Command(p)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatalln("connect stdin: " + err.Error())
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalln("connect stdout: " + err.Error())
	}

	return &Engine{
		Cmd:   cmd,
		State: NotConnected,

		Stdin:  stdin,
		Stdout: stdout,

		Sc:        bufio.NewScanner(stdout),
		EngineOut: make(chan []byte, 10),

		Done: make(chan struct{}),
	}
}

// USIコマンドの実行
func (e *Engine) Exec(b []byte) error {
	_, err := e.Stdin.Write(append(b, '\n'))
	if err != nil {
		log.Println(msg.FailedToExecCommand)
		return err
	}
	return nil
}
