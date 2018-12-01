// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package shogi

// TODO: docsに書く
//{
//  "pos": [9][9]pieceId,
//  "cap_0": [歩, 香, 桂, 銀, 金, 角, 飛], // 先手の持ち駒それぞれの枚数
//  "cap_1": [歩, 香, 桂, 銀, 金, 角, 飛], // 後手の持ち駒それぞれの枚数
//  "turn": turn(先手 == 0, 後手 == 1),
//  "move_count": 何手目か(初期局面 == 0)
//}
type Position struct {
	Pos       [9][9]int `json:"pos"`
	Cap0      [7]int    `json:"cap0"`
	Cap1      [7]int    `json:"cap1"`
	Turn      int       `json:"turn"`
	MoveCount int       `json:"moveCount"`
}
