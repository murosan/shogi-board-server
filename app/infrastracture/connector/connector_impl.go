// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package connector

import (
	"bytes"
	"regexp"
	"time"

	"github.com/murosan/shogi-board-server/app/domain/config"
	"github.com/murosan/shogi-board-server/app/domain/entity/converter"
	"github.com/murosan/shogi-board-server/app/domain/entity/engine"
	"github.com/murosan/shogi-board-server/app/domain/entity/usi"
	"github.com/murosan/shogi-board-server/app/domain/exception"
	conn "github.com/murosan/shogi-board-server/app/domain/infrastracture/connector"
	eg "github.com/murosan/shogi-board-server/app/domain/infrastracture/engine"
	"github.com/murosan/shogi-board-server/app/domain/logger"
	"go.uber.org/zap"
)

var (
	idRegex  = regexp.MustCompile(`id.*`)
	optRegex = regexp.MustCompile(`option.*`)
	name     = "name"
	author   = "author"
)

type connector struct {
	conf config.Config
	pool conn.ConnectionPool
	fu   converter.FromUSI
	log  logger.Log
}

// NewConnector 新しい Connector を返す
func NewConnector(
	c config.Config,
	p conn.ConnectionPool,
	fu converter.FromUSI,
	log logger.Log,
) conn.Connector {
	return &connector{c, p, fu, log}
}

// Connect は将棋エンジンに接続します
func (c *connector) Connect() error {
	if c.pool.NamedEngine() != nil {
		c.log.Debug(exception.EngineIsAlreadyRunning.Error() + " Ignore request...")
		return nil
	}

	c.pool.Initialize() // TODO: Nameを渡して2つのエンジンを使えるように
	egn := c.pool.NamedEngine()
	stt := egn.GetState()
	c.log.Debug("Connect", zap.Any("EngineState", stt))

	if e := egn.Start(); e != nil {
		return e
	}

	egn.Lock()
	egn.SetState(engine.Connected)
	go c.catchOutput(egn.GetChan())
	if e := egn.Exec(usi.CmdUsi); e != nil {
		c.log.Error("ExecUsiError", zap.Error(e))
		return e
	}
	if e := c.waitFor(usi.ResOk, true, egn.GetChan()); e != nil {
		c.log.Error("WaitUsiOkError", zap.Error(e))
		return e
	}
	if e := egn.Exec(usi.CmdIsReady); e != nil {
		c.log.Error("ExecIsReadyError", zap.Error(e))
		return e
	}
	if e := c.waitFor(usi.ResReady, false, egn.GetChan()); e != nil {
		c.log.Error("WaitReadyOkError", zap.Error(e))
		return e
	}
	egn.Unlock()
	return nil
}

// Close 将棋エンジンとの接続を切ります
// TODO: エンジンに接続済か確認する処理はどうにか共通化して綺麗にしたい
func (c *connector) Close() error {
	defer c.pool.Remove()
	egn := c.pool.NamedEngine()
	if egn == nil || egn.GetState() == engine.NotConnected {
		c.log.Debug("Close", zap.Any("EngineState", engine.NotConnected))
		return nil
	}

	// TODO: timeout
	return egn.Close()
}

// GetEngine 名前を受け取り Engine を返します
func (c *connector) GetEngine(name string) (eg.Engine, error) {
	e := c.pool.NamedEngine( /* name */ )
	if e == nil || e.GetState() == engine.NotConnected {
		return nil, exception.EngineIsNotRunning
	}

	return e, nil
}

func (c *connector) catchOutput(ch chan []byte) {
	egn := c.pool.NamedEngine()
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

func (c *connector) waitFor(exitWord []byte, parseOpt bool, egnOut chan []byte) error {
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
				c.setID(k, v)
			}

			if parseOpt && optRegex.Match(b) {
				o, e := c.fu.Option(string(b))
				if e != nil {
					return e
				}
				c.appendOption(o)
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

func (c *connector) setID(k, v string) {
	egn := c.pool.NamedEngine()
	switch k {
	case name:
		egn.SetName(v)
	case author:
		egn.SetAuthor(v)
	default:
		panic("unknown id name")
	}
}

// appendOption パース済みのオプションを受け取り、将棋エンジンが保持している一覧に追加
func (c *connector) appendOption(i interface{}) {
	c.pool.NamedEngine().AddOption(i)
}
