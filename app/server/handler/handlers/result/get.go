package result

import (
	"net/http"

	"github.com/murosan/shogi-board-server/app/domain/service"
	"github.com/murosan/shogi-board-server/app/logger"
	"github.com/murosan/shogi-board-server/app/server/handler"
	"github.com/murosan/shogi-board-server/app/server/handler/handlers"
)

// GetHandler is a handler for getting thought result of the engine.
type GetHandler struct {
	es     service.EngineService
	logger logger.Logger
}

func NewGetHandler(es service.EngineService, logger logger.Logger) handler.Handler {
	return &GetHandler{es: es, logger: logger}
}

func (hdr *GetHandler) Func(ctx *handler.Context) error {
	id, err := handlers.GetEngineID(ctx)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, hdr.es.GetResult(id))
}

func (*GetHandler) Description() string {
	return "" // TODO
}

func (*GetHandler) Methods() []string {
	return []string{
		http.MethodHead,
		http.MethodGet,
	}
}
