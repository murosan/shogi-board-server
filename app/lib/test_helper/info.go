// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test_helper

import (
	"reflect"

	"github.com/murosan/shogi-proxy-server/app/domain/entity/engine/result"
)

func InfoEquals(a, b *result.Info) bool {
	if a == nil || b == nil {
		return false
	}
	return reflect.DeepEqual(a.Values, b.Values) &&
		a.Score == b.Score &&
		MoveSliceEquals(a.Moves, b.Moves)
}
