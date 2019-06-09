// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package engine

import (
	"bytes"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"strings"
	"time"

	"github.com/murosan/shogi-board-server/app/domain/entity/option"
	"github.com/murosan/shogi-board-server/app/domain/model/exception"
	"github.com/murosan/shogi-board-server/app/domain/model/infrastructure/os"
	"github.com/murosan/shogi-board-server/app/domain/model/usi"
	parserUSI "github.com/murosan/shogi-board-server/app/lib/parser/usi"
	"github.com/murosan/shogi-board-server/app/logger"
)

var (
	engineNamePrefix   = "id name "
	engineAuthorPrefix = "id author "
	engineOptionPrefix = "option "
)

// Engine represents shogi engine model.
type Engine struct {
	// The engine's key name. The key of the engine's executable path
	// in configuration file will be used. The key must be unique.
	Key string

	// The engine's name given when initializing.
	Name string

	// Author of the engine given when initializing.
	Author string

	// Hols the options given when initializing.
	Options *option.Options

	// Engine state.
	State State

	// Thought result of the engine.
	Result *Result

	// Shogi engine external command. The path written in
	// app config is used. It must be executable.
	// See app/config/config.go.
	Cmd *os.Cmd

	// message receiver
	Ch chan []byte

	logger logger.Logger
}

// New creates new Engine and returns it.
func New(key string, cmd *os.Cmd, logger logger.Logger) (*Engine, error) {
	// first, initialize with nil values
	engine := &Engine{
		Key:    key,
		Name:   "",
		Author: "",
		Options: &option.Options{
			Buttons: make(map[string]*option.Button),
			Checks:  make(map[string]*option.Check),
			Ranges:  make(map[string]*option.Range),
			Selects: make(map[string]*option.Select),
			Texts:   make(map[string]*option.Text),
		},
		State:  NotConnected,
		Result: &Result{}, // TODO
		Cmd:    cmd,
		Ch:     make(chan []byte),
		logger: logger,
	}

	// exec external command
	if err := cmd.Start(); err != nil {
		msg := "failed to start engine. check if the command is executable"
		return nil, errors.Wrap(err, msg)
	}

	engine.State = Connected

	// output receiver
	go engine.catchOutput(engine.Ch)

	// execute initial usi commands
	logger.Info("[engine.New]", zap.ByteString("exec", usi.Usi))
	if err := cmd.Write(usi.Usi); err != nil {
		return nil, errors.Wrap(err, "failed to write 'usi' to the engine")
	}

	logger.Info("[engine.New]", zap.ByteString("wait", usi.UsiOK))
	if err := engine.waitFor(engine.Ch, usi.UsiOK, true); err != nil {
		return nil, errors.Wrap(err, "could not get 'usiok' from the engine")
	}

	logger.Info("[engine.New]", zap.ByteString("exec", usi.IsReady))
	if err := cmd.Write(usi.IsReady); err != nil {
		return nil, errors.Wrap(err, "failed to write 'isready' to the engine")
	}

	logger.Info("[engine.New]", zap.ByteString("wait", usi.ReadyOK))
	if err := engine.waitFor(engine.Ch, usi.ReadyOK, false); err != nil {
		return nil, errors.Wrap(err, "could not get 'readyok' from the engine")
	}

	logger.Info("[engine.New]", zap.String("finished", ""))
	return engine, nil
}

// FlushResult resets result
func (e *Engine) FlushResult() {
	e.Result = &Result{}
}

// Close closes the connection with the shogi engine.
func (e *Engine) Close() error {
	timeout := make(chan struct{})
	ch := make(chan error)
	go func() {
		time.Sleep(time.Second * 10)
		timeout <- struct{}{}
		close(timeout)
	}()
	go func() {
		if err := e.Cmd.Write(usi.Quit); err != nil {
			ch <- errors.Wrap(err, "failed to close engine")
		}
		ch <- e.Cmd.Wait()
	}()

	for {
		select {
		case err := <-ch:
			return err
		case <-timeout:
			err := exception.ErrTimeout
			e.logger.Error("[close] timed out", zap.Error(err))
			return err
		}
	}
}

