// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package server

import (
	"bytes"
	"log"
	"net/http"
	"time"

	"github.com/murosan/shogi-proxy-server/pkg/engine"
	"github.com/murosan/shogi-proxy-server/pkg/msg"
	"github.com/murosan/shogi-proxy-server/pkg/usi"
)

var (
	connected = "Successfully connected."
)

func (s *Server) Connect(w http.ResponseWriter, r *http.Request) {
	// TODO: HTTPメソッドチェックとログは別のラッパーメソッドを書いて移す
	if r.Method != "GET" {
		log.Printf("%s %s", msg.MethodNotAllowed, ConnectPath)
		http.Error(w, msg.MethodNotAllowed.Error(), http.StatusMethodNotAllowed)
		return
	}
	log.Println(r.Method + " " + ConnectPath)

	if engine.Egn != nil {
		log.Println(msg.EngineIsAlreadyRunning)
		http.Error(w, msg.EngineIsAlreadyRunning.Error(), http.StatusBadRequest)
		return
	}

	engine.NewEngine()
	if err := engine.Egn.Cmd.Start(); err != nil {
		log.Fatalln(msg.FailedToConnect)
		http.Error(w, msg.FailedToConnect.Error(), http.StatusInternalServerError)
		return
	}

	go engine.CatchEngineOutput()
	engine.Egn.Mux.Lock()

	for _, mess := range usi.ConnectCmds {
		engine.Egn.Exec(mess)
		log.Println("Send message. " + string(mess))
		if bytes.Equal(mess, usi.CmdUsi) || bytes.Equal(mess, usi.CmdIsReady) {
			if e := waitStart(); e != nil {
				http.Error(w, msg.ConnectionTimeout.Error(), http.StatusInternalServerError)
				return
			}
		}
	}

	engine.Egn.State = engine.Connected
	engine.Egn.Mux.Unlock()
	log.Println(connected)
	w.WriteHeader(http.StatusOK)
}

func waitStart() error {
	timeout := make(chan struct{})
	go func() {
		time.Sleep(time.Second * 10)
		timeout <- struct{}{}
	}()
	for {
		select {
		case b := <-engine.Egn.EngineOut:
			log.Println("Received. " + string(b))
			if len(b) == 0 {
				continue
			}

			if bytes.Equal(b, usi.ResOk) || bytes.Equal(b, usi.ResReady) {
				return nil
			}

			// id でパースしてみて、失敗したら option でパース
			if e := engine.Egn.ParseId(b); e != nil {
				e := engine.Egn.ParseOpt(b)
				if e != nil {
					log.Println(e)
				}
			}
		case <-timeout:
			log.Println("Failed to connect to engine.")
			return msg.ConnectionTimeout
		}
	}
}

func (s *Server) Close(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Printf("%s %s", msg.MethodNotAllowed, ClosePath)
		http.Error(w, msg.MethodNotAllowed.Error(), http.StatusMethodNotAllowed)
		return
	}

	log.Println(r.Method + " " + ClosePath)

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
