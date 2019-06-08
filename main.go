// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"github.com/murosan/shogi-board-server/app/server/controllers"
	"github.com/murosan/shogi-board-server/app/services/context"
	"net/http"

	"github.com/murosan/shogi-board-server/app/services/config"
	"github.com/murosan/shogi-board-server/app/services/logger"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	port          = flag.String("port", "8080", "http server port")
	appConfigPath = flag.String(
		"appConfig",
		"./config/app.yml",
		"application config path",
	)
	logConfigPath = flag.String(
		"logConfig",
		"./config/log.yml",
		"log config path",
	)
)

func main() {
	flag.Parse()

	config.Init(*appConfigPath, *logConfigPath)
	logger.Init(config.Use())
	context.Init(logger.Use(), config.Use())

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/ok", ok)
	e.HEAD("/ok", ok)

	e.POST("/init", controllers.Init(context.Use()))
	e.POST("/connect", controllers.Connect(context.Use()))
	e.POST("/close", controllers.Close(context.Use()))
	e.POST("/start", controllers.Start(context.Use()))
	e.POST("/stop", controllers.Stop(context.Use()))
	e.GET("/options/get", controllers.GetOptions(context.Use()))
	e.POST("/options/update/button", controllers.UpdateButton(context.Use()))
	e.POST("/options/update/check", controllers.UpdateCheck(context.Use()))
	e.POST("/options/update/range", controllers.UpdateRange(context.Use()))
	e.POST("/options/update/select", controllers.UpdateSelect(context.Use()))
	e.POST("/options/update/string", controllers.UpdateString(context.Use()))
	e.GET("/result/get", controllers.GetResult(context.Use()))
	e.POST("/position/set", controllers.SetPosition(context.Use()))

	err := e.Start(":" + *port)
	if err != nil {
		panic(err)
	}
}

func ok(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}
