package controllers

import (
	"github.com/labstack/echo"
	"net/http"

	"github.com/murosan/shogi-board-server/app/server/context"
)

// Connect establishes connection with the shogi engine.
// This must be called before executing other usi commands.
func Connect(sbc *context.Context) func(echo.Context) error {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "")
	}
}
