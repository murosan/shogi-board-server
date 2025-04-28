package handlers

import (
	"net/http"

	"github.com/murosan/shogi-board-server/app/domain/service"
	"github.com/murosan/shogi-board-server/app/logger"
	"github.com/murosan/shogi-board-server/app/server/handler"
)

// CloseHandler is a handler to close the engine connection.
// When the engine state is
//   - NotConnected, then returns NOT_FOUND or BAD_REQUEST
//   - Connected, StandBy or Thinking, then closes connection and returns OK
//
// See domain/entity/engine/state.go about engine state.
type CloseHandler struct {
	es     service.EngineService
	logger logger.Logger
}

func NewCloseHandler(es service.EngineService, logger logger.Logger) handler.Handler {
	return &CloseHandler{es: es, logger: logger}
}

func (hdr *CloseHandler) Func(ctx *handler.Context) error {
	if err := WithEngineID(ctx, hdr.es.Close); err != nil {
		return err
	}
	return ctx.NoContent(http.StatusOK)
}

func (*CloseHandler) Description() string {
	return "" // TODO
}

func (*CloseHandler) Methods() []string {
	return []string{
		http.MethodPost,
	}
}
