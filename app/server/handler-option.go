// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/murosan/shogi-proxy-server/app/domain/entity/engine/option"
	"github.com/murosan/shogi-proxy-server/app/domain/entity/exception"
	"github.com/murosan/shogi-proxy-server/app/service/logger"
	"go.uber.org/zap"
)

func (s *server) getOptionList(w http.ResponseWriter, r *http.Request) {
	d, err := json.Marshal(s.conn.GetOptions())
	if err != nil {
		logger.Use().Error("Failed to marshal option list.", zap.Error(err))
		s.internalServerError(w, err)
		return
	}

	logger.Use().Info("GetOptions", zap.ByteString("Marshaled value", d))
	w.Header().Set(contentType, mimeJson)
	w.Write(d)
}

func (s *server) setOption(w http.ResponseWriter, r *http.Request) {
	l, err := strconv.Atoi(r.Header.Get(contentLength))
	if err != nil {
		http.Error(w, err.Error(), 411) // Length Required
		logger.Use().Error("Could not read "+contentLength, zap.Error(err))
		return
	}

	body := make([]byte, l)
	l, err = r.Body.Read(body)
	if err != nil && err != io.EOF {
		m := fmt.Sprintf("%v\ncaused by:\n%v", exception.FailedToReadBody.Error(), err.Error())
		http.Error(w, m, http.StatusInternalServerError)
		logger.Use().Error(m)
		return
	}

	var osv option.OptionSetValue
	if err := json.Unmarshal(body, &osv); err != nil {
		s.internalServerError(w, err)
		return
	}

	s.conn.GetOptions()

	w.WriteHeader(http.StatusOK)
}
