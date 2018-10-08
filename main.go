// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/murosan/shogi-proxy-server/pkg/config"
	"github.com/murosan/shogi-proxy-server/pkg/engine"
	"github.com/murosan/shogi-proxy-server/pkg/handler"
)

var (
	addr = flag.String("addr", "127.0.0.1:8080", "http service address")
)

func serveHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home.html")
}

func main() {
	defer engine.Close() // for safety
	config.Load()

	log.Println("Listening. " + *addr)
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/start", handler.Start)
	http.HandleFunc("/quit", handler.Quit)
	http.HandleFunc("/position", handler.Position)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
