// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package handler

import (
	"log"
	"net/http"

	"github.com/murosan/shogi-proxy-server/pkg/engine"
)

func Quit(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Println(methodNotAllowed + " " + r.Method + " " + quitPath)
		http.Error(w, methodNotAllowed, http.StatusMethodNotAllowed)
		return
	}

	log.Println(r.Method + " " + quitPath)

	// TODO: エラー
	if engine.Engine == nil {
		http.Error(w, notRunning, http.StatusBadRequest)
		log.Println(notRunning)
		return
	}

	// TODO: ちゃんと終了したかどうか確認できるようにしたい
	engine.Close()
	if engine.Engine != nil {
		panic("Cloud not shutdown engine successfully.")
	}
	log.Println("Successfully closed.")
	w.WriteHeader(http.StatusOK)
}
