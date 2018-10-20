// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package server

import (
	"log"
	"net/http"
)

func (s *Server) Connect(w http.ResponseWriter, r *http.Request) {
	if err := s.cli.Connect(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Successfully connected.")
	w.WriteHeader(http.StatusCreated)
}

func (s *Server) Close(w http.ResponseWriter, r *http.Request) {
	if err := s.cli.Connect(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Successfully closed.")
	w.WriteHeader(http.StatusOK)
}
