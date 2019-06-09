// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package controllers

import (
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"net/http"

	"github.com/murosan/shogi-board-server/app/domain/entity/engine"
	"github.com/murosan/shogi-board-server/app/server/context"
)

// Close closes the connection of shogi engine.
func Close(sbc *context.Context) func(echo.Context) error {
	return func(c echo.Context) error {
		name := c.QueryParam(ParamEngineName)
		egn, ok := sbc.Engines[name]

		if !ok {
			return c.NoContent(http.StatusNotFound)
		}

		sbc.Logger.Info(
			"[Close]",
			zap.String("name", name),
			zap.String("target key", egn.Key),
		)

		if err := closeEngine(sbc, egn); err != nil {
			return c.NoContent(http.StatusInternalServerError)
		}

		return c.NoContent(http.StatusOK)
	}
}

func closeEngine(c *context.Context, e *engine.Engine) error {
	if e.State == engine.NotConnected {
		return nil
	}

	// exec quit command
	if err := e.Close(); err != nil {
		werr := errors.Wrap(err, "failed to close")
		c.Logger.Error("[closeEngine]", zap.Error(werr))
		return werr
	}

	defer delete(c.Engines, e.Key)
	return nil
}
