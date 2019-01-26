// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package server

import (
	"github.com/murosan/shogi-board-server/app/domain/entity/converter"
	"github.com/murosan/shogi-board-server/app/domain/infrastracture/connector"
	"github.com/murosan/shogi-board-server/app/domain/logger"
	"go.uber.org/zap"
)

// Server は protocol buffers で定義された要件を満たすサーバー
type Server struct {
	conn connector.Connector
	fu   converter.FromUSI
	tu   converter.ToUSI
	log  logger.Log
}

// NewServer returns new protocol buffer server
func NewServer(
	conn connector.Connector,
	fu converter.FromUSI,
	tu converter.ToUSI,
	log logger.Log,
) *Server {
	return &Server{
		conn,
		fu,
		tu,
		log,
	}
}

func (s *Server) accessLog(rpcName string) {
	s.log.Info("access log", zap.String("rpc", rpcName))
}
