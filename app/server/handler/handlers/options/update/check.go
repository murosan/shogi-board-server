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

type CheckHandler struct {
	es     service.EngineService
	logger logger.Logger
}

func NewCheckHandler(es service.EngineService, logger logger.Logger) handler.Handler {
	return &CheckHandler{es: es, logger: logger}
}

func (hdr *CheckHandler) Func(ctx *handler.Context) error {
	var option engine.Check
	if err := ctx.Bind(&option); err != nil {
		return framework.NewBadRequestError("body required", err)
	}

	err := handlers.WithEngineID(ctx, func(id engine.ID) error {
		return hdr.es.UpdateCheckOption(id, &option)
	})

	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusOK)
}

func (*CheckHandler) Description() string {
	return "" // TODO
}

func (*CheckHandler) Methods() []string {
	return []string{
		http.MethodPost,
	}
}
