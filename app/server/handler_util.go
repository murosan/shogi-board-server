// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package server

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/murosan/shogi-board-server/app/domain/exception"
	"go.uber.org/zap"
)

const (
	get  = "GET"
	post = "POST"

	mimeJSON = "application/json"

	contentType   = "Content-Type"
	contentLength = "Content-Length"
	userAgent     = "User-Agent"
)

func (s *server) Handling(w http.ResponseWriter, r *http.Request) {
	uri := r.RequestURI
	s.log.Info(
		"AccessLog",
		zap.String("uri", uri),
		zap.String("method", r.Method),
		zap.String("addr", r.RemoteAddr),
		zap.String("ua", r.Header.Get(userAgent)),
	)

	switch uri {
	case IndexPath:
		s.withMethod(get, w, r, s.serveHome)
	case ConnectPath:
		s.withMethod(post, w, r, s.connect)
	case ClosePath:
		s.withMethod(post, w, r, s.close)
	case ListOptPath:
		s.withMethod(get, w, r, s.getOptionList)
	case SetOptPath:
		s.withMethod(post, w, r, s.contentTypeCheck(mimeJSON, s.updateOption))
	case SetPositionPath:
		s.withMethod(post, w, r, s.contentTypeCheck(mimeJSON, s.setPosition))
	case StartPath:
		s.withMethod(post, w, r, s.start)
	case StopPath:
		s.withMethod(post, w, r, s.stop)
	case GetResultPath:
		s.withMethod(get, w, r, s.getResult)
	case InitAnalyze:
		s.withMethod(post, w, r, s.initAnalyze)
	case StartAnalyze:
		s.withMethod(post, w, r, s.startAnalyze)
	default:
		s.notFound(w, r, uri)
	}
}

func (s *server) contentTypeCheck(
	tpe string,
	h http.HandlerFunc,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ct := r.Header.Get(contentType)
		if ct != tpe {
			s.badRequest(w, fmt.Sprintf(
				"%s must be %s, but got %s",
				contentType,
				tpe,
				ct,
			))
			return
		}

		h(w, r)
	}
}

func (s *server) withMethod(
	meth string,
	w http.ResponseWriter,
	r *http.Request,
	h http.HandlerFunc,
) {
	if r.Method != meth {
		s.log.Debug(
			fmt.Sprintf(
				"%s, uri: %s\n",
				exception.MethodNotAllowed,
				r.RequestURI))
		s.badRequest(w, exception.MethodNotAllowed.Error())
		return
	}

	h(w, r)
}

// JSONのbodyを読み取ってバイト配列にして返す
// error1: exception.ContentLengthRequired
//         HttpHeaderにContent-Lengthがなかったり、int にできなかったエラー
// error2: exception.FailedToReadBody
//         body の読み取りに失敗したエラー
func (s *server) readJSONBody(r *http.Request) ([]byte, error) {
	l, err := strconv.Atoi(r.Header.Get(contentLength))
	if err != nil {
		s.log.Error("Could not read "+contentLength, zap.Error(err))
		return nil, exception.ContentLengthRequired
	}

	body := make([]byte, l)
	_, err = r.Body.Read(body)
	if err != nil && err != io.EOF {
		m := fmt.Sprintf(
			"%v\ncaused by:\n%v",
			exception.FailedToReadBody.Error(),
			err.Error(),
		)
		s.log.Error(m)
		return nil, exception.FailedToReadBody
	}

	return body, nil
}

func (s *server) internalServerError(w http.ResponseWriter, e error) {
	s.log.Error(exception.InternalServerError.Error(), zap.Error(e))
	http.Error(w, e.Error(), http.StatusInternalServerError)
}

func (s *server) badRequest(w http.ResponseWriter, m string) {
	s.log.Debug(m)
	http.Error(w, m, http.StatusBadRequest)
}

func (s *server) notFound(w http.ResponseWriter, r *http.Request, uri string) {
	s.log.Debug(exception.NotFound.Error(), zap.String("uri", uri))
	http.NotFound(w, r)
}
