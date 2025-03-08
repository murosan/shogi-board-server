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

type RangeHandler struct {
	es     service.EngineService
	logger logger.Logger
}

func NewRangeHandler(es service.EngineService, logger logger.Logger) handler.Handler {
	return &RangeHandler{es: es, logger: logger}
}

func (hdr *RangeHandler) Func(ctx *handler.Context) error {
	var option engine.Range
	if err := ctx.Bind(&option); err != nil {
		return framework.ErrBadRequest.With("body required").WithErr(err)
	}

	err := handlers.WithEngineID(ctx, func(id engine.ID) error {
		return hdr.es.UpdateRangeOption(id, &option)
	})

	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusOK)
}

func (*RangeHandler) Description() string {
	return "" // TODO
}

func (*RangeHandler) Methods() []string {
	return []string{
		http.MethodPost,
	}
}
