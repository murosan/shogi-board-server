// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package client

import (
	"bufio"
	"bytes"
	"log"
	"time"

	"github.com/murosan/shogi-proxy-server/pkg/config"
	"github.com/murosan/shogi-proxy-server/pkg/engine"
	"github.com/murosan/shogi-proxy-server/pkg/msg"
	"github.com/murosan/shogi-proxy-server/pkg/usi"
)

// TODO: 2つのエンジンを同時に使えるようにする。Poolとか作る
type Client struct {
	conf *config.Config
	egn  *engine.Engine
}

func NewClient(c *config.Config) *Client {
	return &Client{c, nil}
}

func (c *Client) Connect() error {
	if c.egn != nil {
		log.Println(msg.EngineIsAlreadyRunning.Error() + " Ignore request...")
		return nil
	}

	c.egn = engine.NewEngine(c.conf.EnginePath)
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
	c.SetState(engine.Connected)
	c.egn.Mux.Unlock()
	return nil
}

func (c *Client) Close() error {
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
	return nil
}

// byte配列を受け取ってエンジンに渡す
func (c *Client) Exec(b []byte) error {
	// TODO
	c.egn.Exec(b)
	return nil
}

// エンジンの出力を受け取り続ける
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

func (c *Client) SetState(s struct{}) {
	c.egn.State = s
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

			if parseOpt {
				// id でパースしてみて、失敗したら option でパース
				if e := c.egn.ParseId(b); e == nil {
					continue
				}

				if e := c.egn.ParseOpt(b); e != nil {
					log.Println(e)
					return e
				}
			}
		case <-timeout:
			log.Println("Connection timeout.")
			return msg.ConnectionTimeout
		}
	}
}
