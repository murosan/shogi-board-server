package controllers

import (
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"net/http"
	"strconv"

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
		name := c.QueryParam(ParamEngineName)
		egn, ok := sbc.Engines[name]

		if !ok {
			return c.NoContent(http.StatusNotFound)
		}

		optName := c.Param("name")
		v := c.Param("value")

		sbc.Logger.Info(
			"[UpdateButton]",
			zap.String("value", v),
			zap.String("name", optName),
		)

		btn, ok := egn.Options.Buttons[optName]
		if !ok {
			return c.NoContent(http.StatusNotFound)
		}

		if err := egn.Cmd.Write([]byte(btn.ToUSI())); err != nil {
			sbc.Logger.Error("[UpdateButton]", zap.Error(err))
			return c.NoContent(http.StatusInternalServerError)
		}

		return c.NoContent(http.StatusOK)
	}
}

// UpdateCheck executes setoption USI command
func UpdateCheck(sbc *context.Context) func(echo.Context) error {
	return func(c echo.Context) error {
		name := c.QueryParam(ParamEngineName)
		egn, ok := sbc.Engines[name]

		if !ok {
			return c.NoContent(http.StatusNotFound)
		}

		optName := c.Param("name")
		v := c.Param("value")

		sbc.Logger.Info(
			"[UpdateCheck]",
			zap.String("value", v),
			zap.String("name", optName),
		)

		chk, ok := egn.Options.Checks[optName]

		if !ok || (v != "true" && v != "false") {
			return c.NoContent(http.StatusNotFound)
		}

		chk.Set(v == "true")

		if err := egn.Cmd.Write([]byte(chk.ToUSI())); err != nil {
			sbc.Logger.Error("[UpdateCheck]", zap.Error(err))
			return c.NoContent(http.StatusInternalServerError)
		}

		return c.NoContent(http.StatusOK)
	}
}

// UpdateRange executes setoption USI command
func UpdateRange(sbc *context.Context) func(echo.Context) error {
	return func(c echo.Context) error {
		name := c.QueryParam(ParamEngineName)
		egn, ok := sbc.Engines[name]

		if !ok {
			return c.NoContent(http.StatusNotFound)
		}

		optName := c.Param("name")
		v := c.Param("value")

		sbc.Logger.Info(
			"[UpdateRange]",
			zap.String("value", v),
			zap.String("name", optName),
		)

		rng, ok := egn.Options.Ranges[optName]

		if !ok {
			return c.NoContent(http.StatusNotFound)
		}

		n, err := strconv.Atoi(v)

		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		if err := rng.Set(n); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		if err := egn.Cmd.Write([]byte(rng.ToUSI())); err != nil {
			sbc.Logger.Error("[UpdateRange]", zap.Error(err))
			return c.NoContent(http.StatusInternalServerError)
		}

		return c.NoContent(http.StatusOK)
	}
}

// UpdateSelect executes setoption USI command
func UpdateSelect(sbc *context.Context) func(echo.Context) error {
	return func(c echo.Context) error {
		name := c.QueryParam(ParamEngineName)
		egn, ok := sbc.Engines[name]

		if !ok {
			return c.NoContent(http.StatusNotFound)
		}

		optName := c.Param("name")
		v := c.Param("value")

		sbc.Logger.Info(
			"[UpdateSelect]",
			zap.String("value", v),
			zap.String("name", optName),
		)

		sel, ok := egn.Options.Selects[optName]

		if !ok {
			return c.NoContent(http.StatusNotFound)
		}

		if err := sel.Set(v); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		if err := egn.Cmd.Write([]byte(sel.ToUSI())); err != nil {
			sbc.Logger.Error("[UpdateSelect]", zap.Error(err))
			return c.NoContent(http.StatusInternalServerError)
		}

		return c.NoContent(http.StatusOK)
	}
}

// UpdateText executes setoption USI command
func UpdateText(sbc *context.Context) func(echo.Context) error {
	return func(c echo.Context) error {
		name := c.QueryParam(ParamEngineName)
		egn, ok := sbc.Engines[name]

		if !ok {
			return c.NoContent(http.StatusNotFound)
		}

		optName := c.Param("name")
		v := c.Param("value")

		sbc.Logger.Info(
			"[UpdateText]",
			zap.String("value", v),
			zap.String("name", optName),
		)

		txt, ok := egn.Options.Texts[optName]

		if !ok {
			return c.NoContent(http.StatusNotFound)
		}

		txt.Set(v)

		if err := egn.Cmd.Write([]byte(txt.ToUSI())); err != nil {
			sbc.Logger.Error("[UpdateText]", zap.Error(err))
			return c.NoContent(http.StatusInternalServerError)
		}

		return c.NoContent(http.StatusOK)
	}
}
