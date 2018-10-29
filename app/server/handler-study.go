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

	"github.com/murosan/shogi-proxy-server/app/domain/entity/exception"
)

// Content-Type は application/json である必要がある
func (s *Server) SetPosition(w http.ResponseWriter, r *http.Request) {
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
		m := fmt.Sprintf("%v\ncaused by:\n%v", exception.FailedToReadBody.Error(), err.Error())
		http.Error(w, m, http.StatusInternalServerError)
		log.Println(m)
		return
	}

	pos, err := s.fj.Position(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	usi, err := s.tu.Position(pos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	if err := s.conn.Exec(&usi); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
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
