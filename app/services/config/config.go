// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

import (
	"github.com/murosan/shogi-board-server/app/config"
)

// Holds as a singleton instance.
var c *config.Config

// Init initializes Config instance.
// This must be called before using config and called only once.
func Init(appPath, logPath string) {
	if c != nil {
		panic("Config is already initialized")
	}

	c = config.New(appPath, logPath)
}

// Use returns a Config instance.
// Call Init once before.
func Use() *config.Config {
	return c
}
