// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package context

import (
	"sync"

	"github.com/murosan/shogi-board-server/app/config"
	"github.com/murosan/shogi-board-server/app/domain/entity/engine"
	"github.com/murosan/shogi-board-server/app/logger"
)

// Context represents shogi-board-server.
// It holds entities, modules.
type Context struct {
	sync.RWMutex
	Logger  logger.Logger
	Config  *config.Config
	Engines map[string]*engine.Engine
}

// New returns new Context.
func New(logger logger.Logger, config *config.Config) *Context {
	return &Context{
		Engines: make(map[string]*engine.Engine),
		Config:  config,
		Logger:  logger,
	}
}

// SetEngine set the engine to Engines
func (c *Context) SetEngine(name string, e *engine.Engine) {
	c.Lock()
	c.Engines[name] = e
	c.Unlock()
}

// GetEngine find a engine from the given key and returns it.
func (c *Context) GetEngine(name string) (*engine.Engine, bool) {
	c.RLock()
	e, ok := c.Engines[name]
	c.RUnlock()
	return e, ok
}
