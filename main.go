// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/murosan/shogi-proxy-server/app/server"
	"github.com/murosan/shogi-proxy-server/app/service/connector"
	"github.com/murosan/shogi-proxy-server/app/service/converter"
)

var (
	addr = flag.String("addr", "127.0.0.1:8080", "http service address")
)

func main() {
	conn := connector.UseConnector()
	defer conn.Close() // for safety

	s := server.NewServer(conn, converter.UseFromJson(), converter.UseToUsi())

	log.Println("Listening. localhost:" + *addr)
	http.HandleFunc(server.IndexPath, s.Handling("GET", s.ServeHome))
	http.HandleFunc(server.ConnectPath, s.Handling("POST", s.Connect))
	http.HandleFunc(server.ClosePath, s.Handling("POST", s.Close))
	http.HandleFunc(server.ListOptPath, s.Handling("GET", s.ListOption))
	http.HandleFunc(server.SetOptPath, s.Handling("POST", s.SetOption))
	http.HandleFunc(server.SetPositionPath, s.Handling("POST", s.ContentTypeCheck("application/json", s.SetPosition)))
	http.HandleFunc(server.StartPath, s.Handling("POST", s.Start))
	http.HandleFunc(server.GetValuesPath, s.Handling("GET", s.GetValues))
	http.HandleFunc(server.InitAnalyze, s.Handling("POST", s.InitAnalyze))
	http.HandleFunc(server.StartAnalyze, s.Handling("POST", s.StartAnalyze))
	log.Fatal(http.ListenAndServe(*addr, nil))
}
