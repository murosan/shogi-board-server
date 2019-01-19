// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package engine

import (
	"bufio"

	"github.com/murosan/shogi-board-server/app/domain/entity/engine/option"
	"github.com/murosan/shogi-board-server/app/domain/entity/engine/result"
	"github.com/murosan/shogi-board-server/app/domain/entity/engine/state"
)

// Engine 将棋エンジンを表す
type Engine interface {
	// GetName 将棋エンジンの名前を返す
	// JSON のキーではなく、将棋エンジンから最初に出力されるもの
	GetName() string
	// SetName 将棋エンジンの名前を変更する。GetName の値が変更される
	SetName(string)

	// GetAuthor 将棋エンジンの作者を返す
	// 将棋エンジンから最初に出力されるもの
	GetAuthor() string
	// SetAuthor 将棋エンジンの作者を変更する。GetAuthor の値が変更される
	SetAuthor(string)

	// SetOption オプションを Engine で保持している OptMap に追加する
	// 将棋エンジンから最初に出力されるものを追加するために使う
	// USI の setoption を実行するわけではないので注意
	SetOption(string, option.Option)
	// GetOptions オプションの一覧を返す
	// Engine で保持している OptMap から一覧を取得している
	GetOptions() option.OptMap
	// UpdateOption オプションの値を更新する
	// USI の setoption が実行される
	UpdateOption(option.UpdateOptionValue) error

	// SetState 将棋エンジンの状態を更新する
	SetState(state.State)
	// GetState 将棋エンジンの現在の状態を取得する
	GetState() state.State

	// SetResult 将棋エンジンの思考結果をセットする
	// USI の info をパースして、pv や multiPv だった場合に実行される
	SetResult(i result.Info, key int)
	// GetResult 将棋エンジンの思考結果を取得する
	// SetResult によって保持されている値を返す
	GetResult() *result.Result
	// FlushResult 将棋エンジンの思考結果一覧を削除する
	// Position を更新したときに実行する
	FlushResult()

	// Lock mutex lock
	Lock()
	// Unlock mutex unlock
	Unlock()

	// Exec 将棋エンジンに対して標準出力を渡す(コマンド実行)
	Exec([]byte) error

	// Start 将棋エンジンに対して、思考開始コマンドを実行(USI の go inf)
	Start() error
	// Close 将棋エンジンとの接続を切る。scanner 等も捨てる
	Close() error

	// GetScanner 将棋エンジンが出力した値を読み取る Scanner を作る
	GetScanner() *bufio.Scanner

	// GetChan 将棋エンジンが出力した値が送信されるチャネルを返す
	// このチャネルにたいして scanner で読み取った値を流し続ける
	GetChan() chan []byte
}
