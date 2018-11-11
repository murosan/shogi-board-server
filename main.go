// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/murosan/shogi-proxy-server/app/server"
	"github.com/murosan/shogi-proxy-server/app/service/config"
	"github.com/murosan/shogi-proxy-server/app/service/connector"
	"github.com/murosan/shogi-proxy-server/app/service/converter"
	"github.com/murosan/shogi-proxy-server/app/service/logger"
	"go.uber.org/zap"
)

var (
	addr       = flag.String("addr", "127.0.0.1:8080", "http service address")
	configPath = flag.String("config", "./config.json", "config path")
)

func main() {
	config.InitConfig(*configPath)
	conn := connector.UseConnector()
	defer conn.Close() // for safety

	s := server.NewServer(conn, converter.UseFromJson(), converter.UseToUsi())

	logger.Use().Info("Listening...", zap.String("address", *addr))
	http.HandleFunc("/", s.Handling)
	log.Fatalln(http.ListenAndServe(*addr, nil))
}
