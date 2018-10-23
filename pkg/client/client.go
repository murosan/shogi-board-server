// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package client

import (
	"bufio"
	"bytes"
	"log"
	"regexp"
	"time"

	"github.com/murosan/shogi-proxy-server/pkg/config"
	"github.com/murosan/shogi-proxy-server/pkg/converter/models"
	"github.com/murosan/shogi-proxy-server/pkg/converter/usi2go"
	"github.com/murosan/shogi-proxy-server/pkg/engine"
	"github.com/murosan/shogi-proxy-server/pkg/msg"
	"github.com/murosan/shogi-proxy-server/pkg/usi"
)

var (
	idRegex  = regexp.MustCompile(`id.*`)
	optRegex = regexp.MustCompile(`option.*`)
)

// TODO: 2つのエンジンを同時に使えるようにする。Poolとか作る
type Client struct {
	conf    *config.Config
	egn     *engine.Engine
	options map[string]models.Option
}

func NewClient(c *config.Config) *Client {
	return &Client{c, nil, make(map[string]models.Option)}
}

func (c *Client) Connect() error {
	if c.egn != nil {
		log.Println(msg.EngineIsAlreadyRunning.Error() + " Ignore request...")
		return nil
	}

	c.egn = engine.NewEngine(c.conf.EnginePath)
	c.SetState(engine.Connected)
	c.egn.Cmd.Start()
	go c.CatchOutput()

	c.egn.Mux.Lock()
	if e := c.Exec(usi.CmdUsi); e != nil {
		return e
	}
	if e := c.waitFor(usi.ResOk, true); e != nil {
		return e
	}
	if e := c.Exec(usi.CmdIsReady); e != nil {
		return e
	}
	if e := c.waitFor(usi.ResReady, false); e != nil {
		return e
	}
	c.egn.Mux.Unlock()
	return nil
}

func (c *Client) Close() error {
	if c.egn == nil {
		return nil
	}
	// TODO: エラーをちゃんと返せない
	c.egn.Mux.Lock()
	c.Exec(usi.CmdQuit)
	timeout := make(chan struct{})
	go func() {
		time.Sleep(time.Second * 10)
		timeout <- struct{}{}
	}()
	c.egn.Cmd.Wait() // TODO: timeout と wait が適当すぎて、ちゃんと動かないだろう
	defer c.egn.Mux.Unlock()
	for {
		select {
		case <-c.egn.Done:
			return nil
		case <-timeout:
			return msg.ConnectionTimeout
		}
	}
	c.egn = nil
	return nil
}

func (c *Client) Exec(b []byte) error {
	if c.GetState() == engine.NotConnected {
		return msg.EngineIsNotRunning
	}
	return c.egn.Exec(b)
}

func (c *Client) CatchOutput() {
	defer func() {}()
	s := bufio.NewScanner(c.egn.Stdout)

	for s.Scan() {
		b := s.Bytes()
		log.Println("[EngineOut] " + string(b))
		c.egn.EngineOut <- b
	}

	if s.Err() != nil {
		log.Println("scan:", s.Err())
	}
	close(c.egn.Done)
}

func (c *Client) SetState(s int) {
	c.egn.State = s
}

func (c *Client) GetState() int {
	return c.egn.State
}

func (c *Client) SetId(k []byte, v []byte) error {
	switch string(k) {
	case "name":
		c.egn.Name = v
		return nil
	case "author":
		c.egn.Author = v
		return nil
	default:
		return msg.InvalidIdSyntax.WithMsg("Key was not 'name' or 'author'.")
	}
}

func (c *Client) SetupOption(o models.Option) {
	c.options[string(o.GetName())] = o
}

func (c *Client) OptionList() []models.Option {
	// TODO: どの形で渡すのがいいかなぁ
	return nil
}

func (c *Client) waitFor(exitWord []byte, parseOpt bool) error {
	timeout := make(chan struct{})
	go func() {
		time.Sleep(time.Second * 10)
		timeout <- struct{}{}
	}()
	for {
		select {
		case b := <-c.egn.EngineOut:
			if len(b) == 0 {
				continue
			}

			if bytes.Equal(b, exitWord) {
				return nil
			}

			if parseOpt && idRegex.Match(b) {
				k, v, e := usi2go.ParseId(b)
				if e != nil {
					c.SetId(k, v)
					continue
				}
				return e
			}

			if parseOpt && optRegex.Match(b) {
				o, e := usi2go.ParseOpt(b)
				if e == nil {
					c.SetupOption(o)
					continue
				}
				return e
			}
		case <-timeout:
			log.Println(msg.ConnectionTimeout.Error())
			return msg.ConnectionTimeout
		}
	}
}
