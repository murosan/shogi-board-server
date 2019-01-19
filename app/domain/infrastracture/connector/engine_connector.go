// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package connector

import (
	"github.com/murosan/shogi-board-server/app/domain/infrastracture/engine"
)

// Connector 将棋エンジンの接続を扱う
type Connector interface {
	// 将棋エンジンと接続する
	Connect() error

	// 接続を切る
	Close() error

	// WithEngine 将棋エンジンの JSON キー と、callback を受け取って
	// ConnectionPool から Engine を取得し、
	// Engine を callback に渡して実行する
	WithEngine(string, func(engine.Engine)) error
}
