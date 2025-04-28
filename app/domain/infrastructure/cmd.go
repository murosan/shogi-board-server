// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package infrastructure

import (
	"bufio"
	"errors"
	"io"
	"os/exec"
	"path/filepath"
	"time"

	"golang.org/x/xerrors"
)

// Cmd is simple interface of exec.Cmd.
type Cmd interface {
	io.Writer
	Start() error
	Wait(timeout time.Duration) error
	Scanner() *bufio.Scanner
	Chdir(dir string)
}

type cmd struct {
	cmd     *exec.Cmd
	in      io.WriteCloser
	out     io.ReadCloser
	scanner *bufio.Scanner
}

// NewCmd returns new Cmd.
func NewCmd(path string) Cmd {
	return &cmd{
		cmd:     exec.Command(path),
		in:      nil, // lazily initialized on Start
		out:     nil,
		scanner: nil,
	}
}

func (c *cmd) Start() error {
	stdin, err := c.cmd.StdinPipe()
	if err != nil {
		return xerrors.Errorf("get stdin pipe: %w", err)
	}
	c.in = stdin

	stdout, err := c.cmd.StdoutPipe()
	if err != nil {
		return xerrors.Errorf("get stdout pipe: %w", err)
	}
	c.out = stdout
	c.scanner = bufio.NewScanner(stdout)
	c.scanner.Split(bufio.ScanLines) // just to make sure

	return c.cmd.Start()
}

func (c *cmd) Wait(timeout time.Duration) error {
	wait := make(chan error, 1)
	go func() { wait <- c.cmd.Wait() }()

	select {
	case err := <-wait:
		return err
	case <-time.After(timeout):
		return errors.New("timeout on closing cmd")
	}
}

func (c *cmd) Write(b []byte) (int, error) { return c.in.Write(b) }

func (c *cmd) Scanner() *bufio.Scanner { return c.scanner }

func (c *cmd) Chdir(dir string) { c.cmd.Dir = filepath.Clean(dir) }
