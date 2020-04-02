package handlers

import (
	"net/http"

	"github.com/murosan/shogi-board-server/app/domain/service"
	"github.com/murosan/shogi-board-server/app/logger"
	"github.com/murosan/shogi-board-server/app/server/handler"
)

// ConnectHandler is a handler for connecting (initializing) to the engine.
// When the engine state is Connected, StandBy or Thinking, then returns BAD_REQUEST
// See domain/entity/engine/state.go about engine state.
type ConnectHandler struct {
	es     service.EngineService
	logger logger.Logger
}

func NewConnectHandler(es service.EngineService, logger logger.Logger) handler.Handler {
	return &ConnectHandler{es: es, logger: logger}
}

func (hdr *ConnectHandler) Func(ctx *handler.Context) error {
	if err := WithEngineID(ctx, hdr.es.Connect); err != nil {
		return err
	}
	return ctx.NoContent(http.StatusOK)
}

func (*ConnectHandler) Description() string {
	return "" // TODO
}

func (*ConnectHandler) Methods() []string {
	return []string{
		http.MethodPost,
	}
}
