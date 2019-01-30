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

	// GetEngine は指定された名前の Engine を返します
	GetEngine(string) (engine.Engine, error)

	// GetEngineNames は app.yml で設定された接続可能な将棋エンジンの名前一覧を返す
	GetEngineNames() []string
}
