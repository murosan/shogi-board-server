// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

import (
	"encoding/json"

	confModel "github.com/murosan/shogi-board-server/app/domain/config"
	"go.uber.org/zap"
)

type config struct {
	EnginePath map[string]string `json:"enginePath"`
	Log        zap.Config        `json:"Log"`
}

// NewConfig return Config
func NewConfig(b []byte) confModel.Config {
	var c config
	if err := json.Unmarshal(b, &c); err != nil {
		panic(err)
	}
	return &c
}

func (c *config) GetEnginePath(key string) string {
	return c.EnginePath[key]
}

func (c *config) GetEngineNames() []string {
	l := make([]string, len(c.EnginePath))
	i := 0
	for name := range c.EnginePath {
		l[i] = name
		i++
	}
	return l
}

func (c *config) GetLogConf() zap.Config {
	return c.Log
}
