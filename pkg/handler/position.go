// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/murosan/shogi-proxy-server/pkg/engine"
)

func Position(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Println(methodNotAllowed + " " + r.Method + " " + positionPath)
		http.Error(w, methodNotAllowed, http.StatusBadRequest)
		return
	}

	log.Println(r.Method + " " + positionPath)

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Content-Type was not application/json.", http.StatusBadRequest)
		log.Println("Content-Type was not application/json.")
		return
	}

	// TODO: エラー
	if engine.Engine == nil {
		http.Error(w, notRunning+"\nYou need to start Engine first.", http.StatusBadRequest)
		log.Println(notRunning)
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
		log.Println("Failed to read body. " + err.Error())
		return
	}

	//parse json
	var jsonBody map[string]interface{}
	err = json.Unmarshal(body, &jsonBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Failed to parse body. " + err.Error())
		return
	}
	fmt.Printf("%v\n", jsonBody)

	// TODO: positionの実行

	w.WriteHeader(http.StatusOK)
}
