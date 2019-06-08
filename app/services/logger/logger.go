// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package logger

import (
	"github.com/murosan/shogi-board-server/app/config"
	"github.com/murosan/shogi-board-server/app/logger"
)

// Hold as a singleton instance.
var l logger.Logger

// Init initializes Logger instance.
// This must be called before using logger and called only once.
func Init(c *config.Config) {
	if l != nil {
		panic("Logger is already initialized")
	}

	l = logger.New(c)
}

// Use returns a Logger instance.
// Call Init once before.
func Use() logger.Logger {
	return l
}
