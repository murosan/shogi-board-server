// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package connector

import (
	"github.com/murosan/shogi-proxy-server/app/domain/entity/config"
	connModel "github.com/murosan/shogi-proxy-server/app/domain/infrastracture/connector"
	egnModel "github.com/murosan/shogi-proxy-server/app/domain/infrastracture/engine"
	"github.com/murosan/shogi-proxy-server/app/service/engine"
)

type connectionPool struct {
	conf config.Config

	// TODO: 今はとりあえず1つだけ
	e egnModel.Engine
	//em   map[string]egnModel.Engine
}

func NewConnectionPool(c config.Config) connModel.ConnectionPool {
	return &connectionPool{c, nil}
}

func (cp *connectionPool) Initialize() {
	// TODO: 今は1つだけ使える
	names := cp.conf.GetEngineNames()
	cp.e = engine.UseEngine(names[0])
}

func (cp *connectionPool) NamedEngine() egnModel.Engine {
	return cp.e
}
