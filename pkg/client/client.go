// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package client

import (
	"bufio"
	"github.com/murosan/shogi-proxy-server/pkg/config"
	"log"

	"github.com/murosan/shogi-proxy-server/pkg/engine"
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
	c.egn = engine.NewEngine(c.conf.EnginePath)
	// TODO
	return nil
}

func (c *Client) Close() error {
	c.egn.Mux.Lock()
	log.Println("Close engine.")
	// TODO: エラーとかね
	c.egn.Stdin.Write(append(usi.CmdQuit, '\n'))
	c.egn.Cmd.Wait()
	<-c.egn.Done
	c.egn.Mux.Unlock()
	return nil
}

// byte配列を受け取ってエンジンに渡す
func (c *Client) Exec(b []byte) error {
	// TODO
	c.egn.Exec(b)
	return nil
}

// エンジンの出力を受け取り続ける
func (c *Client) CatchEngineOutput() {
	defer func() {}()
	s := bufio.NewScanner(c.egn.Stdout)

	for s.Scan() {
		b := s.Bytes()
		c.egn.EngineOut <- b
	}

	if s.Err() != nil {
		log.Println("scan: ", s.Err())
	}
	close(c.egn.Done)
}
