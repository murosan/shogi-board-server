package handlers

import (
	"net/http"

	"github.com/murosan/shogi-board-server/app/domain/service"
	"github.com/murosan/shogi-board-server/app/logger"
	"github.com/murosan/shogi-board-server/app/server/handler"
)

// StopHandler is a handler for stopping the engine think.
// When the engine state is
//   - NotConnected, then returns NOT_FOUND or BAD_REQUEST
//   - Connected or StandBy, then returns OK (do nothing)
//   - Thinking, then stops thinking and returns OK
//
// See domain/entity/engine/state.go about engine state.
type StopHandler struct {
	es     service.EngineService
	logger logger.Logger
}

func NewStopHandler(es service.EngineService, logger logger.Logger) handler.Handler {
	return &StopHandler{es: es, logger: logger}
}

func (hdr *StopHandler) Func(ctx *handler.Context) error {
	if err := WithEngineID(ctx, hdr.es.Stop); err != nil {
		return err
	}
	return ctx.NoContent(http.StatusOK)
}

func (*StopHandler) Description() string {
	return "" // TODO
}

func (*StopHandler) Methods() []string {
	return []string{
		http.MethodPost,
	}
}
