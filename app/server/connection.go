// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package server

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/murosan/shogi-board-server/app/domain/exception"
	pb "github.com/murosan/shogi-board-server/app/proto"
)

// Initialize は app.yml で設定された接続可能な将棋エンジンの名前一覧を返す
func (s *Server) Initialize(ctx context.Context, in *pb.Request) (*pb.EngineNames, error) {
	s.accessLog("GetEngineNames")
	n := s.conn.GetEngineNames()
	s.log.Info("GetEngineNames", zap.Strings("result", n))
	return pb.NewEngineNames(n), nil
}

// Connect 指定の将棋エンジンに接続する
func (s *Server) Connect(ctx context.Context, in *pb.EngineName) (*pb.Response, error) {
	s.accessLog("Connect")
	if err := s.conn.Connect(); err != nil {
		msg := exception.FailedToConnect.WithMsg(err.Error()).Error()
		return nil, status.Error(codes.Unknown, msg)
	}

	s.log.Info("Successfully connected")
	return pb.NewResponse(), nil
}

// Close 指定の将棋エンジンとの接続を切る
func (s *Server) Close(ctx context.Context, in *pb.EngineName) (*pb.Response, error) {
	s.accessLog("Close")
	if err := s.conn.Close(); err != nil {
		msg := exception.FailedToClose.WithMsg(err.Error()).Error()
		return nil, status.Error(codes.Unknown, msg)
	}

	s.log.Info("Successfully closed")
	return pb.NewResponse(), nil
}
