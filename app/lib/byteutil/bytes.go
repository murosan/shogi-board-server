// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package byteutil

import "bytes"

// b の中に t があるか
// ない時は -1 を返す
func IndexOfBytes(b [][]byte, t []byte) int {
	if b == nil || t == nil {
		return -1
	}
	for n, e := range b {
		if bytes.Equal(e, t) {
			return n
		}
	}
	return -1
}

// b1 と b2 が等しければ true
func EqualBytes(b1, b2 [][]byte) bool {
	if (b1 == nil) != (b2 == nil) {
		return false
	}
	if len(b1) != len(b2) {
		return false
	}
	for i, b := range b1 {
		if !bytes.Equal(b, b2[i]) {
			return false
		}
	}
	return true
}
