package controllers

import (
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"net/http"

	"github.com/murosan/shogi-board-server/app/domain/entity/engine"
	"github.com/murosan/shogi-board-server/app/domain/model/usi"
	"github.com/murosan/shogi-board-server/app/server/context"
)

// Stop lets the shogi engine stop thinking.
func Stop(sbc *context.Context) func(echo.Context) error {
	return func(c echo.Context) error {
		name := c.QueryParam(ParamEngineName)
		egn, ok := sbc.GetEngine(name)

		// engine was not found
		if !ok {
			return c.NoContent(http.StatusNotFound)
		}

		sbc.Logger.Info(
			"[Stop] param check",
			zap.String("name", name),
			zap.String("egn.Name", egn.Name),
		)

		// engine is not thinking. nothing to do
		if egn.State.Get() != engine.Thinking {
			sbc.Logger.Info("[Stop]", zap.Any("nothing to do", egn.State))
			return c.NoContent(http.StatusOK)
		}

		// execute 'stop'
		if err := egn.Cmd.Write(usi.Stop); err != nil {
			sbc.Logger.Error("[Stop] error", zap.Error(err))
			return c.NoContent(http.StatusInternalServerError)
		}

		// set state and reset thought result.
		egn.State.Set(engine.StandBy)
		egn.Result.Flush()

		return c.NoContent(http.StatusOK)
	}
}
