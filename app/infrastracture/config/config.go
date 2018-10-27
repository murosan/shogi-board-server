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
	EnginePath string `json:"engine_path"`
}

// TODO: ローダーを分離する
func NewConfig(p string) cnf.Config {
	b, err := ioutil.ReadFile(p)
	if err != nil {
		log.Fatalln("load: " + err.Error())
	}

	var c config

	if err := json.Unmarshal(b, &c); err != nil {
		log.Fatalln("unmarshal: " + err.Error())
	}

	return &c
}

func (c *config) GetEnginePath() string {
	return c.EnginePath
}
