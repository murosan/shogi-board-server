// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package server

import "net/http"

var (
	// API gRPC への移行を考えたい

	// IndexPath HTML を返す API
	IndexPath = "/"

	// ConnectPath 将棋エンジンに接続する API
	ConnectPath = "/connect"
	// ClosePath 将棋エンジンとの接続を切る API
	ClosePath = "/close"

	// ListOptPath 将棋エンジンのオプション一覧を取得する API
	ListOptPath = "/option/list"
	// SetOptPath 将棋エンジンのオプションを更新する API
	SetOptPath = "/option/update"

	// StartPath 思考開始 API
	StartPath = "/start"
	// StopPath 思考停止 API
	StopPath = "/stop"

	// SetPositionPath 現在局面更新 API
	SetPositionPath = "/position/set"

	// GetResultPath 将棋エンジンの思考結果を取得する API
	GetResultPath = "/result/get"

	// InitAnalyze 棋譜解析を初期化する API
	InitAnalyze = "/analyze/init"
	// StartAnalyze 棋譜解析を開始する API
	StartAnalyze = "/analyze/start"
)

// ShogiBoardServer サーバーのインターフェース
type ShogiBoardServer interface {
	// ルーターの役目
	// メソッドチェックしたりもする方
	Handling(w http.ResponseWriter, r *http.Request)
}
