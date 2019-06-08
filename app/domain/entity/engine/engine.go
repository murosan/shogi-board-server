// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package engine

import (
	"github.com/murosan/shogi-board-server/app/domain/model/infrastructure/os"
)

// Engine represents shogi engine model.
type Engine struct {
	// The engine's key name. The key of the engine's executable path
	// in configuration file will be used. The key must be unique.
	Key string

	// The engine's name given when initializing.
	Name string

	// Author of the engine given when initializing.
	Author string

	// Hols the options given when initializing.
	Options *Options

	// Engine state.
	State State

	// Thought result of the engine.
	Result *Result

	// Shogi engine external command. The path written in
	// app config is used. It must be executable.
	// See app/config/config.go.
	Cmd *os.Cmd
}

// New creates new Engine and returns it.
// All the constructors are initialized by nil values.
func New(key string) *Engine {
	return &Engine{
		Key:    key,
		Name:   "",
		Author: "",
		Options: &Options{
			Buttons: make(map[string]*Button),
			Checks:  make(map[string]*Check),
			Ranges:  make(map[string]*Range),
			Selects: make(map[string]*Select),
			Texts:   make(map[string]*Text),
		},
		State:  NotConnected,
		Result: &Result{},
		Cmd:    nil,
	}
}
