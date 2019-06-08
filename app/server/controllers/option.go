package controllers

import (
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"net/http"

	"github.com/murosan/shogi-board-server/app/server/context"
)

// GetOptions returns options of the shogi engine.
func GetOptions(sbc *context.Context) func(echo.Context) error {
	return func(c echo.Context) error {
		// get engine name from query parameter,
		// then check the name exists in configuration
		name := c.QueryParam(ParamEngineName)
		egn, ok := sbc.Engines[name]

		// engine name was not found
		if !ok {
			return c.NoContent(http.StatusNotFound)
		}

		sbc.Logger.Info("[GetOptions] param check", zap.String("name", name))
		sbc.Logger.Info("[GetOptions] options", zap.Any("opts", egn.Options))

		return c.JSON(http.StatusOK, egn.Options)
	}
}

// UpdateButton executes setoption USI command
func UpdateButton(sbc *context.Context) func(echo.Context) error {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "")
	}
}

// UpdateCheck executes setoption USI command
func UpdateCheck(sbc *context.Context) func(echo.Context) error {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "")
	}
}

// UpdateRange executes setoption USI command
func UpdateRange(sbc *context.Context) func(echo.Context) error {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "")
	}
}

// UpdateSelect executes setoption USI command
func UpdateSelect(sbc *context.Context) func(echo.Context) error {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "")
	}
}

// UpdateString executes setoption USI command
func UpdateString(sbc *context.Context) func(echo.Context) error {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "")
	}
}
