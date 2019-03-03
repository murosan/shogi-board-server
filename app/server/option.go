// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package server

import (
	"context"

	"github.com/murosan/shogi-board-server/app/domain/entity/engine"
	"github.com/murosan/shogi-board-server/app/domain/exception"
	pb "github.com/murosan/shogi-board-server/app/proto"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetOptions は将棋エンジンのオプション一覧を返す
func (s *Server) GetOptions(ctx context.Context, in *pb.EngineName) (*pb.Options, error) {
	s.accessLog("GetOptions")
	egn, err := s.conn.GetEngine(in.Name)

	if err != nil {
		msg := exception.FailedToConnect.WithMsg(err.Error()).Error()
		return nil, status.Error(codes.Unknown, msg)
	}

	opts := egn.GetOptions()
	s.log.Info("GetOptions", zap.Any("got options", opts))
	return opts, nil
}

// TODO: コピペの嵐。うまく抽象化する方法がわからん。誰か教えて

// UpdateButton は Options の Button の値を更新し、
// 将棋エンジンに setoption USIコマンドを実行します
func (s *Server) UpdateButton(ctx context.Context, in *pb.UpdateButtonRequest) (*pb.Response, error) {
	s.accessLog("UpdateButton")
	egn, err := s.conn.GetEngine(in.Engine.Name)

	if err != nil {
		msg := exception.FailedToConnect.WithMsg(err.Error()).Error()
		return nil, status.Error(codes.Unknown, msg)
	}

	usi := engine.ButtonUSI(in.Button)
	if err := egn.Exec([]byte(usi)); err != nil {
		msg := exception.FailedToUpdateOption.WithMsg(err.Error()).Error()
		return nil, status.Error(codes.Unknown, msg)
	}

	return pb.NewResponse(), nil
}

// UpdateCheck は Options の Check の値を更新し、
// 将棋エンジンに setoption USIコマンドを実行します
func (s *Server) UpdateCheck(ctx context.Context, in *pb.UpdateCheckRequest) (*pb.Response, error) {
	s.accessLog("UpdateCheck")
	egn, err := s.conn.GetEngine(in.Engine.Name)

	if err != nil {
		msg := exception.FailedToConnect.WithMsg(err.Error()).Error()
		return nil, status.Error(codes.Unknown, msg)
	}

	opts := egn.GetOptions()
	engine.UpdateCheck(opts, in.Check)

	usi := engine.CheckUSI(in.Check)
	if err := egn.Exec([]byte(usi)); err != nil {
		msg := exception.FailedToUpdateOption.WithMsg(err.Error()).Error()
		return nil, status.Error(codes.Unknown, msg)
	}

	s.log.Info("Successfully updated")
	return pb.NewResponse(), nil
}

// UpdateSpin は Options の Spin の値を更新し、
// 将棋エンジンに setoption USIコマンドを実行します
func (s *Server) UpdateSpin(ctx context.Context, in *pb.UpdateSpinRequest) (*pb.Response, error) {
	s.accessLog("UpdateSpin")
	egn, err := s.conn.GetEngine(in.Engine.Name)

	if err != nil {
		msg := exception.FailedToConnect.WithMsg(err.Error()).Error()
		return nil, status.Error(codes.Unknown, msg)
	}

	opts := egn.GetOptions()
	engine.UpdateSpin(opts, in.Spin)

	usi := engine.SpinUSI(in.Spin)
	if err := egn.Exec([]byte(usi)); err != nil {
		msg := exception.FailedToUpdateOption.WithMsg(err.Error()).Error()
		return nil, status.Error(codes.Unknown, msg)
	}

	return pb.NewResponse(), nil
}

// UpdateSelect は Options の Select の値を更新し、
// 将棋エンジンに setoption USIコマンドを実行します
func (s *Server) UpdateSelect(ctx context.Context, in *pb.UpdateSelectRequest) (*pb.Response, error) {
	s.accessLog("UpdateSelect")
	egn, err := s.conn.GetEngine(in.Engine.Name)

	if err != nil {
		msg := exception.FailedToConnect.WithMsg(err.Error()).Error()
		return nil, status.Error(codes.Unknown, msg)
	}

	opts := egn.GetOptions()
	engine.UpdateSelect(opts, in.Select)

	usi := engine.SelectUSI(in.Select)
	if err := egn.Exec([]byte(usi)); err != nil {
		msg := exception.FailedToUpdateOption.WithMsg(err.Error()).Error()
		return nil, status.Error(codes.Unknown, msg)
	}

	return pb.NewResponse(), nil
}

// UpdateString は Options の String の値を更新し、
// 将棋エンジンに setoption USIコマンドを実行します
func (s *Server) UpdateString(ctx context.Context, in *pb.UpdateStringRequest) (*pb.Response, error) {
	s.accessLog("UpdateString")
	egn, err := s.conn.GetEngine(in.Engine.Name)

	if err != nil {
		msg := exception.FailedToConnect.WithMsg(err.Error()).Error()
		return nil, status.Error(codes.Unknown, msg)
	}

	opts := egn.GetOptions()
	engine.UpdateString(opts, in.String_)

	usi := engine.StringUSI(in.String_)
	if err := egn.Exec([]byte(usi)); err != nil {
		msg := exception.FailedToUpdateOption.WithMsg(err.Error()).Error()
		return nil, status.Error(codes.Unknown, msg)
	}

	return pb.NewResponse(), nil
}

// UpdateFilename は Options の Filename の値を更新し、
// 将棋エンジンに setoption USIコマンドを実行します
func (s *Server) UpdateFilename(ctx context.Context, in *pb.UpdateFilenameRequest) (*pb.Response, error) {
	s.accessLog("UpdateFilename")
	egn, err := s.conn.GetEngine(in.Engine.Name)

	if err != nil {
		msg := exception.FailedToConnect.WithMsg(err.Error()).Error()
		return nil, status.Error(codes.Unknown, msg)
	}

	opts := egn.GetOptions()
	engine.UpdateFilename(opts, in.Filename)

	usi := engine.FilenameUSI(in.Filename)
	if err := egn.Exec([]byte(usi)); err != nil {
		msg := exception.FailedToUpdateOption.WithMsg(err.Error()).Error()
		return nil, status.Error(codes.Unknown, msg)
	}

	return pb.NewResponse(), nil
}
