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

type ButtonHandler struct {
	es     service.EngineService
	logger logger.Logger
}

func NewButtonHandler(es service.EngineService, logger logger.Logger) handler.Handler {
	return &ButtonHandler{es: es, logger: logger}
}

func (hdr *ButtonHandler) Func(ctx *handler.Context) error {
	var option engine.Button
	if err := ctx.Bind(&option); err != nil {
		return framework.ErrBadRequest.With("body required").WithErr(err)
	}

	err := handlers.WithEngineID(ctx, func(id engine.ID) error {
		return hdr.es.UpdateButtonOption(id, &option)
	})

	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusOK)
}

func (*ButtonHandler) Description() string {
	return "" // TODO
}

func (*ButtonHandler) Methods() []string {
	return []string{
		http.MethodPost,
	}
}
