// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

import (
	cnf "github.com/murosan/shogi-proxy-server/app/config"
	"github.com/murosan/shogi-proxy-server/app/domain/entity/config"
)

var c config.Config = nil

func UseConfig() config.Config {
	if c == nil {
		c = cnf.NewConfig()
	}
	return c
}
