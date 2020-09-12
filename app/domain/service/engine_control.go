package service

import (
	"bytes"
	"strings"
	"time"

	"go.uber.org/zap"

	"github.com/murosan/shogi-board-server/app/domain/entity/engine"
	"github.com/murosan/shogi-board-server/app/domain/entity/shogi"
	"github.com/murosan/shogi-board-server/app/domain/entity/usi"
	"github.com/murosan/shogi-board-server/app/domain/framework"
	"github.com/murosan/shogi-board-server/app/domain/infrastructure"
	"github.com/murosan/shogi-board-server/app/domain/infrastructure/store"
	"github.com/murosan/shogi-board-server/app/lib/usi/convert"
	"github.com/murosan/shogi-board-server/app/lib/usi/parse"
	"github.com/murosan/shogi-board-server/app/logger"
)

const (
	connectTimeout = time.Second * 5
	closeTimeout   = time.Second * 5
	readyTimeout   = time.Second * 5
)

var (
	namePrefix   = []byte("id name ")
	authorPrefix = []byte("id author ")
	optionPrefix = []byte("option ")
)

// EngineControlService is a service for controlling engine.
type EngineControlService interface {
	Connect() error
	Close() error
	Start() error
	Stop() error
	UpdateButtonOption(*engine.Button) error
	UpdateCheckOption(*engine.Check) error
	UpdateRangeOption(*engine.Range) error
	UpdateSelectOption(*engine.Select) error
	UpdateTextOption(*engine.Text) error
	UpdatePosition(*shogi.Position) error
}

// NewEngineControlService returns new EngineControlService.
func NewEngineControlService(
	engine *engine.Engine,
	connector infrastructure.Connector,
	store store.EngineInfoStore,
	logger logger.Logger,
) EngineControlService {
	return &engineControlService{
		engine:          engine,
		connector:       connector,
		engineInfoStore: store,
		logger:          logger,
	}
}

type engineControlService struct {
	engine          *engine.Engine
	connector       infrastructure.Connector
	engineInfoStore store.EngineInfoStore
	logger          logger.Logger
}

func (service *engineControlService) Connect() error {
	egn := service.engine
	service.logger.Info("[Connecting to Engine]", zap.String("engine id", egn.GetID().String()))

	if err := service.connector.Connect(); err != nil {
		return framework.NewInternalServerError("call connect", err)
	}

	if err := service.write(usi.Command.USI); err != nil {
		return framework.NewInternalServerError("write "+string(usi.Command.USI), err)
	}

	done := make(chan struct{})

	go service.connector.OnReceive(func(b []byte) bool {
		service.logger.Info("[EngineOutput]", zap.ByteString("value", b))
		if bytes.Equal(b, usi.Response.OK) {
			done <- struct{}{}
			return true
		}

		if bytes.Equal(b, usi.Response.ReadyOK) {
			close(done)
			return false
		}

		// set name if s starts with 'id name '
		if bytes.HasPrefix(b, namePrefix) {
			bn := bytes.TrimLeft(b, string(namePrefix))
			sn := string(bytes.TrimSpace(bn))
			egn.SetName(sn)
			service.logger.Info("[EngineName]", zap.String("value", sn))
			return true
		}

		// set author if s starts with 'id author '
		if bytes.HasPrefix(b, authorPrefix) {
			ba := bytes.TrimLeft(b, string(authorPrefix))
			sa := string(bytes.TrimSpace(ba))
			egn.SetAuthor(sa)
			service.logger.Info("[EngineAuthor]", zap.String("value", sa))
			return true
		}

		// parse option
		if bytes.HasPrefix(b, optionPrefix) {
			s := string(b)
			switch {
			case strings.Contains(s, parse.TypeButton):
				opt, err := parse.Button(s)
				if err != nil {
					service.logger.Error("parse button", zap.Error(err))
				}
				service.logger.Info("parsed button", zap.Any("value", opt))
				egn.GetOptions().PutButton(opt)

			case strings.Contains(s, parse.TypeCheck):
				opt, err := parse.Check(s)
				if err != nil {
					service.logger.Error("parse check", zap.Error(err))
				}
				service.logger.Info("parsed check", zap.Any("value", opt))
				egn.GetOptions().PutCheck(opt)

			case strings.Contains(s, parse.TypeRange):
				opt, err := parse.Range(s)
				if err != nil {
					service.logger.Error("parse range", zap.Error(err))
				}
				service.logger.Info("parsed range", zap.Any("value", opt))
				egn.GetOptions().PutRange(opt)

			case strings.Contains(s, parse.TypeSelect):
				opt, err := parse.Select(s)
				if err != nil {
					service.logger.Error("parse select", zap.Error(err))
				}
				service.logger.Info("parsed select", zap.Any("value", opt))
				egn.GetOptions().PutSelect(opt)

			case strings.Contains(s, parse.TypeString):
				opt, err := parse.TextFromStringType(s)
				if err != nil {
					service.logger.Error("parse string", zap.Error(err))
				}
				service.logger.Info("parsed string", zap.Any("value", opt))
				egn.GetOptions().PutText(opt)

			case strings.Contains(s, parse.TypeFilename):
				opt, err := parse.TextFromFilenameType(s)
				if err != nil {
					service.logger.Error("parse filename", zap.Error(err))
				}
				service.logger.Info("parsed filename", zap.Any("value", opt))
				egn.GetOptions().PutText(opt)
			}
		}

		return true
	})

	if err := service.timeout(done, connectTimeout); err != nil {
		close(done)
		return framework.NewInternalServerError("connect timeout. failed to receive usiok", err)
	}

	if err := service.write(usi.Command.IsReady); err != nil {
		return framework.NewInternalServerError("write "+string(usi.Command.IsReady), err)
	}

	if err := service.timeout(done, readyTimeout); err != nil {
		return framework.NewInternalServerError("connect timeout. failed to receive readyok", err)
	}

	egn.SetState(engine.Connected)
	return nil
}

