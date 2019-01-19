// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package state

// State 将棋エンジンの状態を表す
type State int

const (
	// NotConnected 起動前
	NotConnected State = 1

	// Connected 接続済み. usinewgame 前
	Connected State = 2

	// StandBy USI の usinewgame を実行後、思考中ではないとき
	StandBy State = 3

	// Thinking 思考中
	Thinking State = 4
)
