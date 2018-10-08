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

	"github.com/murosan/shogi-proxy-server/pkg/config"
	"github.com/murosan/shogi-proxy-server/pkg/usi"
)

// TODO: 名前をどうにかしたい

var (
	// Singleton
	Engine *Client = nil
)

type Client struct {
	// 将棋エンジンの実行コマンド
	Cmd *exec.Cmd

	// 将棋エンジンへ入力を渡すパイプ
	Stdin io.WriteCloser
	// 将棋エンジンの出力を受け取るパイプ
	Stdout io.ReadCloser

	// 将棋エンジンが出力した値を読み取る Scanner
	Sc *bufio.Scanner
	// 将棋エンジンの出力を渡すチャネル
	EngineOut chan []byte

	// 将棋エンジンが終了したかどうか
	Mux  sync.Mutex
	Done chan struct{}
}

// 将棋エンジンと接続する
// TODO: エンジンの入れ替えをできるようにする
func Connect() {
	log.Println("Connect to engine.")
	cmd := exec.Command(config.Conf.EnginePath)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatalln("connect stdin: " + err.Error())
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalln("connect stdout: " + err.Error())
	}

	Engine = &Client{
		Cmd:    cmd,
		Stdin:  stdin,
		Stdout: stdout,

		Sc:        bufio.NewScanner(stdout),
		EngineOut: make(chan []byte, 10),
		Done:      make(chan struct{}),
	}
}

// 将棋エンジンを終了する
func Close() {
	Engine.Mux.Lock()
	log.Println("Close engine.")
	// TODO: エラーとかね
	if Engine == nil {
		return
	}
	Engine.Stdin.Write(append(usi.CmdQuit, '\n'))
	Engine.Stdin.Close()
	Engine.Stdout.Close()
	Engine.Cmd.Wait()
	<-Engine.Done
	Engine.Mux.Unlock()
	Engine = nil
}
