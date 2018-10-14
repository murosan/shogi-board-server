// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package engine

var (
	// 起動前
	NotConnected = struct{}{}

	// 接続済み
	Connected = struct{}{}

	// コマンド待ち
	StandBy = struct{}{}

	// 思考中
	Thinking = struct{}{}
)
