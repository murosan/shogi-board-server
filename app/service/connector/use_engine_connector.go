// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package connector

import (
	"github.com/murosan/shogi-proxy-server/app/domain/entity/config"
	"github.com/murosan/shogi-proxy-server/app/domain/infrastracture/connector"
	conn "github.com/murosan/shogi-proxy-server/app/infrastracture/connector"
)

var c connector.Connector = nil

func UseConnector(conf config.Config) connector.Connector {
	if c == nil {
		c = conn.NewConnector(conf)
	}
	return c
}
