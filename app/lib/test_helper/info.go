// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test_helper

import "github.com/murosan/shogi-proxy-server/app/domain/entity/engine/result"

func InfoEquals(a, b *result.Info) bool {
	if a == nil || b == nil {
		return false
	}
	return a.Depth == b.Depth &&
		a.SelDepth == b.SelDepth &&
		a.Time == b.Time &&
		a.Nodes == b.Nodes &&
		a.HashRate == b.HashRate &&
		a.Score == b.Score &&
		MoveSliceEquals(a.Moves, b.Moves)
}
