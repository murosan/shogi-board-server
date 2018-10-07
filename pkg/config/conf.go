// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

import "time"

var (
	Conf Config
)

const (
	// Time allowed to write a message to the peer.
	WriteWait = 10 * time.Second

	// Maximum message size allowed from peer.
	MaxMessageSize = 8192

	// Time allowed to read the next pong message from the peer.
	PongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	PingPeriod = (PongWait * 9) / 10

	// Time to wait before force close on connection.
	CloseGracePeriod = 10 * time.Second
)
