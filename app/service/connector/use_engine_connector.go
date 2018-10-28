// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package connector

import (
	"github.com/murosan/shogi-proxy-server/app/domain/infrastracture/connector"
	conn "github.com/murosan/shogi-proxy-server/app/infrastracture/connector"
	"github.com/murosan/shogi-proxy-server/app/service/config"
	"github.com/murosan/shogi-proxy-server/app/service/converter"
)

var c connector.Connector = nil

func UseConnector() connector.Connector {
	if c == nil {
		c = conn.NewConnector(config.UseConfig(), converter.UseFromUsi())
	}
	return c
}
