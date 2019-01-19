// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package shogi

// Move 指し手を表す
// 初手で76歩なら
// Move{
//   Source: { Row: 6, Column: 6 },
//   Dest: { Row: 5, Column: 6 },
// }
// 自然に 0 始まり。だが、右上から数える。
// 持ち駒の歩を打って、59歩なら
// Move{
//   Source: { Row: -1, Column: -1 },
//   Dest: { Row: 8, Column: 4 },
//   PieceId: 1,
//   IsPromoted: false
// }
type Move struct {
	// Source 移動元
	// Point{-1, -1} なら持ち駒から打つことを意味する。
	Source Point `json:"source"`

	// Dest 移動先
	Dest Point `json:"dest"`

	// PieceID 駒のID
	// 盤上の駒を動かすときは常に 0。USIのmoveに情報が入らないから・・・
	// 持ち駒から打つ時はその駒のIDが入る。
	// TODO: 盤上の駒を動かすときもないとおかしいので、そのうち入るようにする
	PieceID int `json:"pieceId"`

	// IsPromoted 成ったかどうか。成: true, 不成: false
	IsPromoted bool `json:"isPromoted"`
}
