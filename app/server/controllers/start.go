package controllers

import (
	"github.com/labstack/echo"
	"net/http"

	"github.com/murosan/shogi-board-server/app/server/context"
)

// Start lets the shogi engine start thinking.
func Start(sbc *context.Context) func(echo.Context) error {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "")
	}
}
