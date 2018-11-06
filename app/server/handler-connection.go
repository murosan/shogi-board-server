// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package server

import (
	"github.com/murosan/shogi-proxy-server/app/service/logger"
	"net/http"
)

func (s *Server) Connect(w http.ResponseWriter, r *http.Request) {
	if err := s.conn.Connect(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Use().Info("Successfully connected.")
	w.WriteHeader(http.StatusCreated)
}

func (s *Server) Close(w http.ResponseWriter, r *http.Request) {
	if err := s.conn.Close(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Use().Info("Successfully closed.")
	w.WriteHeader(http.StatusOK)
}
