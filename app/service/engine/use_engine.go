// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package engine

import (
	"github.com/murosan/shogi-proxy-server/app/domain/entity/config"
	"github.com/murosan/shogi-proxy-server/app/domain/infrastracture/engine"
	egn "github.com/murosan/shogi-proxy-server/app/infrastracture/engine"
)

// TODO: 2つのEngineを扱えるように、connection pool とか作って移す
var e engine.Engine = nil

func UseEngine(c config.Config) engine.Engine {
	if e == nil {
		e = egn.NewEngine(c.GetEnginePath())
	}
	return e
}
