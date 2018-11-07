// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

import (
	"io/ioutil"

	"github.com/murosan/shogi-proxy-server/app/config"
	confModel "github.com/murosan/shogi-proxy-server/app/domain/entity/config"
)

var c confModel.Config = nil

func UseConfig() confModel.Config {
	if c == nil {
		b, err := ioutil.ReadFile("./config.json")
		if err != nil {
			panic(err)
		}
		c = config.NewConfig(b)
	}
	return c
}
