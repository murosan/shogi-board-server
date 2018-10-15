// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package handler

import (
	"log"
	"net/http"

	"github.com/murosan/shogi-proxy-server/pkg/engine"
	"github.com/murosan/shogi-proxy-server/pkg/msg"
)

func Quit(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Printf("%s %s", msg.MethodNotAllowed, QuitPath)
		http.Error(w, msg.MethodNotAllowed.Error(), http.StatusMethodNotAllowed)
		return
	}

	log.Println(r.Method + " " + QuitPath)

	if engine.Egn == nil {
		log.Println(msg.EngineIsNotRunning)
		http.Error(w, msg.EngineIsNotRunning.Error(), http.StatusBadRequest)
		return
	}

	engine.Close()
	if engine.Egn != nil {
		panic(msg.FailedToShutdown)
	}
	log.Println("Successfully closed.")
	w.WriteHeader(http.StatusOK)
}
