// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test_helper

import (
	"github.com/murosan/shogi-proxy-server/app/domain/entity/shogi"
	"github.com/murosan/shogi-proxy-server/app/lib/intutil"
)

func MoveEquals(a, b shogi.Move) bool {
	return intutil.SliceEquals(a.Source, b.Source) &&
		intutil.SliceEquals(a.Dest, b.Dest) &&
		a.PieceId == b.PieceId &&
		a.IsPromoted == b.IsPromoted
}

func MoveSliceEquals(a, b []shogi.Move) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if !MoveEquals(v, b[i]) {
			return false
		}
	}
	return true
}
