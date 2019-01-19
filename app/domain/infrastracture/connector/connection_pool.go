// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package connector

import "github.com/murosan/shogi-board-server/app/domain/infrastracture/engine"

// ConnectionPool 将棋エンジンのコネクションを保持する
// 将来的に複数の将棋エンジンを切り替えるために用意している
type ConnectionPool interface {
	// config にある全ての Engine のコマンドを初期化する
	Initialize()

	// 名前を受け取って Engine を返す
	// TODO: 今はとりあえず1つのEngineだけを使えるようにしてあるので、
	// name を渡してない。あとで修正する
	NamedEngine() engine.Engine

	Remove() error
}
