package position

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"

	"github.com/murosan/shogi-board-server/app/domain/entity/engine"
	"github.com/murosan/shogi-board-server/app/domain/entity/shogi"
	"github.com/murosan/shogi-board-server/app/domain/framework"
	"github.com/murosan/shogi-board-server/app/domain/service"
	"github.com/murosan/shogi-board-server/app/lib/usi/convert"
	"github.com/murosan/shogi-board-server/app/logger"
	"github.com/murosan/shogi-board-server/app/server/handler"
	"github.com/murosan/shogi-board-server/app/server/handler/handlers"
)

var (
	queryKeys = struct {
		format string
	}{
		format: "format",
	}

	queryValues = struct {
		json,
		sfen string
	}{
		json: "json",
		sfen: "sfen",
	}
)

// GetHandler is a handler for getting the current position.
// Returns NOT_FOUND when the engine does not exists or the game has not started yet.
type GetHandler struct {
	es     service.EngineService
	logger logger.Logger
}

func NewGetHandler(es service.EngineService, logger logger.Logger) handler.Handler {
	return &GetHandler{es: es, logger: logger}
}

func (hdr *GetHandler) Func(ctx *handler.Context) error {
	format := ctx.GetQuery(queryKeys.format)
	if format == "" {
		return framework.NewBadRequestError("please specify format query", nil)
	}

	var pos *shogi.Position
	var ok bool
	err := handlers.WithEngineID(ctx, func(id engine.ID) error {
		pos, ok = hdr.es.GetCurrentPosition(id)
		if !ok {
			return framework.NewNotFoundError("position not found. id="+id.String(), nil)
		}
		return nil
	})

	if err != nil {
		return err
	}

	switch format {
	case queryValues.json:
		return ctx.JSON(http.StatusOK, pos)
	case queryValues.sfen:
		usi, err := convert.Position(pos)
		if err != nil {
			hdr.logger.Error("convert", zap.Any("pos", pos), zap.Error(err))
			return framework.NewInternalServerError("convert position error", err)
		}
		return ctx.Text(http.StatusOK, usi)
	default:
		return framework.NewBadRequestError(
			fmt.Sprintf(
				"unknown format. got=%s. availables=%s,%s",
				format,
				queryValues.json,
				queryValues.sfen,
			),
			nil,
		)
	}
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

type posResp struct {
	Object *shogi.Position `json:"object"`
	Sfen   string          `json:"sfen"`
}
