// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package command

import (
	"io"
	"log"
	"os/exec"

	"github.com/murosan/shogi-proxy-server/app/domain/infrastracture/os/command"
)

type osCmd struct {
	// 将棋エンジンの実行コマンド
	cmd *exec.Cmd

	// 将棋エンジンへ入力を渡すパイプ
	in io.WriteCloser

	// 将棋エンジンの出力を受け取るパイプ
	out io.ReadCloser
}

func NewCmd(cmd *exec.Cmd) command.OsCmd {
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatalln("connect stdin: " + err.Error())
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalln("connect stdout: " + err.Error())
	}
	return &osCmd{cmd, stdin, stdout}
}

func (c *osCmd) Start() error { return c.cmd.Start() }

func (c *osCmd) Wait() error { return c.cmd.Wait() }

func (c *osCmd) Write(b []byte) error {
	_, err := c.in.Write(b)
	return err
}

func (c *osCmd) GetStdoutPipe() *io.ReadCloser { return &c.out }