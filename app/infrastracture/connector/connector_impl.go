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
)

var (
	idRegex  = regexp.MustCompile(`id.*`)
	optRegex = regexp.MustCompile(`option.*`)
)

type connector struct {
	conf config.Config
	pool conn.ConnectionPool
	fu   *from_usi.FromUsi

	// TODO: 2つのエンジンから受け取る方法を考える
	egnOut chan []byte
}

func NewConnector(c config.Config, p conn.ConnectionPool, fu *from_usi.FromUsi) conn.Connector {
	return &connector{c, p, fu, make(chan []byte)}
}

func (c *connector) Connect() error {
	egn := c.pool.NamedEngine()

	if egn.GetState() != state.NotConnected {
		log.Println(exception.EngineIsAlreadyRunning.Error() + " Ignore request...")
		return nil
	}

	egn.Start()
	egn.SetState(state.Connected)
	go c.catchOutput(c.egnOut)

	egn.Lock()
	if e := c.Exec(&usi.CmdUsi); e != nil {
		return e
	}
	if e := c.waitFor(usi.ResOk, true, c.egnOut); e != nil {
		return e
	}
	if e := c.Exec(&usi.CmdIsReady); e != nil {
		return e
	}
	if e := c.waitFor(usi.ResReady, false, c.egnOut); e != nil {
		return e
	}
	egn.Unlock()
	return nil
}

func (c *connector) Close() error {
	egn := c.pool.NamedEngine()
	if egn.GetState() == state.NotConnected {
		return nil
	}

	// TODO: timeout
	if err := egn.Close(); err != nil {
		return err
	}

	return nil
}

func (c *connector) Exec(b *[]byte) error {
	egn := c.pool.NamedEngine()
	if egn.GetState() == state.NotConnected {
		return exception.EngineIsNotRunning
	}
	return egn.Exec(b)
}

func (c *connector) catchOutput(ch chan []byte) {
	egn := c.pool.NamedEngine()
	s := egn.GetScanner()

	for s.Scan() {
		b := s.Bytes()
		log.Println("[EngineOut] " + string(b))
		ch <- b
	}

	if s.Err() != nil {
		log.Println("scan:", s.Err())
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
			if len(b) == 0 {
				continue
			}

			if bytes.Equal(b, exitWord) {
				return nil
			}

			if parseOpt && idRegex.Match(b) {
				k, v, e := c.fu.EngineID(b)
				if e != nil {
					c.setId(&k, &v)
					continue
				}
				return e
			}

			if parseOpt && optRegex.Match(b) {
				o, e := c.fu.Option(b)
				if e == nil {
					c.setOption(o)
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

func (c *connector) setId(k, v *[]byte) {
	egn := c.pool.NamedEngine()
	switch string(*k) {
	case "name":
		egn.SetName(v)
	case "author":
		egn.SetAuthor(v)
	default:
		panic("unknown id name")
	}
}

func (c *connector) setOption(o option.Option) {
	c.pool.NamedEngine().SetOption(string(o.GetName()), o)
}