func (service *engineControlService) Close() error {
	egn := service.engine
	service.logger.Info("[Closing Engine]", zap.String("engine name", egn.GetName()))

	if egn.GetState() == engine.NotConnected {
		return nil
	}

	if err := service.write(usi.Command.Quit); err != nil {
		return framework.WrapError("write "+string(usi.Command.Quit), err)
	}

	if err := service.connector.Close(closeTimeout); err != nil {
		return framework.NewInternalServerError("close engine", err)
	}

	return nil
}

func (service *engineControlService) Start() error {
	egn := service.engine
	service.logger.Info("[Starting Engine]", zap.String("engine name", egn.GetName()))

	if egn.GetState() == engine.NotConnected {
		return framework.NewBadRequestError("must initialize engine first", nil)
	}

	if egn.GetState() == engine.Thinking {
		return nil
	}

	if egn.GetState() == engine.Connected {
		if err := service.write(usi.Command.NewGame); err != nil {
			return framework.NewInternalServerError("write "+string(usi.Command.NewGame), err)
		}

		egn.SetState(engine.StandBy)
	}

	// before start thinking, delete all consideration results
	service.engineInfoStore.DeleteAll(egn.GetID())

	done := make(chan struct{})

	// catch call engine outputs on background
	go func() {
		egn := service.engine
		service.connector.OnReceive(func(b []byte) bool {
			service.logger.Info("[EngineOutput]", zap.ByteString("message", b))

			// ignore 'info string' for now
			if bytes.HasPrefix(b, []byte("info string")) {
				return true
			}

			if bytes.HasPrefix(b, []byte("info ")) {
				i, mpv, err := parse.Info(string(b))
				if err != nil {
					service.logger.Error("[start]", zap.Error(err))
					return true // ignore error
				}

				// service.logger.Info("[ParsedInfo]", zap.Any("value", i))

				if mpv <= 1 {
					// If mpv is less than or equal to 1, it means 'best move' usually.
					// We need to delete when the number of candidates is reduced,
					// for example from 5 to 2, not to be left extra information.
					service.engineInfoStore.DeleteAll(egn.GetID())
				}
				if len(i.Moves) != 0 {
					service.engineInfoStore.Upsert(egn.GetID(), mpv, i)
				}
			}

			close(done)
			return egn.GetState() == engine.Thinking
		})
	}()

	if err := service.write(usi.Command.GoInf); err != nil {
		return framework.NewInternalServerError("write "+string(usi.Command.GoInf), nil)
	}

	if err := service.timeout(done, connectTimeout); err != nil {
		return framework.NewInternalServerError("connect timeout", err)
	}
	egn.SetState(engine.Thinking)
	return nil
}

func (service *engineControlService) Stop() error {
	egn := service.engine
	service.logger.Info("[Stopping Engine]", zap.String("engine name", egn.GetName()))

	if egn.GetState() != engine.Thinking {
		return nil
	}

	if err := service.write(usi.Command.Stop); err != nil {
		return framework.WrapError("write "+string(usi.Command.Quit), err)
	}

	egn.SetState(engine.StandBy)
	return nil
}

func (service *engineControlService) UpdateButtonOption(button *engine.Button) error {
	return service.updateOption(button)
}

func (service *engineControlService) UpdateCheckOption(check *engine.Check) error {
	return service.updateOption(check)
}

func (service *engineControlService) UpdateRangeOption(rang *engine.Range) error {
	return service.updateOption(rang)
}

func (service *engineControlService) UpdateSelectOption(sel *engine.Select) error {
	return service.updateOption(sel)
}

func (service *engineControlService) UpdateTextOption(text *engine.Text) error {
	return service.updateOption(text)
}

func (service *engineControlService) updateOption(option engine.Option) error {
	service.logger.Info("[UpdateOption]", zap.String("option", option.String()))
	if err := option.Validate(); err != nil {
		return framework.NewBadRequestError("invalid option value", err)
	}
	return service.write([]byte(option.ToUSI()))
}

func (service *engineControlService) UpdatePosition(position *shogi.Position) error {
	service.logger.Info("[UpdatePosition]", zap.Any("position", position))

	isThinking := service.engine.GetState() == engine.Thinking

	// stop thinking first
	if isThinking {
		if err := service.Stop(); err != nil {
			return err
		}
	}

	b, err := convert.Position(position)
	if err != nil {
		return framework.NewBadRequestError("invalid position", err)
	}
	if err := service.write(b); err != nil {
		return framework.NewInternalServerError("write "+string(b), err)
	}

	// restart thinking
	if isThinking {
		return service.Start()
	}
	return nil
}

func (service *engineControlService) write(bytes []byte) error {
	service.logger.Info("[Write]", zap.ByteString("message", bytes))
	w := service.connector.Writer()
	if _, err := w.Write(append(bytes, '\n')); err != nil {
		return framework.NewInternalServerError("write", err)
	}
	return nil
}

func (service *engineControlService) withTimeout(timeout time.Duration, block func()) error {
	ch := make(chan struct{}, 1)
	go func() {
		block()
		close(ch)
	}()
	return service.timeout(ch, timeout)
}

func (service *engineControlService) timeout(ch chan struct{}, timeout time.Duration) error {
	select {
	case <-ch:
		return nil
	case <-time.After(timeout):
		return framework.NewInternalServerError("timeout", nil)
	}
}
