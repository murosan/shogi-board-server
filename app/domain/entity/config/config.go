// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

// OutputType: stdout | file
// Level: debug | info | warn | error
type LogConfig struct {
	OutputType string `json:"output_type"`
	Level      string `json:"level"`
}

type Config interface {
	GetEnginePath(string) string
	GetEngineNames() []string
	GetLogConf() LogConfig
}