// Catch engine output while the engine is running.
// When caught output, passes it to given channel
func (e *Engine) catchOutput(ch chan []byte) {
	scanner := e.Cmd.Scanner()
	e.logger.Info("[catch]", zap.String("start catching", ""))

	for scanner.Scan() {
		b := scanner.Bytes()
		e.logger.Info("[catch]", zap.ByteString("engine output", b))
		ch <- b
	}

	// this is called when connection is closed.
	if err := scanner.Err(); err != nil {
		e.logger.Warn("[catch]", zap.Error(err))
	}

	close(ch)
}

func (e *Engine) waitFor(ch chan []byte, exitWord []byte, parseOpt bool) error {
	e.logger.Info("[waitFor]", zap.ByteString("exitWord", exitWord))

	timeout := make(chan struct{})
	go func() {
		time.Sleep(time.Second * 10)
		timeout <- struct{}{}
		close(timeout)
	}()

	for {
		select {
		case b := <-ch:
			s := string(b) // TODO?: is not good for performance

			// set name if s starts with 'id name '
			if parseOpt && strings.HasPrefix(s, engineNamePrefix) {
				n := strings.TrimLeft(s, engineNamePrefix)
				e.Name = strings.TrimSpace(n)
				e.logger.Info("[waitFor]", zap.String("name", e.Name))
			}

			// set author if s starts with 'id author '
			if parseOpt && strings.HasPrefix(s, engineAuthorPrefix) {
				a := strings.TrimLeft(s, engineAuthorPrefix)
				e.Author = strings.TrimSpace(a)
				e.logger.Info("[waitFor]", zap.String("author", e.Author))
			}

			// parse and set option if s starts with 'option '
			if parseOpt && strings.HasPrefix(s, engineOptionPrefix) {
				switch {
				case strings.Contains(s, parserUSI.TypeButton):
					btn, err := parserUSI.ParseButton(s)
					if err != nil {
						e.logger.Error("[waitFor] parse button", zap.Error(err))
					}
					e.logger.Info("[waitFor]", zap.Any("button", btn))
					e.Options.Buttons[btn.Name] = btn

				case strings.Contains(s, parserUSI.TypeCheck):
					chk, err := parserUSI.ParseCheck(s)
					if err != nil {
						e.logger.Error("[waitFor] parse check", zap.Error(err))
					}
					e.logger.Info("[waitFor]", zap.Any("check", chk))
					e.Options.Checks[chk.Name] = chk

				case strings.Contains(s, parserUSI.TypeRange):
					rng, err := parserUSI.ParseRange(s)
					if err != nil {
						e.logger.Error("[waitFor] parse range", zap.Error(err))
					}
					e.logger.Info("[waitFor]", zap.Any("range", rng))
					e.Options.Ranges[rng.Name] = rng

				case strings.Contains(s, parserUSI.TypeSelect):
					sel, err := parserUSI.ParseSelect(s)
					if err != nil {
						e.logger.Error("[waitFor] parse select", zap.Error(err))
					}
					e.logger.Info("[waitFor]", zap.Any("select", sel))
					e.Options.Selects[sel.Name] = sel

				case strings.Contains(s, parserUSI.TypeString):
					txt, err := parserUSI.ParseTextFromStringType(s)
					if err != nil {
						e.logger.Error("[waitFor] parse text", zap.Error(err))
					}
					e.logger.Info("[waitFor]", zap.Any("txt from string", txt))
					e.Options.Texts[txt.Name] = txt

				case strings.Contains(s, parserUSI.TypeFilename):
					txt, err := parserUSI.ParseTextFromFilenameType(s)
					if err != nil {
						e.logger.Error("[waitFor] parse text", zap.Error(err))
					}
					e.logger.Info("[waitFor]", zap.Any("txt from filename", txt))
					e.Options.Texts[txt.Name] = txt
				}

			}

			// exit if received exitWord
			if bytes.Equal(b, exitWord) {
				e.logger.Info("[waitFor]", zap.ByteString("exit", b))
				return nil
			}

		case <-timeout:
			err := exception.ErrTimeout
			e.logger.Error("[waitFor] timed out", zap.Error(err))
			return err
		}
	}
}
