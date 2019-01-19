// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package result

// Result 将棋エンジンの思考結果を保持しておく型
type Result struct {
	// TODO: slice にしたい
	// エンジンから必ず順番通り&抜け番なしに出力される保証がない
	// &最大数が分からない(MultiPVがあるか分からないし、
	// MultiPVという名前かどうかも不明だ)
	Values map[int]*Info `json:"values"`
}

// NewResult Result を作成する
func NewResult() *Result {
	return &Result{make(map[int]*Info)}
}

// Set Values に値を追加する
func (r *Result) Set(i Info, key int) {
	r.Values[key] = &i
}

// Flush Values をリセットする
// 新しい Position がセットされた時に実行する
func (r *Result) Flush() {
	r.Values = make(map[int]*Info)
}
