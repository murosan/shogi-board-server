package service

import (
	"github.com/murosan/shogi-board-server/app/domain/config"
	"github.com/murosan/shogi-board-server/app/domain/entity/engine"
	"github.com/murosan/shogi-board-server/app/domain/entity/shogi"
	"github.com/murosan/shogi-board-server/app/domain/entity/usi"
	"github.com/murosan/shogi-board-server/app/domain/framework"
	"github.com/murosan/shogi-board-server/app/domain/infrastructure"
	"github.com/murosan/shogi-board-server/app/domain/infrastructure/store"
	"github.com/murosan/shogi-board-server/app/logger"
)

// EngineService is a service for engine.
// This service controls engine store, and delegates
// actual engine controlling task to EngineControlService.
type EngineService interface {
	Connect(engine.ID) error
	Close(engine.ID) error
	CloseAll() error
	Start(engine.ID) error
	Stop(engine.ID) error
	GetOptions(engine.ID) (*engine.Options, error)
	UpdateButtonOption(engine.ID, *engine.Button) error
	UpdateCheckOption(engine.ID, *engine.Check) error
	UpdateRangeOption(engine.ID, *engine.Range) error
	UpdateSelectOption(engine.ID, *engine.Select) error
	UpdateTextOption(engine.ID, *engine.Text) error
	GetCurrentPosition(engine.ID) (*shogi.Position, bool)
	UpdatePosition(engine.ID, *shogi.Position) error
	GetResult(engine.ID) usi.Result
}

// NewEngineService returns new EngineService.
func NewEngineService(
	engineStore store.EngineStore,
	engineInfoStore store.EngineInfoStore,
	gameStore store.GameStore,
	config *config.Config,
	logger logger.Logger,
	newCmd func(string) infrastructure.Cmd,
	newConnector func(infrastructure.Cmd, logger.Logger) infrastructure.Connector,
) EngineService {
	return &engineService{
		engineStore:     engineStore,
		engineInfoStore: engineInfoStore,
		gameStore:       gameStore,
		config:          config,
		logger:          logger,
		newCmd:          newCmd,
		newConnector:    newConnector,
	}
}

type engineService struct {
	engineStore     store.EngineStore
	engineInfoStore store.EngineInfoStore
	gameStore       store.GameStore

	config *config.Config
	logger logger.Logger

	newCmd       func(string) infrastructure.Cmd
	newConnector func(infrastructure.Cmd, logger.Logger) infrastructure.Connector
}

func (service *engineService) Connect(id engine.ID) error {
	if service.engineStore.Exists(id) {
		return framework.NewBadRequestError("engine already exists. id="+id.String(), nil)
	}

	// TODO
	path, ok := service.config.App.Engines[id.String()]
	if !ok {
		return framework.NewNotFoundError(
			"engine path not found. "+
				"please specify in config. id="+id.String(),
			nil,
		)
	}

	egn := engine.New(id, path)
	cmd := service.newCmd(path)
	conn := service.newConnector(cmd, service.logger)
	if err := service.engineStore.Insert(egn, conn); err != nil {
		return framework.NewInternalServerError("insert new engine", err)
	}

	return service.withControl(id, func(service EngineControlService) error {
		return service.Connect()
	})
}

func (service *engineService) Close(id engine.ID) error {
	return service.withControl(id, func(ecs EngineControlService) error {
		if err := ecs.Stop(); err != nil {
			return err
		}
		if err := ecs.Close(); err != nil {
			return err
		}
		return service.engineStore.Delete(id)
	})
}

func (service *engineService) CloseAll() error {
	for _, id := range service.engineStore.FindAllKeys() {
		if err := service.Close(id); err != nil {
			return framework.NewInternalServerError("delete engine at engine service. ID="+id.String(), err)
		}
	}
	return nil
}

func (service *engineService) Start(id engine.ID) error {
	return service.withControl(id, func(service EngineControlService) error {
		return service.Start()
	})
}

func (service *engineService) Stop(id engine.ID) error {
	return service.withControl(id, func(service EngineControlService) error {
		return service.Stop()
	})
}

func (service *engineService) GetOptions(id engine.ID) (*engine.Options, error) {
	egn, _, ok := service.engineStore.Find(id)
	if !ok {
		return nil, framework.NewNotFoundError("no such engine. ID="+id.String(), nil)
	}
	return egn.GetOptions(), nil
}

func (service *engineService) UpdateButtonOption(id engine.ID, button *engine.Button) error {
	return service.withControl(id, func(service EngineControlService) error {
		return service.UpdateButtonOption(button)
	})
}

func (service *engineService) UpdateCheckOption(id engine.ID, check *engine.Check) error {
	return service.withControl(id, func(service EngineControlService) error {
		return service.UpdateCheckOption(check)
	})
}

func (service *engineService) UpdateRangeOption(id engine.ID, rng *engine.Range) error {
	return service.withControl(id, func(service EngineControlService) error {
		return service.UpdateRangeOption(rng)
	})
}

func (service *engineService) UpdateSelectOption(id engine.ID, sel *engine.Select) error {
	return service.withControl(id, func(service EngineControlService) error {
		return service.UpdateSelectOption(sel)
	})
}

func (service *engineService) UpdateTextOption(id engine.ID, txt *engine.Text) error {
	return service.withControl(id, func(service EngineControlService) error {
		return service.UpdateTextOption(txt)
	})
}

func (service *engineService) GetCurrentPosition(id engine.ID) (*shogi.Position, bool) {
	return service.gameStore.FindPosition(id)
}

func (service *engineService) UpdatePosition(id engine.ID, pos *shogi.Position) error {
	return service.withControl(id, func(ecs EngineControlService) error {
		if err := ecs.UpdatePosition(pos); err != nil {
			return err
		}
		service.gameStore.UpsertPosition(id, pos)
		return nil
	})
}

func (service *engineService) GetResult(id engine.ID) usi.Result {
	return service.engineInfoStore.FindAll(id)
}

func (service *engineService) withControl(
	id engine.ID,
	block func(EngineControlService) error,
) error {
	egn, conn, ok := service.engineStore.Find(id)
	if !ok {
		return framework.NewNotFoundError("no such engine. ID="+id.String(), nil)
	}

	s := NewEngineControlService(egn, conn, service.engineInfoStore, service.logger)
	return block(s)
}
