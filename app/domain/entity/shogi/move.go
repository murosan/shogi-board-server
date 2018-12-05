// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package shogi

// 指し手を表す
// 初手で76歩なら
// Move{Source: []int{6, 6}, Dest: []int{6, 5}}
// 自然に 0 始まり。だが、右上から数える。
// 持ち駒の歩を打って、59歩なら
// Move{Source[]int{-1, -1}, Dest: []int{4, 8}, PieceId: 1, Extras: []string{"打"}}
type Move struct {
	// 移動元 TODO: Point型とか作るかなぁ。
	// []int{column, row} で表す。 column/row は盤上の駒を動かすなら 0-8
	// []int{-1, -1} なら持ち駒から打つことを意味する。
	Source []int `json:"source"`

	// 移動先
	// []int{column, row} で表す。 column/row は 0-8
	Dest []int `json:"dest"`

	// 盤上の駒を動かすときは常に 0。USIのmoveに情報が入らないから・・・
	// 持ち駒から打つ時はその駒のIDが入る。
	// TODO: 盤上の駒を動かすときもないとおかしいので、そのうち入るようにする
	PieceId int `json:"pieceId"`

	// 成ったかどうか。成: true, 不成: false
	IsPromoted bool `json:"isPromoted"`
}
