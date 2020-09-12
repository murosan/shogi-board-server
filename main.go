// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"os"
	"path"
	"path/filepath"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/murosan/shogi-board-server/app/module"
	"github.com/murosan/shogi-board-server/app/server/handler/routes"
)

var (
	port          = flag.String("port", "8080", "http server port")
	appConfigPath = flag.String(
		"app_config",
		"",
		"application config path. default=config/app.config.yml",
	)
	logConfigPath = flag.String("log_config", "", "log config path. optional")
)

func main() {
	flag.Parse()

	if *appConfigPath == "" {
		acp := filepath.Join(path.Dir(os.Args[0]), "config", "app.config.yml")
		appConfigPath = &acp
	}

	e := echo.New()
	// e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	module.Initialize(*appConfigPath, *logConfigPath)
	routes.Initialize(
		e,
		module.Config,
		module.Logger,
		module.Services.Engine,
	)

	err := e.Start(":" + *port)
	if err != nil {
		panic(err)
	}
}
