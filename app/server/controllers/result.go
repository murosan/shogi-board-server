package controllers

import (
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"net/http"

	"github.com/murosan/shogi-board-server/app/server/context"
)

// GetResult returns thought result of the shogi engine.
func GetResult(sbc *context.Context) func(echo.Context) error {
	return func(c echo.Context) error {
		name := c.QueryParam(ParamEngineName)
		egn, ok := sbc.Engines[name]

		// engine was not found
		if !ok {
			return c.NoContent(http.StatusNotFound)
		}

		sbc.Logger.Info(
			"[GetResult] param check",
			zap.String("name", name),
			zap.String("egn.Name", egn.Name),
		)

		sbc.Logger.Info("[GetResult]", zap.Any("result", egn.Result))

		return c.JSON(http.StatusOK, egn.Result)
	}
}
