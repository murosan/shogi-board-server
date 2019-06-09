package controllers

import (
	"bytes"
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"net/http"

	"github.com/murosan/shogi-board-server/app/domain/entity/engine"
	"github.com/murosan/shogi-board-server/app/domain/model/usi"
	usiParser "github.com/murosan/shogi-board-server/app/lib/parser/usi"
	"github.com/murosan/shogi-board-server/app/server/context"
)

// Start lets the shogi engine start thinking.
func Start(sbc *context.Context) func(echo.Context) error {
	return func(c echo.Context) error {
		name := c.QueryParam(ParamEngineName)
		egn, ok := sbc.Engines[name]

		// engine was not found
		if !ok {
			return c.NoContent(http.StatusNotFound)
		}

		sbc.Logger.Info(
			"[Start] param check",
			zap.String("name", name),
			zap.String("egn.Name", egn.Name),
		)

		if err := start(sbc, egn); err != nil {
			sbc.Logger.Error("[Start] failed to start", zap.Error(err))
			return c.NoContent(http.StatusInternalServerError)
		}

		return c.NoContent(http.StatusOK)
	}
}

func start(sbc *context.Context, egn *engine.Engine) error {
	if egn.State == engine.Thinking {
		sbc.Logger.Info("[start]", zap.String("nothing to do", ""))
		return nil
	}

	// if state is connected, execute 'usinewgame'
	if egn.State == engine.Connected {
		if err := egn.Cmd.Write(usi.NewGame); err != nil {
			sbc.Logger.Error("[start] failed to start", zap.Error(err))
			return err
		}
		egn.State = engine.StandBy
	}

	// start thinking
	if err := egn.Cmd.Write(usi.GoInf); err != nil {
		sbc.Logger.Error("[start] failed to start", zap.Error(err))
		return err
	}

	// receives outputs while thinking, and set those to result of engine.
	go func() {
		for b := range egn.Ch {
			sbc.Logger.Info("[engine output]", zap.ByteString("msg", b))

			// ignore 'info string'
			if bytes.HasPrefix(b, []byte("info string")) {
				continue
			}

			if bytes.HasPrefix(b, []byte("info ")) {
				i, mpv, err := usiParser.ParseInfo(string(b))
				if err != nil {
					sbc.Logger.Error("[start]", zap.Error(err))
					continue
				}

				sbc.Logger.Info("[start]", zap.Any("parsed", i))

				if mpv == 1 {
					// If mpv is 1, it means 'best move' usually.
					// If the number of candidates is reduced from 5 to 2,
					// there will be extra information left, so delete
					egn.FlushResult()
				}
				if len(i.Moves) != 0 {
					egn.Result[mpv] = i
				}
			}
		}
	}()

	egn.State = engine.Thinking
	return nil
}
