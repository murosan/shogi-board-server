// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package server

import (
	"encoding/json"
	"net/http"

	"github.com/murosan/shogi-proxy-server/app/domain/entity/engine/option"
	"github.com/murosan/shogi-proxy-server/app/domain/exception"
	"github.com/murosan/shogi-proxy-server/app/domain/infrastracture/engine"
	"go.uber.org/zap"
)

func (s *server) getOptionList(w http.ResponseWriter, r *http.Request) {
	err := s.conn.WithEngine("", func(e engine.Engine) {
		d, err := json.Marshal(e.GetOptions())
		if err != nil {
			s.log.Error("Failed to marshal option list.", zap.Error(err))
			s.internalServerError(w, err)
			return
		}

		s.log.Info("GetOptions", zap.ByteString("Marshaled value", d))
		w.Header().Set(contentType, mimeJson)
		w.Write(d)
	})
	if err != nil {
		s.internalServerError(w, err)
	}
}

func (s *server) updateOption(w http.ResponseWriter, r *http.Request) {
	body, err := s.readJsonBody(r)
	if err != nil && err == exception.ContentLengthRequired {
		http.Error(w, err.Error(), 411) // Length Required
		return
	}
	if err != nil && err == exception.FailedToReadBody {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var osv option.UpdateOptionValue
	if err := json.Unmarshal(body, &osv); err != nil {
		s.internalServerError(w, err)
		return
	}
	s.log.Info("UpdateOptionBody", zap.Any("Unmarshal", osv))

	er := s.conn.WithEngine("", func(e engine.Engine) {
		if err := e.UpdateOption(osv); err != nil {
			// TODO: InternalServerError ではないな・・
			s.internalServerError(w, err)
			return
		}
		w.WriteHeader(http.StatusOK)
	})
	if er != nil {
		s.internalServerError(w, er)
	}
}
