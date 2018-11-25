// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package server

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/murosan/shogi-proxy-server/app/domain/entity/exception"
	"github.com/murosan/shogi-proxy-server/app/domain/entity/usi"
	"github.com/murosan/shogi-proxy-server/app/service/logger"
	"go.uber.org/zap"
)

// Content-Type は application/json である必要がある
func (s *server) setPosition(w http.ResponseWriter, r *http.Request) {
	l, err := strconv.Atoi(r.Header.Get(contentLength))
	if err != nil {
		http.Error(w, err.Error(), 411) // Length Required
		logger.Use().Error("Could not read "+contentLength, zap.Error(err))
		return
	}

	// read body
	body := make([]byte, l)
	l, err = r.Body.Read(body)
	if err != nil && err != io.EOF {
		m := fmt.Sprintf("%v\ncaused by:\n%v", exception.FailedToReadBody.Error(), err.Error())
		http.Error(w, m, http.StatusInternalServerError)
		logger.Use().Error(m)
		return
	}

	pos, err := s.fj.Position(body)
	if err != nil {
		s.internalServerError(w, err)
		return
	}

	usiCmd, err := s.tu.Position(pos)
	if err != nil {
		s.internalServerError(w, err)
		return
	}

	if err := s.conn.Exec(usiCmd); err != nil {
		s.internalServerError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (s *server) start(w http.ResponseWriter, r *http.Request) {
	if err := s.conn.Exec(usi.CmdNewGame); err != nil {
		s.internalServerError(w, err)
		return
	}
	if err := s.conn.Start(); err != nil {
		s.internalServerError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (s *server) getValues(w http.ResponseWriter, r *http.Request) {
	// TODO
}
