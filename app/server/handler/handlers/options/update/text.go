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

type TextHandler struct {
	es     service.EngineService
	logger logger.Logger
}

func NewTextHandler(es service.EngineService, logger logger.Logger) handler.Handler {
	return &TextHandler{es: es, logger: logger}
}

func (hdr *TextHandler) Func(ctx *handler.Context) error {
	var option engine.Text
	if err := ctx.Bind(&option); err != nil {
		return framework.NewBadRequestError("body required", err)
	}

	err := handlers.WithEngineID(ctx, func(id engine.ID) error {
		return hdr.es.UpdateTextOption(id, &option)
	})

	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusOK)
}

func (*TextHandler) Description() string {
	return "" // TODO
}

func (*TextHandler) Methods() []string {
	return []string{
		http.MethodPost,
	}
}
