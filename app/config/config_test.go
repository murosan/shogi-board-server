// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

import (
	"os"
	"path"
	"testing"

	smap "github.com/murosan/goutils/map/strings"
	sslice "github.com/murosan/goutils/slice/strings"
	"github.com/murosan/goutils/testutils"
)

var (
	pwd, _  = os.Getwd()
	dataDir = "./testdata"
)

func TestNew(t *testing.T) {
	cases := []struct {
		appPath string
		logPath string
		app     App
	}{
		{
			path.Join(pwd, dataDir, "app.config.yml"),
			path.Join(pwd, dataDir, "log.config.yml"),
			App{
				Engines:     map[string]string{"com": "/home/user/path/to/engine/bin"},
				EngineNames: []string{"com"},
			},
		},
		{
			path.Join(pwd, dataDir, "app.config.yml"),
			"",
			App{
				Engines:     map[string]string{"com": "/home/user/path/to/engine/bin"},
				EngineNames: []string{"com"},
			},
		},
	}

	for i, c := range cases {
		conf := New(c.appPath, c.logPath)

		failed := func(key string, expected, actual interface{}) {
			t.Helper()
			t.Errorf(`
[app > config > TestNew] %s was not equal to as expected.
Index:    %d
Expected: %v
Actual:   %v
`, key, i, expected, actual)
		}

		if sslice.NotEqual(conf.App.EngineNames, c.app.EngineNames) {
			failed("EngineNames", c.app.EngineNames, conf.App.EngineNames)
		}
		if smap.NotEqual(conf.App.Engines, c.app.Engines) {
			failed("Engines", c.app.Engines, conf.App.Engines)
		}
	}
}

func TestNew2(t *testing.T) {
	cases := []struct {
		appPath string
		logPath string
	}{
		{
			path.Join(pwd, dataDir, "app_invalid.config.yml"),
			path.Join(pwd, dataDir, "log.config.yml"),
		},
		{
			path.Join(pwd, dataDir, "app.config.yml"),
			path.Join(pwd, dataDir, "log_invalid.config.yml"),
		},
		{
			path.Join(pwd, dataDir, "app_empty.config.yml"),
			path.Join(pwd, dataDir, "log.config.yml"),
		},
		{
			path.Join(pwd, dataDir),
			path.Join(pwd, dataDir, "log.config.yml"),
		},
		{
			path.Join(pwd, dataDir, "app.config.yml"),
			path.Join(pwd, dataDir),
		},
	}

	for i, c := range cases {
		causePanic := func() { New(c.appPath, c.logPath) }
		onFail := func(t *testing.T) {
			t.Helper()
			t.Errorf(`
[app > config > TestNew2] Expected panic but none was found.
Index:        %d
InputAppPath: %s
InputLogPath: %s
`, i, c.appPath, c.logPath)
		}
		testutils.MustPanic(t, causePanic, onFail)
	}
}
