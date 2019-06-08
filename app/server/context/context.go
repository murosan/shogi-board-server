// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package context

import (
	"github.com/murosan/shogi-board-server/app/config"
	"github.com/murosan/shogi-board-server/app/domain/entity/engine"
	"github.com/murosan/shogi-board-server/app/logger"
)

// Context represents shogi-board-server.
// It holds entities, modules.
type Context struct {
	Logger  logger.Logger
	Config  *config.Config
	Engines map[string]engine.Engine
}

// New returns new Context.
func New(logger logger.Logger, config *config.Config) *Context {
	return &Context{
		Engines: make(map[string]engine.Engine),
		Config:  config,
		Logger:  logger,
	}
}

// ActiveEngines returns list of the shogi engine connected.
func (c *Context) ActiveEngines() (e []engine.Engine) {
	for _, v := range c.Engines {
		e = append(e, v)
	}
	return
}
