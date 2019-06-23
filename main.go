// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/murosan/shogi-board-server/app/server/controllers"
	"github.com/murosan/shogi-board-server/app/services/config"
	"github.com/murosan/shogi-board-server/app/services/context"
	"github.com/murosan/shogi-board-server/app/services/logger"
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

	config.Init(*appConfigPath, *logConfigPath)
	logger.Init(config.Use())
	context.Init(logger.Use(), config.Use())

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	e.GET("/ok", ok)
	e.HEAD("/ok", ok)

	ctx := context.Use()

	e.POST("/init", controllers.Init(ctx))
	e.POST("/connect", controllers.Connect(ctx))
	e.POST("/close", controllers.Close(ctx))
	e.POST("/start", controllers.Start(ctx))
	e.POST("/stop", controllers.Stop(ctx))
	e.GET("/options/get", controllers.GetOptions(ctx))
	e.POST("/options/update/button", controllers.UpdateButton(ctx))
	e.POST("/options/update/check", controllers.UpdateCheck(ctx))
	e.POST("/options/update/range", controllers.UpdateRange(ctx))
	e.POST("/options/update/select", controllers.UpdateSelect(ctx))
	e.POST("/options/update/text", controllers.UpdateText(ctx))
	e.GET("/result/get", controllers.GetResult(ctx))
	e.POST("/position/set", controllers.SetPosition(ctx))

	err := e.Start(":" + *port)
	if err != nil {
		panic(err)
	}
}

func ok(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}
