// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package engine

import (
	"bytes"
	"io"
	"log"
	"time"

	"github.com/gorilla/websocket"
	"github.com/murosan/shogi-proxy-server/pkg/config"
	"github.com/murosan/shogi-proxy-server/pkg/usi"
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
		ml, err := usi.Convert(message)
		if err != nil {
			log.Println("convert error")
			// TODO: エラーをちゃんと返す
			if err := ws.WriteMessage(websocket.TextMessage, []byte("error")); err != nil {
				ws.Close()
				break
			}
		}

		for _, msg := range ml {
			log.Printf("コマンドを実行: %v\n", string(msg))
			if _, err := w.Write(append(msg, '\n')); err != nil {
				break
			}

			if bytes.Equal(msg, usi.CmdUsi) || bytes.Equal(msg, usi.CmdIsReady) {
				waitReady()
			}
		}
	}
}

// usiok か readyok を待つ
func waitReady() {
	for {
		select {
		case b := <-Engine.EngineOut:
			log.Println("受け取り: ", string(b))
			if bytes.Equal(b, usi.ResOk) || bytes.Equal(b, usi.ResReady) {
				return
			}
		}
	}
}
