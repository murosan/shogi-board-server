// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package handler

import (
	"log"
	"net/http"

	"github.com/murosan/shogi-proxy-server/pkg/engine"
	"github.com/murosan/shogi-proxy-server/pkg/msg"
	"github.com/murosan/shogi-proxy-server/pkg/usi"
)

func StudyStart(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Printf("%s %s", msg.MethodNotAllowed, StudyInitPath)
		http.Error(w, msg.MethodNotAllowed.Error(), http.StatusBadRequest)
		return
	}
	log.Println(r.Method + " " + StudyInitPath)

	if engine.Engine == nil {
		http.Error(w, msg.EngineIsNotRunning.Error(), http.StatusInternalServerError)
		return
	}

	// 将棋エンジンが思考中なら何もしない
	if engine.Engine.State == engine.Thinking {
		w.WriteHeader(http.StatusOK)
		return
	}

	// 将棋エンジンが StandBy 状態なら思考を開始させる
	if engine.Engine.State == engine.StandBy {
		if e := engine.Engine.Exec(usi.CmdGoInf); e != nil {
			log.Println(msg.FailedToExecCommand.Error() + " cmd: " + string(usi.CmdGoInf))
			http.Error(w, msg.FailedToExecCommand.Error(), http.StatusInternalServerError)
			return
		}
		log.Println("Start thinking...")
		w.WriteHeader(http.StatusOK)
		return
	}

	e := msg.FailedToExecCommand.WithMsg("Before to start 'study mode', 'usinewgame' command should be sent to the engine.")
	http.Error(w, e.Error(), http.StatusInternalServerError)
}

func StudyStop(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Printf("%s %s", msg.MethodNotAllowed, StudyInitPath)
		http.Error(w, msg.MethodNotAllowed.Error(), http.StatusBadRequest)
		return
	}
	log.Println(r.Method + " " + StudyInitPath)

	if engine.Engine == nil {
		http.Error(w, msg.EngineIsNotRunning.Error(), http.StatusInternalServerError)
		return
	}

	// 将棋エンジンが思考中ならStop
	if engine.Engine.State == engine.Thinking {
		if e := engine.Engine.Exec(usi.CmdStop); e != nil {
			log.Println(msg.FailedToExecCommand.Error() + " cmd: " + string(usi.CmdStop))
			http.Error(w, msg.FailedToExecCommand.Error(), http.StatusInternalServerError)
			return
		}
		log.Println("Stopped.")
	}

	// close してしまう
	// TODO
	engine.Close()

	w.WriteHeader(http.StatusOK)
	return
}
