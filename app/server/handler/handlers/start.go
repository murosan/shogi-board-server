package handlers

import (
	"net/http"

	"github.com/murosan/shogi-board-server/app/domain/service"
	"github.com/murosan/shogi-board-server/app/logger"
	"github.com/murosan/shogi-board-server/app/server/handler"
)

// StartHandler is a handler for starting the engine think.
// When the engine state is
//   - NotConnected, then returns NOT_FOUND or BAD_REQUEST
//   - Connected or StandBy, then starts thinking and returns OK
//   - Thinking, then returns OK (do nothing)
// See domain/entity/engine/state.go about engine state.
type StartHandler struct {
	es     service.EngineService
	logger logger.Logger
}

func NewStartHandler(es service.EngineService, logger logger.Logger) handler.Handler {
	return &StartHandler{es: es, logger: logger}
}

func (hdr *StartHandler) Func(ctx *handler.Context) error {
	if err := WithEngineID(ctx, hdr.es.Start); err != nil {
		return err
	}
	return ctx.NoContent(http.StatusOK)
}

func (*StartHandler) Description() string {
	return "" // TODO
}

func (*StartHandler) Methods() []string {
	return []string{
		http.MethodPost,
	}
}
