package module

import (
	"github.com/murosan/shogi-board-server/app/domain/config"
	"github.com/murosan/shogi-board-server/app/domain/infrastructure"
	"github.com/murosan/shogi-board-server/app/domain/infrastructure/store"
	"github.com/murosan/shogi-board-server/app/domain/service"
	"github.com/murosan/shogi-board-server/app/logger"
)

var (
	Config *config.Config
	Logger logger.Logger

	Stores = stores{
		Engine:     store.NewEngineStore(),
		EngineInfo: store.NewEngineInfoStore(),
		Game:       store.NewGameStore(),
	}

	Services services
)

type (
	stores struct {
		Engine     store.EngineStore
		EngineInfo store.EngineInfoStore
		Game       store.GameStore
	}

	services struct {
		Engine service.EngineService
	}
)

// Initialize setups module variables.
func Initialize(appConfigPath, logConfigPath string) {
	Config = config.New(appConfigPath, logConfigPath)
	Logger = logger.New(Config)

	Services = services{
		Engine: service.NewEngineService(
			Stores.Engine,
			Stores.EngineInfo,
			Stores.Game,
			Config,
			Logger,
			infrastructure.NewCmd,
			infrastructure.NewConnector,
		),
	}
}
