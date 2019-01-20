// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

import (
	confModel "github.com/murosan/shogi-board-server/app/domain/config"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

type config struct {
	app appConfig
	log zap.Config
}

type appConfig struct {
	Engines map[string]string `json:"engines" yaml:"engines"`
}

// NewConfig return Config
func NewConfig(appBytes, logBytes []byte) confModel.Config {
	var app appConfig
	var log zap.Config

	if err := yaml.Unmarshal(appBytes, &app); err != nil {
		panic(err)
	}
	if err := yaml.Unmarshal(logBytes, &log); err != nil {
		panic(err)
	}

	return &config{app, log}
}

func (c *config) GetEnginePath(key string) string {
	return c.app.Engines[key]
}

func (c *config) GetEngineNames() []string {
	l := make([]string, len(c.app.Engines))
	i := 0
	for name := range c.app.Engines {
		l[i] = name
		i++
	}
	return l
}

func (c *config) GetLogConf() zap.Config {
	return c.log
}
