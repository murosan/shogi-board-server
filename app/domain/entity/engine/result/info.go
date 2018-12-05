// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package result

import "github.com/murosan/shogi-proxy-server/app/domain/entity/shogi"

const (
	Depth    = "depth"
	SelDepth = "seldepth"
	Time     = "time"
	Nodes    = "nodes"
	HashFull = "hashfull"
	Nps      = "nps"
	Score    = "score"
	Pv       = "pv"
	MultiPv  = "multipv"
)

// TODO: ぱっと見伝わらないので、名前変えたい

type Info struct {
	// depth, seldepth, time, nodes, nps, hashfull が入る
	Values map[string]int `json:"values"`

	// cp, mate 両方
	Score int `json:"score"`

	// USIでいうpv
	Moves []shogi.Move `json:"moves"`
}

func NewInfo() *Info {
	return &Info{
		Values: make(map[string]int),
	}
}
