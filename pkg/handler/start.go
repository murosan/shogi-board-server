// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package handler

import (
	"bytes"
	"log"
	"net/http"

	"github.com/murosan/shogi-proxy-server/pkg/engine"
	"github.com/murosan/shogi-proxy-server/pkg/usi"
)

var (
	startPath    = "/start"
	quitPath     = "/quit"
	positionPath = "/position"

	// TODO: エラーをちゃんとする
	methodNotAllowed = "Method not allowed."
	alreadyRunning   = "Engine is already running."
	notRunning       = "Engine is not running."
	couldNotStart    = "Could not start engine."

	connected = "Successfully connected."
)

func Start(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Println(methodNotAllowed + " " + r.Method + " " + startPath)
		http.Error(w, methodNotAllowed, http.StatusMethodNotAllowed)
		return
	}

	log.Println(r.Method + " " + startPath)

	// TODO: エラー
	if engine.Engine != nil {
		log.Println(alreadyRunning)
		http.Error(w, alreadyRunning, http.StatusBadRequest)
		return
	}

	engine.Connect()
	if err := engine.Engine.Cmd.Start(); err != nil {
		http.Error(w, couldNotStart, http.StatusInternalServerError)
		log.Fatalln(couldNotStart)
		return
	}

	go engine.CatchEngineOutput()

	for _, msg := range usi.StartCmds {
		engine.Engine.Stdin.Write(append(msg, '\n'))
		log.Println("Send message. " + string(msg))
		if bytes.Equal(msg, usi.CmdUsi) || bytes.Equal(msg, usi.CmdIsReady) {
			waitStart()
		}
	}

	log.Println(connected)
	w.WriteHeader(http.StatusOK)
}

func waitStart() {
	// TODO: timeout
	for {
		select {
		case b := <-engine.Engine.EngineOut:
			log.Println("Received. " + string(b))
			if bytes.Equal(b, usi.ResOk) || bytes.Equal(b, usi.ResReady) {
				return
			}
		}
	}
}
