// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package engine

import (
	"os/exec"

	"github.com/murosan/shogi-proxy-server/app/domain/infrastracture/engine"
	egn "github.com/murosan/shogi-proxy-server/app/infrastracture/engine"
	"github.com/murosan/shogi-proxy-server/app/infrastracture/os/command"
	"github.com/murosan/shogi-proxy-server/app/service/config"
)

// TODO: 2つのEngineを扱えるように、connection pool とか作ってこのファイルは消す
var e engine.Engine = nil

func UseEngine(name string) engine.Engine {
	if e == nil {
		ec := exec.Command(config.UseConfig().GetEnginePath(name))
		e = egn.NewEngine(command.NewCmd(ec))
	}
	return e
}
