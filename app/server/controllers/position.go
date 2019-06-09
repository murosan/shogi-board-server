package controllers

import (
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"net/http"

	"github.com/murosan/shogi-board-server/app/domain/model/shogi"
	"github.com/murosan/shogi-board-server/app/server/context"
)

// SetPosition returns options of the shogi engine.
func SetPosition(sbc *context.Context) func(echo.Context) error {
	return func(c echo.Context) error {
		name := c.QueryParam(ParamEngineName)
		egn, ok := sbc.Engines[name]

		// engine was not found
		if !ok {
			return c.NoContent(http.StatusNotFound)
		}

		sbc.Logger.Info(
			"[SetPosition] param check",
			zap.String("name", name),
			zap.String("egn.Name", egn.Name),
		)

		var p shogi.Position
		if err := c.Bind(&p); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		sbc.Logger.Info("[SetPosition]", zap.Any("position", p))

		usi, err := p.ToUSI()
		if err != nil {
			sbc.Logger.Warn("[SetPosition] failed to convert to usi", zap.Error(err))
			return c.NoContent(http.StatusBadRequest)
		}

		sbc.Logger.Info("[SetPosition]", zap.ByteString("usi", usi))

		if err := egn.Cmd.Write(usi); err != nil {
			sbc.Logger.Error(
				"[SetPosition] error at write position command",
				zap.Error(err),
			)
			return c.NoContent(http.StatusInternalServerError)
		}

		return c.NoContent(http.StatusOK)
	}
}
