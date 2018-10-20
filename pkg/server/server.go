// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package server

import (
	"github.com/murosan/shogi-proxy-server/pkg/client"
)

type Server struct {
	cli client.Connector
}

func NewServer(cli client.Connector) *Server {
	return &Server{cli}
}
