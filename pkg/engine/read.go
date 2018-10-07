// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package engine

import (
	"bufio"
	"io"
	"log"
	"time"

	"github.com/gorilla/websocket"
	"github.com/murosan/shogi-proxy-server/pkg/config"
)

// Engine の出力を読み取る
func Read(ws *websocket.Conn, r io.Reader, done chan struct{}) {
	defer func() {}()
	s := bufio.NewScanner(r)

	for s.Scan() {
		ws.SetWriteDeadline(time.Now().Add(config.WriteWait))
		// TODO: フロントに渡すデータに変換する
		b := s.Bytes()
		if err := ws.WriteMessage(websocket.TextMessage, b); err != nil {
			ws.Close()
			break
		}
	}

	if s.Err() != nil {
		log.Println("scan:", s.Err())
	}
	close(done)

	ws.SetWriteDeadline(time.Now().Add(config.WriteWait))
	ws.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	time.Sleep(config.CloseGracePeriod)
	ws.Close()
}
