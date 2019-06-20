package controllers

import (
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"net/http"

	"github.com/murosan/shogi-board-server/app/domain/entity/engine"
	"github.com/murosan/shogi-board-server/app/domain/model/shogi"
	"github.com/murosan/shogi-board-server/app/domain/model/usi"
	"github.com/murosan/shogi-board-server/app/server/context"
)

// SetPosition returns options of the shogi engine.
func SetPosition(sbc *context.Context) func(echo.Context) error {
	return func(c echo.Context) error {
		name := c.QueryParam(ParamEngineName)
		egn, ok := sbc.GetEngine(name)

		// engine was not found
		if !ok {
			return c.NoContent(http.StatusNotFound)
		}

		sbc.Logger.Info(
			"[SetPosition] param check",
			zap.String("name", name),
			zap.String("egn.Name", egn.Name),
		)

		var p shogi.Position
		if err := c.Bind(&p); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		sbc.Logger.Info("[SetPosition]", zap.Any("position", p))

		usipos, err := p.ToUSI()
		if err != nil {
			sbc.Logger.Warn("[SetPosition] failed to convert to usi", zap.Error(err))
			return c.NoContent(http.StatusBadRequest)
		}

		sbc.Logger.Info("[SetPosition]", zap.ByteString("usi", usipos))

		// exec stop if thinking
		isThinking := egn.State.Get() == engine.Thinking
		if isThinking {
			if err := egn.Cmd.Write(usi.Stop); err != nil {
				err2 := errors.Wrap(err, "failed to stop")
				sbc.Logger.Error("[SetPosition]", zap.Error(err2))
				return c.NoContent(http.StatusInternalServerError)
			}
			egn.State.Set(engine.StandBy)
		}

		if err := egn.Cmd.Write(usipos); err != nil {
			sbc.Logger.Error(
				"[SetPosition] error at write position command",
				zap.Error(err),
			)
			return c.NoContent(http.StatusInternalServerError)
		}

		egn.Result.Flush()

		// start thinking if isThinking is true
		if isThinking {
			if err := start(sbc, egn); err != nil {
				sbc.Logger.Error("[SetPosition] failed to restart", zap.Error(err))
				return c.NoContent(http.StatusInternalServerError)
			}
		}

		return c.NoContent(http.StatusOK)
	}
}
