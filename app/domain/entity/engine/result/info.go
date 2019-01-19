// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package result

import "github.com/murosan/shogi-board-server/app/domain/entity/shogi"

const (
	// Depth USIプロトコルの depth キー
	Depth string = "depth"

	// SelDepth USIプロトコルの seldepth キー
	SelDepth = "seldepth"

	// Time USIプロトコルの time キー
	Time = "time"

	// Nodes USIプロトコルの nodes キー
	Nodes = "nodes"

	// HashFull USIプロトコルの hashfull キー
	HashFull = "hashfull"

	// Nps USIプロトコルの nps キー
	Nps = "nps"

	// Score USIプロトコルの score キー
	Score = "score"

	// Pv USIプロトコルの pv キー
	Pv = "pv"

	// MultiPv USIプロトコルの multipv キー
	MultiPv = "multipv"
)

// TODO: ぱっと見伝わらないので、名前変えたい

// Info 将棋エンジンの思考結果
type Info struct {
	// depth, seldepth, time, nodes, nps, hashfull が入る
	Values map[string]int `json:"values"`

	// cp, mate 両方
	Score int `json:"score"`

	// USIでいうpv
	Moves []shogi.Move `json:"moves"`
}

// NewInfo 新しい Info を返す
func NewInfo() Info {
	return Info{
		Values: make(map[string]int),
	}
}
