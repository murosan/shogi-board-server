package update

import (
	"net/http"

	"github.com/murosan/shogi-board-server/app/domain/entity/engine"
	"github.com/murosan/shogi-board-server/app/domain/framework"
	"github.com/murosan/shogi-board-server/app/domain/service"
	"github.com/murosan/shogi-board-server/app/logger"
	"github.com/murosan/shogi-board-server/app/server/handler"
	"github.com/murosan/shogi-board-server/app/server/handler/handlers"
)

type SelectHandler struct {
	es     service.EngineService
	logger logger.Logger
}

func NewSelectHandler(es service.EngineService, logger logger.Logger) handler.Handler {
	return &SelectHandler{es: es, logger: logger}
}

func (hdr *SelectHandler) Func(ctx *handler.Context) error {
	var option engine.Select
	if err := ctx.Bind(&option); err != nil {
		return framework.NewBadRequestError("body required", err)
	}

	err := handlers.WithEngineID(ctx, func(id engine.ID) error {
		return hdr.es.UpdateSelectOption(id, &option)
	})

	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusOK)
}

func (*SelectHandler) Description() string {
	return "" // TODO
}

func (*SelectHandler) Methods() []string {
	return []string{
		http.MethodPost,
	}
}
