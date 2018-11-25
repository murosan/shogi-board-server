// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package server

import (
	"github.com/murosan/shogi-proxy-server/app/domain/entity/engine/state"
	"github.com/murosan/shogi-proxy-server/app/domain/entity/exception"
	"github.com/murosan/shogi-proxy-server/app/domain/entity/usi"
	"github.com/murosan/shogi-proxy-server/app/domain/infrastracture/engine"
	"github.com/murosan/shogi-proxy-server/app/service/logger"
	"net/http"
)

// Content-Type は application/json である必要がある
func (s *server) setPosition(w http.ResponseWriter, r *http.Request) {
	body, err := s.readJsonBody(r)
	if err != nil && err == exception.ContentLengthRequired {
		http.Error(w, err.Error(), 411) // Length Required
		return
	}
	if err != nil && err == exception.FailedToReadBody {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pos, err := s.fj.Position(body)
	if err != nil {
		s.internalServerError(w, err)
		return
	}

	setPosUsi, err := s.tu.Position(pos)
	if err != nil {
		s.internalServerError(w, err)
		return
	}

	s.conn.WithEngine("", func(e engine.Engine) {
		if e == nil || e.GetState() == state.NotConnected {
			// TODO: internal server error ではないな
			s.internalServerError(w, exception.EngineIsNotRunning)
			return
		}
		isThinking := e.GetState() == state.Thinking
		// 思考中なら stop
		if isThinking {
			if err := e.Exec(usi.CmdStop); err != nil {
				s.internalServerError(w, err)
				return
			}
			e.SetState(state.StandBy)
		}
		if err := e.Exec(setPosUsi); err != nil {
			s.internalServerError(w, err)
			return
		}
		// さっき思考中だったら再スタート
		if isThinking {
			s.start(w, r)
			return
		}
		w.WriteHeader(http.StatusOK)
	})
}

func (s *server) start(w http.ResponseWriter, r *http.Request) {
	s.conn.WithEngine("", func(e engine.Engine) {
		if e == nil || e.GetState() == state.NotConnected {
			// TODO: internal server error ではないな
			s.internalServerError(w, exception.EngineIsNotRunning)
			return
		}

		switch e.GetState() {
		case state.Connected: // usinewgame 前なら実行
			if err := e.Exec(usi.CmdNewGame); err != nil {
				s.internalServerError(w, err)
				return
			}
		case state.StandBy: // 思考開始
			if err := e.Exec(usi.CmdGoInf); err != nil {
				s.internalServerError(w, err)
				return
			}
		case state.Thinking:
			logger.Use().Debug("Engine is thinking. Nothing to do.")
		default:
			panic("unknown engine state.")
		}

		w.WriteHeader(http.StatusOK)
	})
}

func (s *server) getValues(w http.ResponseWriter, r *http.Request) {
	// TODO
}
