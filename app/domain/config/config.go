// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

import "go.uber.org/zap"

// Config は設定を取得するためのインターフェース
type Config interface {

	// 将棋エンジンの実行パスを名前(JSONのキー)から取得する
	// 名前は GetEngineNames で取得できる
	GetEnginePath(string) string

	// 将棋エンジンの名前一覧を取得する
	// この名前を使って GetEnginePath からパスを取得できる
	GetEngineNames() []string

	// ログの設定を取得
	GetLogConf() zap.Config
}
