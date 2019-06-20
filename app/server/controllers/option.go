package controllers

import (
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"net/http"

	"github.com/murosan/shogi-board-server/app/domain/entity/option"
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
		var b option.Button

		if err := c.Bind(&b); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		name := c.QueryParam(ParamEngineName)
		egn, ok := sbc.Engines[name]

		if !ok {
			return c.NoContent(http.StatusNotFound)
		}

		sbc.Logger.Info("[UpdateButton]", zap.Any("button", b))

		btn, ok := egn.Options.Buttons[b.Name]
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
		var v option.Check

		if err := c.Bind(&v); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		name := c.QueryParam(ParamEngineName)
		egn, ok := sbc.Engines[name]

		if !ok {
			return c.NoContent(http.StatusNotFound)
		}

		sbc.Logger.Info("[UpdateCheck]", zap.Any("check", v))

		chk, ok := egn.Options.Checks[v.Name]

		if !ok {
			return c.NoContent(http.StatusNotFound)
		}

		chk.Set(v.Value)

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
		var v option.Range

		if err := c.Bind(&v); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		name := c.QueryParam(ParamEngineName)
		egn, ok := sbc.Engines[name]

		if !ok {
			return c.NoContent(http.StatusNotFound)
		}

		sbc.Logger.Info("[UpdateRange]", zap.Any("range", v))

		rng, ok := egn.Options.Ranges[v.Name]

		if !ok {
			sbc.Logger.Info("[UpdateRange]", zap.String("option not found", v.Name))
			return c.NoContent(http.StatusNotFound)
		}

		if err := rng.Set(v.Value); err != nil {
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
		var v option.Select

		if err := c.Bind(&v); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		name := c.QueryParam(ParamEngineName)
		egn, ok := sbc.Engines[name]

		if !ok {
			return c.NoContent(http.StatusNotFound)
		}

		sbc.Logger.Info("[UpdateSelect]", zap.Any("select", v))

		sel, ok := egn.Options.Selects[v.Name]

		if !ok {
			return c.NoContent(http.StatusNotFound)
		}

		if err := sel.Set(v.Value); err != nil {
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
		var v option.Text

		if err := c.Bind(&v); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		name := c.QueryParam(ParamEngineName)
		egn, ok := sbc.GetEngine(name)

		if !ok {
			return c.NoContent(http.StatusNotFound)
		}

		sbc.Logger.Info("[UpdateText]", zap.Any("text", v))

		txt, ok := egn.Options.Texts[v.Name]

		if !ok {
			return c.NoContent(http.StatusNotFound)
		}

		txt.Set(v.Value)

		if err := egn.Cmd.Write([]byte(txt.ToUSI())); err != nil {
			sbc.Logger.Error("[UpdateText]", zap.Error(err))
			return c.NoContent(http.StatusInternalServerError)
		}

		return c.NoContent(http.StatusOK)
	}
}
