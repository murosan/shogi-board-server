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

func Load() {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln("load: " + err.Error())
	}

	if err := json.Unmarshal(b, &Conf); err != nil {
		log.Fatalln("unmarshal: " + err.Error())
	}
}
