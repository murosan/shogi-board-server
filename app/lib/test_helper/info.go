// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package testhelper

import (
	"reflect"

	"github.com/murosan/shogi-board-server/app/domain/entity/engine/result"
)

// InfoEquals テスト用メソッド 2つの Info が同じかどうか判定する
// 同じ: true
// 違う: false
func InfoEquals(a, b result.Info) bool {
	return reflect.DeepEqual(a.Values, b.Values) &&
		a.Score == b.Score &&
		MoveSliceEquals(a.Moves, b.Moves)
}
