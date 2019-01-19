// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package shogi

// Point は位置を表す
// ７六 は Point{Row: 5, Column: 6}
// 持ち駒 は Point{Row: -1, Column: -1}
type Point struct {
	// Row は -1 | 0-8
	Row int `json:"row"`

	// Column は -1 | 0-8
	Column int `json:"column"`
}
