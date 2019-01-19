// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package option

// UpdateOptionValue オプションを更新するときに API で受ける値
// gRPC を検討したい
// Engine 型の UpdateOption に渡す
type UpdateOptionValue struct {
	// オプション名
	Name string `json:"name"`

	// button, check, spin, combo, string, filename のどれか
	Type string `json:"type"`

	// 新しい値
	Value interface{} `json:"value"`
}
