// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package server

import (
	"net/http"

	"github.com/murosan/shogi-proxy-server/app/domain/entity/engine/state"
	"github.com/murosan/shogi-proxy-server/app/domain/entity/usi"
	"github.com/murosan/shogi-proxy-server/app/domain/exception"
	"github.com/murosan/shogi-proxy-server/app/domain/infrastracture/engine"
	"go.uber.org/zap"
)

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

	er := s.conn.WithEngine("", func(e engine.Engine) {
		isThinking := e.GetState() == state.Thinking
		// 思考中なら stop
		// TODO: bestmove受け取ったかなど知らん
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

		e.FlushResult()

		// さっき思考中だったら再スタート
		if isThinking {
			s.start(w, r)
			return
		}
		w.WriteHeader(http.StatusOK)
	})

	if er != nil {
		s.internalServerError(w, er)
	}
}

func (s *server) start(w http.ResponseWriter, r *http.Request) {
	err := s.conn.WithEngine("", func(e engine.Engine) {
		stt := e.GetState()
		if stt == state.Thinking {
			s.log.Debug("Engine is thinking. Nothing to do.")
			w.WriteHeader(http.StatusOK)
			return
		}

		// usinewgame 前なら usinewgame
		if stt == state.Connected {
			if err := e.Exec(usi.CmdNewGame); err != nil {
				s.internalServerError(w, err)
				return
			}
			e.SetState(state.StandBy)
		}

		// 思考開始
		if stt == state.StandBy {
			if err := e.Exec(usi.CmdGoInf); err != nil {
				s.internalServerError(w, err)
				return
			}
			e.SetState(state.Thinking)
		}

		w.WriteHeader(http.StatusOK)
	})

	if err != nil {
		s.internalServerError(w, err)
	}
}

func (s *server) stop(w http.ResponseWriter, r *http.Request) {
	err := s.conn.WithEngine("", func(e engine.Engine) {
		// 細かいことは後から考える。とりあえず stop してしまう
		if err := e.Exec(usi.CmdStop); err != nil {
			s.log.Error("Failed to exec stop", zap.Error(err))
			s.internalServerError(w, err)
			return
		}
		e.FlushResult()
		w.WriteHeader(http.StatusOK)
		return
	})
	if err != nil {
		s.internalServerError(w, err)
	}
}

func (s *server) getValues(w http.ResponseWriter, r *http.Request) {
	// TODO
}
