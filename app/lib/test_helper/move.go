// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package testhelper

import (
	pb "github.com/murosan/shogi-board-server/app/proto"
)

// MoveEquals テスト用メソッド 2つの Move が同じかどうか判定する
// 同じ: true
// 違う: false
func MoveEquals(a, b *pb.Move) bool {
	return a.Source.Row == b.Source.Row &&
		a.Source.Column == b.Source.Column &&
		a.Dest.Row == b.Dest.Row &&
		a.Dest.Column == b.Dest.Column &&
		a.PieceID == b.PieceID &&
		a.IsPromoted == b.IsPromoted
}

// MoveSliceEquals テスト用メソッド 2つの Move スライスが同じかどうか判定する
// 同じ: true
// 違う: false
func MoveSliceEquals(a, b []*pb.Move) bool {
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
