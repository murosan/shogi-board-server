// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package context

import (
	"github.com/murosan/shogi-board-server/app/config"
	"github.com/murosan/shogi-board-server/app/logger"
	"github.com/murosan/shogi-board-server/app/server/context"
)

var c *context.Context

// Init initializes Context instance.
// This must be called before using context and called only once.
func Init(logger logger.Logger, config *config.Config) {
	if c != nil {
		panic("Context is already initialized")
	}

	c = context.New(logger, config)
}

// Use returns a Context instance.
// Call Init once before.
func Use() *context.Context {
	return c
}
