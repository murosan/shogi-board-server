// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package handler

import (
	"bytes"
	"log"
	"net/http"

	"github.com/murosan/shogi-proxy-server/pkg/engine"
	"github.com/murosan/shogi-proxy-server/pkg/msg"
	"github.com/murosan/shogi-proxy-server/pkg/usi"
)

var (
	startPath    = "/start"
	quitPath     = "/quit"
	positionPath = "/position"

	connected = "Successfully connected."
)

func Start(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Printf("%s %s", msg.MethodNotAllowed, startPath)
		http.Error(w, msg.MethodNotAllowed.Error(), http.StatusMethodNotAllowed)
		return
	}

	log.Println(r.Method + " " + startPath)

	if engine.Engine != nil {
		log.Println(msg.EngineIsAlreadyRunning)
		http.Error(w, msg.EngineIsAlreadyRunning.Error(), http.StatusBadRequest)
		return
	}

	engine.Connect()
	if err := engine.Engine.Cmd.Start(); err != nil {
		log.Fatalln(msg.FailedToStart)
		http.Error(w, msg.FailedToStart.Error(), http.StatusInternalServerError)
		return
	}

	go engine.CatchEngineOutput()

	engine.Engine.Mux.Lock()

	for _, mess := range usi.StartCmds {
		engine.Engine.Stdin.Write(append(mess, '\n'))
		log.Println("Send message. " + string(mess))
		if bytes.Equal(mess, usi.CmdUsi) || bytes.Equal(mess, usi.CmdIsReady) {
			waitStart()
		}
	}

	engine.Engine.Mux.Unlock()

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
