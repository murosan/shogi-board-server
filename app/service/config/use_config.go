// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

import (
	"io/ioutil"

	"github.com/murosan/shogi-board-server/app/config"
	confModel "github.com/murosan/shogi-board-server/app/domain/config"
)

var c confModel.Config

// InitConfig Config の初期化
// Config はシングルトンで持っておく
func InitConfig(path string) {
	if c == nil {
		b, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}
		c = config.NewConfig(b)
	}
}

// UseConfig シングルトンで保持している Config を返す
func UseConfig() confModel.Config {
	if c == nil {
		panic("Must run InitConfig(), first.")
	}
	return c
}
