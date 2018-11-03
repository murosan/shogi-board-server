// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package connector

import (
	confModel "github.com/murosan/shogi-proxy-server/app/domain/entity/config"
	connModel "github.com/murosan/shogi-proxy-server/app/domain/infrastracture/connector"
	egnModel "github.com/murosan/shogi-proxy-server/app/domain/infrastracture/engine"
	"github.com/murosan/shogi-proxy-server/app/infrastracture/engine"
	"github.com/murosan/shogi-proxy-server/app/infrastracture/os/command"
	"github.com/murosan/shogi-proxy-server/app/service/config"
	"os/exec"
)

type connectionPool struct {
	conf confModel.Config

	// TODO: 今はとりあえず1つだけ
	e egnModel.Engine
	//em   map[string]egnModel.Engine
}

func NewConnectionPool(c confModel.Config) connModel.ConnectionPool {
	return &connectionPool{c, nil}
}

func (cp *connectionPool) Initialize() {
	// TODO: 今は1つだけ使える
	names := cp.conf.GetEngineNames()
	cmd := exec.Command(config.UseConfig().GetEnginePath(names[0]))
	cp.e = engine.NewEngine(command.NewCmd(cmd))
}

func (cp *connectionPool) NamedEngine() egnModel.Engine {
	return cp.e
}

func (cp *connectionPool) Remove() error {
	cp.e = nil
	return nil
}
