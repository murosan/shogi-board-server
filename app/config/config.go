// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

import (
	"encoding/json"
	"io/ioutil"
	"log"

	cnf "github.com/murosan/shogi-proxy-server/app/domain/entity/config"
)

type config struct {
	EnginePath map[string]string `json:"engine_path"`
}

// TODO: ローダーを分離する
func NewConfig() cnf.Config {
	b, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatalln("load: " + err.Error())
	}

	var c config

	if err := json.Unmarshal(b, &c); err != nil {
		log.Fatalln("unmarshal: " + err.Error())
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
