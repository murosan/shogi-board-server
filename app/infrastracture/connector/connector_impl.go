// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package connector

import (
	"bytes"
	"log"
	"regexp"
	"time"

	"github.com/murosan/shogi-proxy-server/app/domain/entity/config"
	"github.com/murosan/shogi-proxy-server/app/domain/entity/converter/from_usi"
	"github.com/murosan/shogi-proxy-server/app/domain/entity/engine/option"
	"github.com/murosan/shogi-proxy-server/app/domain/entity/engine/state"
	"github.com/murosan/shogi-proxy-server/app/domain/entity/exception"
	"github.com/murosan/shogi-proxy-server/app/domain/entity/usi"
	conn "github.com/murosan/shogi-proxy-server/app/domain/infrastracture/connector"
	"github.com/murosan/shogi-proxy-server/app/domain/infrastracture/engine"
	engineService "github.com/murosan/shogi-proxy-server/app/service/engine"
)

var (
	idRegex  = regexp.MustCompile(`id.*`)
	optRegex = regexp.MustCompile(`option.*`)
)

// TODO: 2つのエンジンを同時に使えるようにする。Poolとか作る
type connector struct {
	conf    config.Config
	egn     engine.Engine
	egnOut  chan []byte
	options map[string]option.Option
	fu      *from_usi.FromUsi
}

func NewConnector(c config.Config, fu *from_usi.FromUsi) conn.Connector {
	return &connector{
		c,
		nil,
		make(chan []byte, 10),
		make(map[string]option.Option),
		fu,
	}
}

func (c *connector) Connect() error {
	if c.egn != nil {
		log.Println(exception.EngineIsAlreadyRunning.Error() + " Ignore request...")
		return nil
	}

	// TODO: 引数で渡す。Poolを作ってその中に入れるように修正する
	c.egn = engineService.UseEngine()
	c.SetState(state.Connected)
	c.egn.Start()
	go c.CatchOutput()

	c.egn.Lock()
	if e := c.Exec(&usi.CmdUsi); e != nil {
		return e
	}
	if e := c.waitFor(usi.ResOk, true); e != nil {
		return e
	}
	if e := c.Exec(&usi.CmdIsReady); e != nil {
		return e
	}
	if e := c.waitFor(usi.ResReady, false); e != nil {
		return e
	}
	c.egn.Unlock()
	return nil
}

func (c *connector) Close() error {
	if c.egn == nil {
		return nil
	}

	c.Exec(&usi.CmdQuit)

	timeout := make(chan struct{})
	closeCh := make(chan struct{})

	go func() {
		time.Sleep(time.Second * 10)
		timeout <- struct{}{}
	}()

	c.egn.Close(closeCh)

	for {
		select {
		case <-closeCh:
			c.egn = nil
			return nil
		case <-timeout:
			return exception.ConnectionTimeout
		}
	}
}

func (c *connector) Exec(b *[]byte) error {
	if c.GetState() == state.NotConnected {
		return exception.EngineIsNotRunning
	}
	return c.egn.Exec(b)
}

func (c *connector) CatchOutput() {
	s := c.egn.GetScanner()

	for s.Scan() {
		b := s.Bytes()
		log.Println("[EngineOut] " + string(b))
		c.egnOut <- b
	}

	if s.Err() != nil {
		log.Println("scan:", s.Err())
	}
}

func (c *connector) SetState(s state.State) {
	c.egn.SetState(s)
}

func (c *connector) GetState() state.State {
	return c.egn.GetState()
}

func (c *connector) SetId(k *[]byte, v *[]byte) error {
	switch string(*k) {
	case "name":
		c.egn.SetName(v)
		return nil
	case "author":
		c.egn.SetAuthor(v)
		return nil
	default:
		return exception.InvalidIdSyntax.WithMsg("Key was not 'name' or 'author'.")
	}
}

func (c *connector) SetupOption(o option.Option) {
	c.options[string(o.GetName())] = o
}

func (c *connector) OptionList() []option.Option {
	// TODO: どの形で渡すのがいいかなぁ
	return nil
}

func (c *connector) waitFor(exitWord []byte, parseOpt bool) error {
	timeout := make(chan struct{})
	go func() {
		time.Sleep(time.Second * 10)
		timeout <- struct{}{}
	}()
	for {
		select {
		case b := <-c.egnOut:
			if len(b) == 0 {
				continue
			}

			if bytes.Equal(b, exitWord) {
				return nil
			}

			if parseOpt && idRegex.Match(b) {
				k, v, e := c.fu.EngineID(b)
				if e != nil {
					c.SetId(&k, &v)
					continue
				}
				return e
			}

			if parseOpt && optRegex.Match(b) {
				o, e := c.fu.Option(b)
				if e == nil {
					c.SetupOption(o)
					continue
				}
				return e
			}
		case <-timeout:
			log.Println(exception.ConnectionTimeout.Error())
			return exception.ConnectionTimeout
		}
	}
}
