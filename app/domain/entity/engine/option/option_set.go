// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package option

type OptionSetValue struct {
	// オプション名
	Name string `json:"name"`

	// button, check, spin, combo, string, filename のどれか
	Type string `json:"type"`

	// 新しい値
	Value interface{} `json:"value"`
}
