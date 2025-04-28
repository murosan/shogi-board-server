package handler

import (
	"errors"
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

	var e *framework.ControllerError
	if errors.As(err, &e) {
		if errors.Is(e, framework.ErrInternalServerError) {
			ctx.logger.Error(e.Message, zap.Error(err))
		} else {
			ctx.logger.Info(e.Message, zap.Error(err))
		}
		return errJSON(ctx, e.Status, err)
	}

	ctx.logger.Error(fmt.Sprint(http.StatusInternalServerError), zap.Error(err))
	return errJSON(ctx, http.StatusInternalServerError, err)
}

func errJSON(ctx *Context, status int, err error) error {
	return ctx.JSON(
		status,
		map[string]any{
			"status":  status,
			"message": err.Error(),
		},
	)
}
