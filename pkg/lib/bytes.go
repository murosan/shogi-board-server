package lib

import "bytes"

// b の中に t があるか
// ない時は -1 を返す
func IndexOfBytes(b [][]byte, t []byte) int {
	for n, e := range b {
		if bytes.Equal(e, t) {
			return n
		}
	}
	return -1
}
