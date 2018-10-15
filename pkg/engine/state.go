// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package engine

// TODO: 外から触れないようにする
// State構造体でも作る

var (
	// 起動前
	NotConnected = struct{}{}

	// 接続済み. usinewgame 前
	Connected = struct{}{}

	// usinewgame の後、go コマンドを待っている状態
	StandBy = struct{}{}

	// 思考中
	Thinking = struct{}{}
)
