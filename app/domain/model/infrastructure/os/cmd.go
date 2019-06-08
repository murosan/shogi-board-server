// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

import (
	"io"
	"os/exec"
	"time"

	"github.com/murosan/shogi-board-server/app/config"
	"github.com/murosan/shogi-board-server/app/domain/model/exception"
)

// Cmd is a wrapper of exec.Cmd.
type Cmd struct {
	// Executable command of the shogi engine.
	cmd *exec.Cmd

	// Pipe to pass USI commands to the shogi engine through stdin.
	in io.WriteCloser

	// Pipe to receive responses from the shogi engine through stdout.
	out io.ReadCloser
}

// Start executes the command.
func (c *Cmd) Start() error { return c.cmd.Start() }

// Wait waits the command finishes.
// Returns timeout error when exceeds the time configured in app/config/app.go.
func (c *Cmd) Wait() error {
	ch := make(chan error, 1)
	go func() { ch <- c.cmd.Wait() }()

	select {
	case err := <-ch:
		return err
	case <-time.After(config.Timeout):
		return exception.ErrTimeout
	}
}

// Write passes the received bytes to the command.
func (c *Cmd) Write(b []byte) error {
	_, err := c.in.Write(b)
	return err
}

// Stdout returns stdout pipe.
func (c *Cmd) Stdout() *io.ReadCloser { return &c.out }
