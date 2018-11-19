// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package server

import "net/http"

var (
	IndexPath = "/"

	ConnectPath = "/connect"
	ClosePath   = "/close"

	ListOptPath = "/option/list"
	SetOptPath  = "/option/set"

	StartPath = "/start"

	SetPositionPath = "/position/set"

	GetValuesPath = "/study/values/list"

	InitAnalyze  = "/analyze/init"
	StartAnalyze = "/analyze/start"
)

type ShogiProxyServer interface {
	// ルーターの役目
	// メソッドチェックしたりもする方
	Handling(w http.ResponseWriter, r *http.Request)

	// HTML を返す
	// method: GET
	// returns: html
	//ServeHome(http.ResponseWriter, *http.Request)

	// Engine に接続する
	// method: POST
	// returns: ok | error
	//Connect(http.ResponseWriter, *http.Request)

	// Engine を終了する
	// method: POST
	// returns: ok | error
	//Close(http.ResponseWriter, *http.Request)

	// Option 一覧を返す
	// method: GET
	// returns: list | error
	//ListOption(http.ResponseWriter, *http.Request)

	// Option を設定する
	// method: POST
	// returns: ok | error
	//SetOption(http.ResponseWriter, *http.Request)

	// 局面のセット
	// method: POST
	// returns: ok | error
	//SetPosition(http.ResponseWriter, *http.Request)

	// newgame.
	// method: POST
	// returns: ok | error
	//Start(http.ResponseWriter, *http.Request)

	// 評価値一覧を返す
	// method: GET
	// returns: list | error
	//GetValues(http.ResponseWriter, *http.Request)

	// 棋譜解析の初期化。棋譜を渡す
	// method: POST
	// returns: ok | error
	//InitAnalyze(http.ResponseWriter, *http.Request)

	// 解析する。結果を返す
	// method: POST
	// returns: list | error
	//StartAnalyze(http.ResponseWriter, *http.Request)

	// TODO: 対局用のAPI
}
