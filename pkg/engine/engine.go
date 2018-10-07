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
)

// TODO: 名前をどうにかしたい

// Singleton
var Engine *Client

type Client struct {
	// 将棋エンジンの実行コマンド
	Cmd *exec.Cmd

	// 将棋エンジンの入力パイプ
	Stdin io.WriteCloser
	// 将棋エンジンの出力パイプ
	Stdout io.ReadCloser

	// 将棋エンジンが出力した値を読み取る Scanner
	Sc *bufio.Scanner

	// 将棋エンジンが終了したかどうか
	mux  sync.Mutex
	Done chan struct{}
}

// 将棋エンジンと接続する
// TODO: エンジンの入れ替えをできるようにする
func Connect() {
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

		Sc:   bufio.NewScanner(stdout),
		Done: make(chan struct{}),
	}
}

// 将棋エンジンを終了する
func Close() {
	// TODO: エンジンにquitを送って正常終了できるように
	Engine.Stdin.Close()
	Engine.Stdout.Close()
}
