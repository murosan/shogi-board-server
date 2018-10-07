// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/murosan/shogi-proxy-server/pkg/config"
	"github.com/murosan/shogi-proxy-server/pkg/engine"
)

var (
	addr = flag.String("addr", "127.0.0.1:8080", "http service address")
)

func ping(ws *websocket.Conn, done chan struct{}) {
	ticker := time.NewTicker(config.PingPeriod)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			if err := ws.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(config.WriteWait)); err != nil {
				log.Println("ping:", err)
			}
		case <-done:
			return
		}
	}
}

func internalError(ws *websocket.Conn, msg string, err error) {
	log.Println(msg, err)
	ws.WriteMessage(websocket.TextMessage, []byte("Internal server error."))
}

var upgrader = websocket.Upgrader{}

func serveWs(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}

	defer ws.Close()

	if err := engine.Engine.Cmd.Start(); err != nil {
		panic(err)
	}

	if err != nil {
		internalError(ws, "start:", err)
		return
	}

	go engine.Read(ws, engine.Engine.Stdout, engine.Engine.Done)
	go ping(ws, engine.Engine.Done)

	engine.Exec(ws, engine.Engine.Stdin)

	select {
	case <-engine.Engine.Done:
	case <-time.After(time.Second):
		<-engine.Engine.Done
	}

	if err := engine.Engine.Cmd.Wait(); err != nil {
		log.Println("wait:", err)
	}
}

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
	config.Load()
	engine.Connect()
	defer engine.Close()

	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", serveWs)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
