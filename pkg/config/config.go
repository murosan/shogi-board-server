// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var (
	path = "./config.json"
)

type Config struct {
	EnginePath string `json:"engine_path"`
}

func NewConfig(p string) Config {
	b, err := ioutil.ReadFile(p)
	if err != nil {
		log.Fatalln("load: " + err.Error())
	}

	var c Config

	if err := json.Unmarshal(b, &c); err != nil {
		log.Fatalln("unmarshal: " + err.Error())
	}

	return c
}
