// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package controllers

import (
	"net/http"

	"github.com/murosan/shogi-board-server/app/domain/entity/engine"
	"github.com/murosan/shogi-board-server/app/domain/model/exception"
	"github.com/murosan/shogi-board-server/app/domain/model/usi"
	"github.com/murosan/shogi-board-server/app/server/context"

	"github.com/labstack/echo"
	"go.uber.org/zap"
)

// Close closes the connection of shogi engine.
func Close(sbc *context.Context) func(echo.Context) error {
	return func(c echo.Context) error {
		name := c.Param(ParamEngineName)
		egn, ok := sbc.Engines[name]

		if !ok {
			return c.JSON(http.StatusNotFound, "")
		}

		if err := closeEngine(sbc, egn); err != nil {
			return c.JSON(http.StatusInternalServerError, "")
		}

		return c.JSON(http.StatusOK, "")
	}
}

func closeEngine(c *context.Context, e engine.Engine) error {
	// nothing to do
	if e.State == engine.NotConnected {
		return nil
	}

	// exec quit command
	if err := e.Cmd.Write(usi.Quit); err != nil {
		err = exception.FailedToClose(err)
		c.Logger.Error("failed to exec", zap.Error(err))
		return err
	}

	defer delete(c.Engines, e.Key)
	return e.Cmd.Wait()
}
