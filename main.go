// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"net"

	pb "github.com/murosan/shogi-board-server/app/proto"
	"github.com/murosan/shogi-board-server/app/server"
	"github.com/murosan/shogi-board-server/app/service/config"
	"github.com/murosan/shogi-board-server/app/service/connector"
	"github.com/murosan/shogi-board-server/app/service/converter"
	"github.com/murosan/shogi-board-server/app/service/logger"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
	port          = flag.String("port", ":8080", "http server port")
	appConfigPath = flag.String("appConfig", "./config/app.yml", "application config path")
	logConfigPath = flag.String("logConfig", "./config/log.yml", "log config path")
)

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", *port)
	if err != nil {
		panic(err)
	}

	config.InitConfig(*appConfigPath, *logConfigPath)
	conn := connector.UseConnector()

	svr := server.NewServer(
		conn,
		converter.UseFromUSI(),
		converter.UseToUSI(),
		logger.Use(),
	)
	s := grpc.NewServer()
	pb.RegisterShogiBoardServer(s, svr)

	logger.Use().Info("Listening...", zap.String("port", *port))

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
