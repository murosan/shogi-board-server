// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package server

import (
	"github.com/murosan/shogi-board-server/app/domain/entity/converter"
	"github.com/murosan/shogi-board-server/app/domain/infrastracture/connector"
	"github.com/murosan/shogi-board-server/app/domain/logger"
)

type server struct {
	conn connector.Connector
	fj   converter.FromJSON
	fu   converter.FromUSI
	tu   converter.ToUSI
	log  logger.Log
}

// NewServer 新しいサーバーを返す
func NewServer(
	conn connector.Connector,
	fj converter.FromJSON,
	fu converter.FromUSI,
	tu converter.ToUSI,
	log logger.Log,
) ShogiBoardServer {
	return &server{
		conn,
		fj,
		fu,
		tu,
		log,
	}
}
