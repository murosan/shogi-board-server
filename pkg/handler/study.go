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

func StudyInit(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Printf("%s %s", msg.MethodNotAllowed, StudyInitPath)
		http.Error(w, msg.MethodNotAllowed.Error(), http.StatusBadRequest)
		return
	}
	log.Println(r.Method + " " + StudyInitPath)

	// 将棋エンジンが思考中ならStop
	if engine.Engine.State == engine.Thinking {
		if e := engine.Engine.Exec(usi.CmdStop); e != nil {
			http.Error(w, msg.FailedToStop.Error(), http.StatusInternalServerError)
			return
		}
	}

	if e := engine.Engine.Exec(usi.CmdNew); e != nil {
		http.Error(w, msg.FailedToStart.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
