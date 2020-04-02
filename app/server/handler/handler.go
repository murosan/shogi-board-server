package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	"github.com/murosan/shogi-board-server/app/domain/framework"
	"github.com/murosan/shogi-board-server/app/logger"
)

// Handler is provider of basic handler func
type Handler interface {
	Func(ctx *Context) error // Func is http handler
	Description() string     // Description returns API description.
	Methods() []string       // Methods returns allowed http methods
}

// Build returns HandlerFunc
func Build(handler Handler, logger logger.Logger) echo.HandlerFunc {
	return func(ec echo.Context) error {
		ctx := NewContext(ec, logger)
		return handleError(ctx, handler.Func(ctx))
	}
}

func handleError(ctx *Context, err error) error {
	if err == nil {
		return nil
	}

	if _, ok := err.(*framework.BadRequestError); ok {
		ctx.logger.Info(fmt.Sprint(http.StatusBadRequest), zap.Error(err))
		return errJSON(ctx, http.StatusBadRequest, err)
	}
	if _, ok := err.(*framework.NotFoundError); ok {
		ctx.logger.Info(fmt.Sprint(http.StatusNotFound), zap.Error(err))
		return errJSON(ctx, http.StatusNotFound, err)
	}

	ctx.logger.Error(fmt.Sprint(http.StatusInternalServerError), zap.Error(err))
	return errJSON(ctx, http.StatusInternalServerError, err)
}

func errJSON(ctx *Context, status int, err error) error {
	return ctx.JSON(
		status,
		map[string]interface{}{
			"status":  status,
			"message": err.Error(),
		},
	)
}
