package position

import (
	"net/http"

	"github.com/murosan/shogi-board-server/app/domain/entity/engine"
	"github.com/murosan/shogi-board-server/app/domain/entity/shogi"
	"github.com/murosan/shogi-board-server/app/domain/framework"
	"github.com/murosan/shogi-board-server/app/domain/service"
	"github.com/murosan/shogi-board-server/app/logger"
	"github.com/murosan/shogi-board-server/app/server/handler"
	"github.com/murosan/shogi-board-server/app/server/handler/handlers"
)

// SetHandler is a handler for setting the new position.
// This handler requires the position body as JSON.
// See domain/entity/shogi/position.go about position.
type SetHandler struct {
	es     service.EngineService
	logger logger.Logger
}

func NewSetHandler(es service.EngineService, logger logger.Logger) handler.Handler {
	return &SetHandler{es: es, logger: logger}
}

func (hdr *SetHandler) Func(ctx *handler.Context) error {
	var pos shogi.Position
	if err := ctx.Bind(&pos); err != nil {
		return framework.NewBadRequestError("body required", err)
	}

	err := handlers.WithEngineID(ctx, func(id engine.ID) error {
		return hdr.es.UpdatePosition(id, &pos)
	})

	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusOK)
}

func (*SetHandler) Description() string {
	return "" // TODO
}

func (*SetHandler) Methods() []string {
	return []string{
		http.MethodPost,
	}
}
