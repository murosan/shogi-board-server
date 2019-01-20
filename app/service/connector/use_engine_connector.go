// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package connector

import (
	connModel "github.com/murosan/shogi-board-server/app/domain/infrastracture/connector"
	"github.com/murosan/shogi-board-server/app/infrastracture/connector"
	"github.com/murosan/shogi-board-server/app/service/config"
	"github.com/murosan/shogi-board-server/app/service/converter"
	"github.com/murosan/shogi-board-server/app/service/logger"
)

var initialized = false
var c connModel.Connector

// UseConnector シングルトンで保持している Connector を返す
func UseConnector() connModel.Connector {
	if !initialized {
		p := connector.NewConnectionPool(config.UseConfig(), logger.Use())
		c = connector.NewConnector(
			config.UseConfig(),
			p,
			converter.UseFromUSI(),
			logger.Use(),
		)
		initialized = true
	}
	return c
}
