// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package server

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/murosan/shogi-proxy-server/pkg/engine"
	"github.com/murosan/shogi-proxy-server/pkg/msg"
)

// Content-Type は application/json である必要がある
func (s *Server) SetPosition(w http.ResponseWriter, r *http.Request) {
	// 将棋エンジンへ接続されてなければ BadRequest
	if s.cli.GetState() == engine.NotConnected {
		e := msg.EngineIsNotRunning.WithMsg("You need to start engine first.")
		http.Error(w, e.Error(), http.StatusBadRequest)
		log.Println(e.Error())
		return
	}

	l, err := strconv.Atoi(r.Header.Get("Content-Length"))
	if err != nil {
		http.Error(w, err.Error(), 411) // Length Required
		log.Println("Could not read Content-Length. " + err.Error())
		return
	}

	// read body
	body := make([]byte, l)
	l, err = r.Body.Read(body)
	if err != nil && err != io.EOF {
		m := fmt.Sprintf("%v\ncaused by:\n%v", msg.FailedToReadBody.Error(), err.Error())
		http.Error(w, m, http.StatusInternalServerError)
		log.Println(m)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Server) Start(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (s *Server) GetValues(w http.ResponseWriter, r *http.Request) {
	// TODO
}
