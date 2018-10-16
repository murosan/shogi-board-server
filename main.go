// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/murosan/shogi-proxy-server/pkg/client"
	"github.com/murosan/shogi-proxy-server/pkg/config"
	"github.com/murosan/shogi-proxy-server/pkg/server"
	"github.com/murosan/shogi-proxy-server/pkg/msg"
)

var (
	addr = flag.String("addr", "127.0.0.1:8080", "http service address")
)

func serveHome(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, msg.MethodNotAllowed.Error(), http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home.html")
}

func main() {
	conf := config.NewConfig("./config.json")
	cli := client.NewClient(conf)
	defer cli.Close() // for safety

	// TODO: handlerパッケージはglobalのEngineを直接触っているのでテストできないので修正する
	log.Println("Listening. " + *addr)
	http.HandleFunc("/", serveHome)
	http.HandleFunc(server.ConnectPath, server.Connect)
	http.HandleFunc(server.QuitPath, server.Quit)
	http.HandleFunc(server.SetPositionPath, server.SetPosition)
	http.HandleFunc(server.StudyStartPath, server.StudyStart)
	http.HandleFunc(server.StudyStopPath, server.StudyStop)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
