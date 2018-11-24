// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package server

import (
	"net/http"

	"github.com/murosan/shogi-proxy-server/app/domain/entity/exception"
	"github.com/murosan/shogi-proxy-server/app/service/logger"
)

func (s *server) connect(w http.ResponseWriter, r *http.Request) {
	if err := s.conn.Connect(); err != nil {
		http.Error(w, exception.FailedToConnect.WithMsg(err.Error()).Error(), http.StatusInternalServerError)
		return
	}

	logger.Use().Info("Successfully connected.")
	w.WriteHeader(http.StatusCreated)
}

func (s *server) close(w http.ResponseWriter, r *http.Request) {
	if err := s.conn.Close(); err != nil {
		http.Error(w, exception.FailedToClose.WithMsg(err.Error()).Error(), http.StatusInternalServerError)
		return
	}

	logger.Use().Info("Successfully closed.")
	w.WriteHeader(http.StatusOK)
}
