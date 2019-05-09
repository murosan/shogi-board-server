// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

import (
	"io/ioutil"

	"go.uber.org/zap"
	"gopkg.in/yaml.v2"

	smap "github.com/murosan/goutils/map/strings"
)

// Config is container of configurations.
type Config struct {
	// application config
	App *App

	// zap(logger) config
	Log *zap.Config
}

// App is configuration of application
type App struct {
	// Engines is a map of engine name and engine executable path pairs.
	// This map must not be empty. If empty, panics when called New method.
	//
	// The key is the name of shogi engine. Any string is OK to use.
	// The value is the path of shogi engine.
	// It must be executable, and work with USI protocol.
	//
	// TODO: write document about USI protocol
	Engines map[string]string `yaml:"engines"`

	// Keys of Engines
	EngineNames []string
}

func New(appPath, logPath string) *Config {
	var app App
	var log zap.Config

	// load app config path as byte array
	a, err := ioutil.ReadFile(appPath)
	if err != nil {
		panic(err)
	}

	// load log config path as byte array
	l, err := ioutil.ReadFile(logPath)
	if err != nil {
		panic(err)
	}

	// convert YAML to App
	if err := yaml.Unmarshal(a, &app); err != nil {
		panic(err)
	}

	// convert YAML to zap.Config
	if err := yaml.Unmarshal(l, &log); err != nil {
		panic(err)
	}

	if len(app.Engines) == 0 {
		panic("Engines is empty. You must specify at least one shogi engine.")
	}

	app.EngineNames = smap.Keys(app.Engines)

	return &Config{App: &app, Log: &log}
}
