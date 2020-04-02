package handlers

import (
	"net/http"

	"go.uber.org/zap"

	"github.com/murosan/shogi-board-server/app/domain/config"
	"github.com/murosan/shogi-board-server/app/domain/service"
	"github.com/murosan/shogi-board-server/app/logger"
	"github.com/murosan/shogi-board-server/app/server/handler"
)

// InitHandler is a handler that initializes all engine.
// It closes all engines and returns the list of engine names.
type InitHandler struct {
	es     service.EngineService
	config *config.Config
	logger logger.Logger
}

func NewInitHandler(
	es service.EngineService,
	config *config.Config,
	logger logger.Logger,
) handler.Handler {
	return &InitHandler{es: es, config: config, logger: logger}
}

func (hdr *InitHandler) Func(ctx *handler.Context) error {
	if err := hdr.es.CloseAll(); err != nil {
		return err
	}

	names := hdr.config.App.EngineNames
	hdr.logger.Info("[Init]", zap.Strings("engine names", names))
	return ctx.JSON(http.StatusOK, names)
}

func (*InitHandler) Description() string {
	return "" // TODO
}

func (*InitHandler) Methods() []string {
	return []string{
		http.MethodPost,
	}
}
