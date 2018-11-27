// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package server

import (
	"github.com/murosan/shogi-proxy-server/app/domain/entity/converter/from_json"
	"github.com/murosan/shogi-proxy-server/app/domain/entity/converter/to_usi"
	"github.com/murosan/shogi-proxy-server/app/domain/entity/logger"
	"github.com/murosan/shogi-proxy-server/app/domain/infrastracture/connector"
)

// TODO: Usecase にちゃんと実装し直す
type server struct {
	conn connector.Connector
	fj   *from_json.FromJson
	tu   *to_usi.ToUsi
	log  logger.Log
}

// TODO: interface の方返すかどうか微妙だ・・
func NewServer(
	conn connector.Connector,
	fj *from_json.FromJson, tu *to_usi.ToUsi,
	log logger.Log,
) ShogiProxyServer {
	return &server{
		conn,
		fj,
		tu,
		log,
	}
}
