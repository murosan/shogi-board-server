package controllers

import (
	"github.com/labstack/echo"
	"net/http"

	"github.com/murosan/shogi-board-server/app/server/context"
)

// GetOptions returns options of the shogi engine.
func GetOptions(sbc *context.Context) func(echo.Context) error {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "")
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
