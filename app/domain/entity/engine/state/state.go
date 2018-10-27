// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package state

type State int

const (
	// 起動前
	NotConnected State = 1

	// 接続済み. usinewgame 前
	Connected State = 2

	// usinewgame の後、go コマンドを待っている状態
	StandBy State = 3

	// 思考中
	Thinking State = 4
)
