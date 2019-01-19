// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package connector

import (
	"os/exec"

	confModel "github.com/murosan/shogi-board-server/app/domain/config"
	connModel "github.com/murosan/shogi-board-server/app/domain/infrastracture/connector"
	egnModel "github.com/murosan/shogi-board-server/app/domain/infrastracture/engine"
	"github.com/murosan/shogi-board-server/app/domain/logger"
	"github.com/murosan/shogi-board-server/app/infrastracture/engine"
	"github.com/murosan/shogi-board-server/app/infrastracture/os/command"
	"github.com/murosan/shogi-board-server/app/service/config"
)

type connectionPool struct {
	conf confModel.Config

	// TODO: 今はとりあえず1つだけ
	e egnModel.Engine
	//em   map[string]egnModel.Engine

	log logger.Log
}

// NewConnectionPool 新しい ConnectionPool を返す
func NewConnectionPool(c confModel.Config, log logger.Log) connModel.ConnectionPool {
	return &connectionPool{c, nil, log}
}

func (cp *connectionPool) Initialize() {
	// TODO: 今は1つだけ使える
	names := cp.conf.GetEngineNames()
	cmd := exec.Command(config.UseConfig().GetEnginePath(names[0]))
	cp.e = engine.NewEngine(command.NewCmd(cmd), cp.log)
}

func (cp *connectionPool) NamedEngine() egnModel.Engine {
	return cp.e
}

func (cp *connectionPool) Remove() error {
	cp.e = nil
	return nil
}
