// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package controllers

import (
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"net/http"

	"github.com/murosan/shogi-board-server/app/domain/model/exception"
	"github.com/murosan/shogi-board-server/app/server/context"
)

// Init closes all engine connection and returns the list of engine names
// configured in configuration file.
// This should be called before any other APIs.
func Init(sbc *context.Context) func(echo.Context) error {
	return func(c echo.Context) error {
		for _, ngn := range sbc.Engines {
			sbc.Logger.Info("[Init]", zap.String("close target", ngn.Key))

			if err := closeEngine(sbc, ngn); err != nil {
				msg := "failed to close engine. name = " + ngn.Name
				e := exception.FailedToClose(err)
				sbc.Logger.Error(msg, zap.Error(e))
				return echo.NewHTTPError(http.StatusInternalServerError, msg)
			}
		}

		names := sbc.Config.App.EngineNames
		sbc.Logger.Info("[Init]", zap.Strings("names", names))
		return c.JSON(http.StatusOK, names)
	}
}
