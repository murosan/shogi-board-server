// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package result

import "github.com/murosan/shogi-proxy-server/app/domain/entity/shogi"

// TODO: ぱっと見伝わらないので、名前変えたい
type Info struct {
	Depth    int           `json:"depth"`
	SelDepth int           `json:"selDepth"`
	Time     int           `json:"time"`
	Nodes    int           `json:"nodes"`
	HashRate int           `json:"hashRate"` // 0 <= a <= 1000
	Score    int           `json:"score"`    // cp, mate 両方
	Moves    []*shogi.Move `json:"moves"`    // USIでいうpv
}
