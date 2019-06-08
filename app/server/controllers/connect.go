package controllers

import (
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"net/http"
	"os/exec"

	"github.com/murosan/shogi-board-server/app/domain/entity/engine"
	"github.com/murosan/shogi-board-server/app/domain/model/infrastructure/os"
	"github.com/murosan/shogi-board-server/app/server/context"
)

// Connect establishes connection with the shogi engine.
// This must be called before executing other usi commands.
func Connect(sbc *context.Context) func(echo.Context) error {
	return func(c echo.Context) error {
		// get engine name from query parameter,
		// then check the name exists in configuration
		name := c.QueryParam(ParamEngineName)
		egnPath, ok := sbc.Config.App.Engines[name]

		// engine name was not found
		if !ok {
			return c.NoContent(http.StatusNotFound)
		}

		sbc.Logger.Info(
			"[Connect] param check",
			zap.String("name", name),
			zap.String("path", egnPath),
		)

		// return ok if the engine is already running
		if _, ok := sbc.Engines[name]; ok {
			sbc.Logger.Info(
				"[Connect] state check",
				zap.String("msg", "already running"),
			)
			return c.NoContent(http.StatusOK)
		}

		// create command
		cmd := exec.Command(egnPath)
		osCmd, err := os.New(cmd)

		if err != nil {
			sbc.Logger.Error("[Connect] initialize os.Cmd", zap.Error(err))
			return c.NoContent(http.StatusInternalServerError)
		}

		// initialize the engine
		egn, err := engine.New(name, osCmd, sbc.Logger)

		if err != nil {
			sbc.Logger.Error("[Connect] initialize Engine", zap.Error(err))
			return c.NoContent(http.StatusInternalServerError)
		}

		// set the initialized engine to Context
		sbc.Engines[name] = egn

		return c.NoContent(http.StatusOK)
	}
}
