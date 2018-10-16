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

func SetPosition(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Printf("%s %s", msg.MethodNotAllowed, SetPositionPath)
		http.Error(w, msg.MethodNotAllowed.Error(), http.StatusBadRequest)
		return
	}

	log.Println(r.Method + " " + SetPositionPath)

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Content-Type was not application/json.", http.StatusBadRequest)
		log.Println("Content-Type was not application/json.")
		return
	}

	// TODO: エラー
	if engine.Egn == nil {
		e := msg.EngineIsNotRunning.WithMsg("You need to start engine first.")
		http.Error(w, e.Error(), http.StatusBadRequest)
		log.Println(e)
		return
	}

	l, err := strconv.Atoi(r.Header.Get("Content-Length"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
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
