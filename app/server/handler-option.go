// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package server

import (
	"encoding/json"
	"net/http"

	"github.com/murosan/shogi-proxy-server/app/service/logger"
	"go.uber.org/zap"
)

func (s *Server) GetOptionList(w http.ResponseWriter, r *http.Request) {
	d, err := json.Marshal(s.conn.GetOptions())
	if err != nil {
		logger.Use().Error("Failed to marshal option list.", zap.Error(err))
		s.internalServerError(w, err)
		return
	}

	logger.Use().Info("GetOptions", zap.ByteString("Marshaled value", d))
	w.Header().Set("Content-Type", "application/json")
	w.Write(d)
}

func (s *Server) SetOption(w http.ResponseWriter, r *http.Request) {
	// TODO
}
