package controllers

import (
	"github.com/labstack/echo"
	"net/http"

	"github.com/murosan/shogi-board-server/app/server/context"
)

// Stop lets the shogi engine stop thinking.
func Stop(sbc *context.Context) func(echo.Context) error {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "")
	}
}
