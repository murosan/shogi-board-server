package routes

import (
	"github.com/labstack/echo"

	"github.com/murosan/shogi-board-server/app/domain/config"
	"github.com/murosan/shogi-board-server/app/domain/service"
	"github.com/murosan/shogi-board-server/app/logger"
	"github.com/murosan/shogi-board-server/app/server/handler"
	"github.com/murosan/shogi-board-server/app/server/handler/handlers"
	"github.com/murosan/shogi-board-server/app/server/handler/handlers/options"
	"github.com/murosan/shogi-board-server/app/server/handler/handlers/options/update"
	"github.com/murosan/shogi-board-server/app/server/handler/handlers/position"
	"github.com/murosan/shogi-board-server/app/server/handler/handlers/result"
)

// Initialize setups server routes.
func Initialize(
	e *echo.Echo,
	config *config.Config,
	logger logger.Logger,
	es service.EngineService,
) {
	routes := []route{
		{path: "/ok", handler: handlers.NewOKHandler()},
		{path: "/init", handler: handlers.NewInitHandler(es, config, logger)},
		{path: "/connect", handler: handlers.NewConnectHandler(es, logger)},
		{path: "/close", handler: handlers.NewCloseHandler(es, logger)},
		{path: "/start", handler: handlers.NewStartHandler(es, logger)},
		{path: "/stop", handler: handlers.NewStopHandler(es, logger)},
		{path: "/options/get", handler: options.NewGetHandler(es, logger)},
		{path: "/options/update/button", handler: update.NewButtonHandler(es, logger)},
		{path: "/options/update/check", handler: update.NewCheckHandler(es, logger)},
		{path: "/options/update/range", handler: update.NewRangeHandler(es, logger)},
		{path: "/options/update/select", handler: update.NewSelectHandler(es, logger)},
		{path: "/options/update/text", handler: update.NewTextHandler(es, logger)},
		{path: "/result/get", handler: result.NewGetHandler(es, logger)},
		{path: "/position/set", handler: position.NewSetHandler(es, logger)},
	}

	for _, r := range routes {
		path, hdr := r.path, r.handler
		e.Match(hdr.Methods(), path, handler.Build(hdr, logger))
	}
}

type route struct {
	path    string
	handler handler.Handler
}
