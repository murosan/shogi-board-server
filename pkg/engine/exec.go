// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package engine

import (
	"github.com/gorilla/websocket"
	"github.com/murosan/shogi-proxy-server/pkg/config"
	"io"
	"time"
)

// WebSocket で受け取ったコマンドを USIプロトコル に変換して書き込む
func Exec(ws *websocket.Conn, w io.Writer) {
	defer ws.Close()
	ws.SetReadLimit(config.MaxMessageSize)
	ws.SetReadDeadline(time.Now().Add(config.PongWait))
	ws.SetPongHandler(func(string) error { ws.SetReadDeadline(time.Now().Add(config.PongWait)); return nil })
	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		// TODO: USIに変換する
		message = append(message, '\n')
		if _, err := w.Write(message); err != nil {
			break
		}
	}
}
