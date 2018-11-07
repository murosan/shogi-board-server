// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

import (
	"sort"
	"testing"

	confModel "github.com/murosan/shogi-proxy-server/app/domain/entity/config"
	"github.com/murosan/shogi-proxy-server/app/lib/stringutil"
)

func TestNewConfig(t *testing.T) {
	cases := []struct {
		json        string
		enginePaths []string
		engineNames []string
		log         confModel.LogConfig
	}{
		{`
{
  "engine_path": {
    "com": "/home/user/path/to/engine"
  },
  "log": {
    "output_type": "stdout",
    "level": "debug"
  }
}
`, []string{"/home/user/path/to/engine"}, []string{"com"}, confModel.LogConfig{"stdout", "debug"}},
	}

	for i, c := range cases {
		conf := NewConfig([]byte(c.json))
		names := conf.GetEngineNames()
		sort.Strings(names)

		// GetEngineNames() と GetEnginePath() のテスト
		if stringutil.SliceEquals(conf.GetEngineNames(), c.engineNames) {
			for j := range names {
				p1 := conf.GetEnginePath(names[j])
				p2 := c.enginePaths[j]
				if p1 != p2 {
					t.Errorf(`EnginePath was not equal to as expected.
i, %d
Expected: %s
Actual:   %s`, j, p2, p1)
				}
			}
		} else {
			t.Errorf(`EngineNames was not equal to as expected.
i: %d
Expected: %v
Actual:   %v`, i, c.engineNames, names)
		}
	}
}
