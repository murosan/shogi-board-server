// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package connector

import (
	connModel "github.com/murosan/shogi-proxy-server/app/domain/infrastracture/connector"
	"github.com/murosan/shogi-proxy-server/app/infrastracture/connector"
	"github.com/murosan/shogi-proxy-server/app/service/config"
	"github.com/murosan/shogi-proxy-server/app/service/converter"
)

var c connModel.Connector = nil

func UseConnector() connModel.Connector {
	if c == nil {
		p := connector.NewConnectionPool(config.UseConfig())
		p.Initialize()
		c = connector.NewConnector(config.UseConfig(), p, converter.UseFromUsi())
	}
	return c
}
