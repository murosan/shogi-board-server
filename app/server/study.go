// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package server

import (
	"bytes"
	"context"

	"github.com/murosan/shogi-board-server/app/domain/entity/engine"
	"github.com/murosan/shogi-board-server/app/domain/entity/usi"
	"github.com/murosan/shogi-board-server/app/domain/exception"
	eg "github.com/murosan/shogi-board-server/app/domain/infrastracture/engine"
	pb "github.com/murosan/shogi-board-server/app/proto"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// SetPosition 指定の将棋エンジンに setposition USI コマンド を実行する
func (s *Server) SetPosition(ctx context.Context, in *pb.SetPositionRequest) (*pb.Response, error) {
	s.accessLog("SetPosition")
	egn, err := s.conn.GetEngine(in.Engine.Name)

	if err != nil {
		msg := exception.FailedToConnect.WithMsg(err.Error()).Error()
		return nil, status.Error(codes.Unknown, msg)
	}

	usiSetPos, err := s.tu.Position(in.Pos)
	if err != nil {
		msg := exception.FailedToConvert.WithMsg(err.Error()).Error()
		return nil, status.Error(codes.Unknown, msg)
	}

	// 思考中なら stop
	// bestmove受け取ったかなど知らん
	isThinking := egn.GetState() == engine.Thinking
	if isThinking {
		if err := egn.Exec(usi.CmdStop); err != nil {
			msg := exception.FailedToStop.WithMsg(err.Error()).Error()
			return nil, status.Error(codes.Unknown, msg)
		}
		egn.SetState(engine.StandBy)
	}

	if err := egn.Exec([]byte(usiSetPos)); err != nil {
		msg := exception.FailedToExecUSI.WithMsg(err.Error()).Error()
		return nil, status.Error(codes.Unknown, msg)
	}

	egn.FlushResult()

	// さっき思考中だったら再スタート
	if isThinking {
		if err := s.start(egn); err != nil {
			return nil, status.Error(codes.Unknown, err.Error())
		}
	}

	return pb.NewResponse(), nil
}

// Start 将棋エンジンに思考を開始させる
// go inf USI コマンドが実行される
func (s *Server) Start(ctx context.Context, in *pb.EngineName) (*pb.Response, error) {
	s.accessLog("Start")
	egn, err := s.conn.GetEngine(in.Name)

	if err != nil {
		msg := exception.FailedToConnect.WithMsg(err.Error()).Error()
		return nil, status.Error(codes.Unknown, msg)
	}

	if err := s.start(egn); err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}
	return pb.NewResponse(), nil
}

// Stop 将棋エンジンの思考を停止する
func (s *Server) Stop(ctx context.Context, in *pb.EngineName) (*pb.Response, error) {
	s.accessLog("Stop")
	egn, err := s.conn.GetEngine(in.Name)

	if err != nil {
		msg := exception.FailedToConnect.WithMsg(err.Error()).Error()
		return nil, status.Error(codes.Unknown, msg)
	}

	if err := egn.Exec(usi.CmdStop); err != nil {
		s.log.Error("failed to stop", zap.Error(err))
		msg := exception.FailedToStop.WithMsg(err.Error()).Error()
		return nil, status.Error(codes.Unknown, msg)
	}

	egn.SetState(engine.StandBy)
	egn.FlushResult()
	return pb.NewResponse(), nil
}

// GetResult 将棋エンジンの思考結果を取得する
func (s *Server) GetResult(ctx context.Context, in *pb.EngineName) (*pb.Result, error) {
	s.accessLog("GetResult")
	egn, err := s.conn.GetEngine(in.Name)

	if err != nil {
		msg := exception.FailedToConnect.WithMsg(err.Error()).Error()
		return nil, status.Error(codes.Unknown, msg)
	}

	s.log.Info("got result", zap.Any("result", egn.GetResult()))
	return egn.GetResult(), nil
}

func (s *Server) start(egn eg.Engine) error {
	state := egn.GetState()
	if state == engine.Thinking {
		s.log.Info("Engine is thinking. Nothing to do.")
		return nil
	}

	// usinewgame 前なら usinewgame
	if state == engine.Connected {
		if err := egn.Exec(usi.CmdNewGame); err != nil {
			return exception.FailedToStop.WithMsg(err.Error())
		}
		egn.SetState(engine.StandBy)
	}

	// 思考開始
	if err := egn.Exec(usi.CmdGoInf); err != nil {
		return exception.FailedToStart.WithMsg(err.Error())
	}

	// 出力を受け取り続けて、engine の Result に追加していく
	go func() {
		for b := range egn.GetChan() {
			s.log.Info("engine output", zap.ByteString("receive", b))
			if bytes.HasPrefix(b, []byte("info string")) {
				continue // info string は無視
			}
			if bytes.HasPrefix(b, []byte("info ")) {
				i, mpv, err := s.fu.Info(string(b))
				if err != nil {
					s.log.Error("ParseInfoError", zap.Error(err))
					continue
				}
				if len(i.Moves) > 0 {
					egn.SetResult(i, mpv)
				}
			}
		}
	}()

	egn.SetState(engine.Thinking)
	return nil
}
