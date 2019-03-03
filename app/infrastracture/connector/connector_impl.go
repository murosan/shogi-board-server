// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package connector

import (
	"bytes"
	"os/exec"
	"regexp"
	"time"

	"github.com/murosan/shogi-board-server/app/domain/config"
	"github.com/murosan/shogi-board-server/app/domain/entity/converter"
	"github.com/murosan/shogi-board-server/app/domain/entity/engine"
	"github.com/murosan/shogi-board-server/app/domain/entity/usi"
	"github.com/murosan/shogi-board-server/app/domain/exception"
	conn "github.com/murosan/shogi-board-server/app/domain/infrastracture/connector"
	mengine "github.com/murosan/shogi-board-server/app/domain/infrastracture/engine"
	"github.com/murosan/shogi-board-server/app/domain/logger"
	iengine "github.com/murosan/shogi-board-server/app/infrastracture/engine"
	"github.com/murosan/shogi-board-server/app/infrastracture/os/command"
	"github.com/murosan/shogi-board-server/app/lib/stringutil"

	"go.uber.org/zap"
)

var (
	idRegex  = regexp.MustCompile(`id.*`)
	optRegex = regexp.MustCompile(`option.*`)
)

type connector struct {
	conf config.Config
	em   map[string]mengine.Engine
	fu   converter.FromUSI
	log  logger.Log
}

// NewConnector 新しい Connector を返す
func NewConnector(
	c config.Config,
	fu converter.FromUSI,
	log logger.Log,
) conn.Connector {
	em := make(map[string]mengine.Engine)
	return &connector{c, em, fu, log}
}

// Connect は将棋エンジンに接続します
func (c *connector) Connect(name string) error {
	if _, ok := c.em[name]; ok {
		c.log.Debug(exception.EngineIsAlreadyRunning.Error() + " Ignore request...")
		return nil
	}

	// 将棋エンジンの名前一覧を Conf から取得
	names := c.conf.GetEngineNames()

	// Conf にエンジンの名前があるかチェックする
	if !stringutil.Contains(names, name) {
		return exception.UnknownEngineName
	}

	// バイナリ実行
	cmd := exec.Command(c.conf.GetEnginePath(name))
	egn := iengine.NewEngine(command.NewCmd(cmd), c.log)
	c.em[name] = egn

	stt := egn.GetState()
	c.log.Debug("Connect", zap.Any("EngineState", stt))

	if e := egn.Start(); e != nil {
		return e
	}

	egn.Lock()
	egn.SetState(engine.Connected)
	go c.catchOutput(egn.GetChan(), name)
	if e := egn.Exec(usi.CmdUsi); e != nil {
		c.log.Error("ExecUsiError", zap.Error(e))
		return e
	}
	if e := c.waitFor(usi.ResOk, true, egn.GetChan(), name); e != nil {
		c.log.Error("WaitUsiOkError", zap.Error(e))
		return e
	}
	if e := egn.Exec(usi.CmdIsReady); e != nil {
		c.log.Error("ExecIsReadyError", zap.Error(e))
		return e
	}
	if e := c.waitFor(usi.ResReady, false, egn.GetChan(), name); e != nil {
		c.log.Error("WaitReadyOkError", zap.Error(e))
		return e
	}
	egn.Unlock()
	return nil
}

// Close 将棋エンジンとの接続を切ります
func (c *connector) Close(name string) error {
	// config にエンジンの名前があるかチェックする
	if !stringutil.Contains(c.conf.GetEngineNames(), name) {
		return exception.UnknownEngineName
	}

	egn, ok := c.em[name]

	if !ok {
		c.log.Debug("Close", zap.Any("EngineState", engine.NotConnected))
		return nil
	}

	defer delete(c.em, name)
	return egn.Close()
}

// GetEngine 名前を受け取り Engine を返します
func (c *connector) GetEngine(name string) (mengine.Engine, error) {
	egn, ok := c.em[name]
	if !ok || egn.GetState() == engine.NotConnected {
		return nil, exception.EngineIsNotRunning
	}

	return egn, nil
}

// GetEngineNames は app.yml で設定された接続可能な将棋エンジンの名前一覧を返します
func (c *connector) GetEngineNames() []string {
	return c.conf.GetEngineNames()
}

func (c *connector) catchOutput(ch chan []byte, name string) {
	egn := c.em[name]
	s := egn.GetScanner()

	for s.Scan() {
		b := s.Bytes()
		c.log.Info("StdoutPipe", zap.ByteString("EngineOut", b))
		ch <- b
	}

	if s.Err() != nil {
		c.log.Debug("scan:" + s.Err().Error())
	}
}

func (c *connector) waitFor(
	exitWord []byte,
	parseOpt bool,
	egnOut chan []byte,
	name string,
) error {
	timeout := make(chan struct{})
	go func() {
		time.Sleep(time.Second * 10)
		timeout <- struct{}{}
		close(timeout)
	}()
	for {
		select {
		case b := <-egnOut:
			if parseOpt && idRegex.Match(b) {
				k, v, e := c.fu.EngineID(string(b))
				if e != nil {
					return e
				}
				c.setID(k, v, name)
			}

			if parseOpt && optRegex.Match(b) {
				o, e := c.fu.Option(string(b))
				if e != nil {
					return e
				}
				c.appendOption(o, name)
			}

			if bytes.Equal(b, exitWord) {
				return nil
			}
		case <-timeout:
			c.log.Error(exception.ConnectionTimeout.Error())
			return exception.ConnectionTimeout
		}
	}
}

func (c *connector) setID(k, v, name string) {
	egn := c.em[name]
	switch k {
	case "name":
		egn.SetName(v)
	case "author":
		egn.SetAuthor(v)
	default:
		panic("unknown id name")
	}
}

// appendOption パース済みのオプションを受け取り、将棋エンジンが保持している一覧に追加
func (c *connector) appendOption(i interface{}, name string) {
	egn := c.em[name]
	egn.AddOption(i)
}
