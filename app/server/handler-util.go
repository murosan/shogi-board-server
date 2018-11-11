// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package server

import (
	"fmt"
	"net/http"

	"github.com/murosan/shogi-proxy-server/app/domain/entity/exception"
	"github.com/murosan/shogi-proxy-server/app/service/logger"
	"go.uber.org/zap"
)

const (
	get  = "GET"
	post = "POST"

	appliJson = "application/json"

	contentType = "Content-Type"
	userAgent   = "User-Agent"
)

func (s *Server) Handling(w http.ResponseWriter, r *http.Request) {
	uri := r.RequestURI
	logger.Use().Info(
		"AccessLog",
		zap.String("uri", uri),
		zap.String("method", r.Method),
		zap.String("addr", r.RemoteAddr),
		zap.String("ua", r.Header.Get(userAgent)),
	)

	switch uri {
	case IndexPath:
		s.withMethod(get, w, r, s.ServeHome)
	case ConnectPath:
		s.withMethod(post, w, r, s.Connect)
	case ClosePath:
		s.withMethod(post, w, r, s.Close)
	case ListOptPath:
		s.withMethod(get, w, r, s.GetOptionList)
	case SetOptPath:
		s.withMethod(post, w, r, s.SetOption)
	case SetPositionPath:
		s.contentTypeCheck(appliJson, w, r, s.SetPosition)
	case StartPath:
		s.withMethod(post, w, r, s.Start)
	case GetValuesPath:
		s.withMethod(get, w, r, s.GetValues)
	case InitAnalyze:
		s.withMethod(post, w, r, s.InitAnalyze)
	case StartAnalyze:
		s.withMethod(post, w, r, s.StartAnalyze)
	default:
		s.notFound(w, r, uri)
	}
}

func (s *Server) contentTypeCheck(tpe string, w http.ResponseWriter, r *http.Request, h http.HandlerFunc) {
	ct := r.Header.Get(contentType)
	if ct != tpe {
		s.badRequest(w, fmt.Sprintf("%s must be %s, but got %s", contentType, tpe, ct))
		return
	}

	h(w, r)
}

func (s *Server) withMethod(meth string, w http.ResponseWriter, r *http.Request, h http.HandlerFunc) {
	if r.Method != meth {
		logger.Use().Debug(fmt.Sprintf("%s, uri: %s\n", exception.MethodNotAllowed, r.RequestURI))
		s.badRequest(w, exception.MethodNotAllowed.Error())
		return
	}

	h(w, r)
}

func (s *Server) internalServerError(w http.ResponseWriter, e error) {
	logger.Use().Error(exception.InternalServerError.Error(), zap.Error(e))
	http.Error(w, e.Error(), http.StatusInternalServerError)
}

func (s *Server) badRequest(w http.ResponseWriter, m string) {
	logger.Use().Debug(m)
	http.Error(w, m, http.StatusBadRequest)
}

func (s *Server) notFound(w http.ResponseWriter, r *http.Request, uri string) {
	logger.Use().Debug(exception.NotFound.Error(), zap.String("uri", uri))
	http.NotFound(w, r)
}
