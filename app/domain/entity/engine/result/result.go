// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package result

type Result struct {
	Values map[int]*Info `json:"values"`
}

// Values に値を追加する
func (r *Result) Set(i *Info, key int) {
	r.Values[key] = i
}

// Values をリセットする
// 新らしい Position がセットされた時に実行する
func (r *Result) Flush() {
	r.Values = make(map[int]*Info)
}
