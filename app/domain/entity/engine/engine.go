// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package engine

import (
	"fmt"
	"sync"
)

// ID represents an id of shogi engine.
type ID string

func (id ID) String() string { return string(id) }

// Engine represents shogi engine.
type Engine struct {
	sync.RWMutex

	// The key of the engine's executable path in configuration file.
	// This ID must be unique.
	id ID

	// The engine's name. Lazily initialized on start.
	name string

	// Author of the engine. Lazily initialized on start.
	author string

	// Holder of engine options. Lazily initialized on start.
	options *Options

	// Engine state.
	state State

	// Shogi engine external command path. The path written in
	// app config is used. It must be executable.
	// See app/config/config.go.
	path string
}

// New creates new Engine and returns it.
func New(id ID, path string) *Engine {
	return &Engine{
		id:      id,
		name:    "",
		author:  "",
		options: NewOptions(),
		state:   NotConnected,
		path:    path,
	}
}

func (e *Engine) String() string {
	return fmt.Sprintf(
		"Engine{id:%s,name:%s,author:%s,options:%s,state:%s,path:%s}",
		e.GetID(), e.GetName(), e.GetAuthor(), e.GetOptions(), e.GetState(), e.GetPath(),
	)
}

func (e *Engine) GetID() ID { return e.id }

func (e *Engine) GetName() string {
	e.RLock()
	defer e.RUnlock()
	return e.name
}

func (e *Engine) SetName(name string) {
	e.Lock()
	e.name = name
	e.Unlock()
}

func (e *Engine) GetAuthor() string {
	e.RLock()
	defer e.RUnlock()
	return e.author
}

func (e *Engine) SetAuthor(author string) {
	e.Lock()
	e.author = author
	e.Unlock()
}

func (e *Engine) GetState() State {
	e.RLock()
	defer e.RUnlock()
	return e.state
}

func (e *Engine) SetState(state State) {
	if !state.isValid() {
		panic(fmt.Sprintf("unknown state '%s'", state))
	}
	e.Lock()
	e.state = state
	e.Unlock()
}

func (e *Engine) GetOptions() *Options {
	e.RLock()
	defer e.RUnlock()
	return e.options
}

func (e *Engine) GetPath() string { return e.path }
