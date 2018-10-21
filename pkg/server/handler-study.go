// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/murosan/shogi-proxy-server/pkg/engine"
	"github.com/murosan/shogi-proxy-server/pkg/msg"
)

func (s *Server) SetPosition(w http.ResponseWriter, r *http.Request) {

	if r.Header.Get("Content-Type") != "application/json" {
		m := "Content-Type was not application/json."
		http.Error(w, m, http.StatusBadRequest)
		log.Println(m)
		return
	}

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
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(msg.FailedToReadBody.Error() + err.Error())
		return
	}

	//parse json
	var jsonBody map[string]interface{}
	err = json.Unmarshal(body, &jsonBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(msg.FailedToParseBody.Error() + err.Error())
		return
	}
	fmt.Printf("%v\n", jsonBody)

	// TODO: positionの実行

	w.WriteHeader(http.StatusOK)
}

func (s *Server) Start(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (s *Server) GetValues(w http.ResponseWriter, r *http.Request) {
	// TODO
}
