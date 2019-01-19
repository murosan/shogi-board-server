// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stringutil

// SliceContains a の中に b が含まれているか判定する
// 含まれていたら true を返す
func SliceContains(a []string, b string) bool {
	if a == nil {
		return false
	}
	for _, s := range a {
		if s == b {
			return true
		}
	}
	return false
}

// SliceEquals a と b が同じ値を持っているか判定する
// 同じ場合 true
func SliceEquals(a, b []string) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
