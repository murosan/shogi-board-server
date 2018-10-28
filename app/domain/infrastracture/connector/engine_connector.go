// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package connector

import (
	"github.com/murosan/shogi-proxy-server/app/domain/entity/engine/option"
	"github.com/murosan/shogi-proxy-server/app/domain/entity/engine/state"
)

type Connector interface {
	// 将棋エンジンと接続する
	Connect() error

	// 接続を切る
	Close() error

	// 将棋エンジンにコマンドを実行する
	Exec(*[]byte) error

	// 将棋エンジンの出力を受け取り続ける
	CatchOutput()

	// State の更新
	SetState(s state.State)

	// State の取得
	GetState() state.State

	// ID(author | name) をセットする
	SetId(*[]byte, *[]byte) error

	// 設定可能な将棋エンジンのオプションを追加する
	// オプション一覧は自分で持つ
	SetupOption(option.Option)

	OptionList() []option.Option
}