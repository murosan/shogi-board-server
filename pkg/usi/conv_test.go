// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package usi

import (
	"bytes"
	"testing"
)

func TestConvert(t *testing.T) {
	r := "lnsgk1snl/6gb1/p1pppp2p/6R2/9/1rP6/P2PPPP1P/1BG6/LNS1KGSNL w 3P2p 1"
	json := `
{
  "version": 1,
  "command": "position",
  "data": {
    "position": [
      [-2, -3, -4, -5, -8, 0, -4, -3, -2],
      [0, 0, 0, 0, 0, 0, -5, -6, 0],
      [-1, 0, -1, -1, -1, -1, 0, 0, -1],
      [0, 0, 0, 0, 0, 0, 7, 0, 0],
      [0, 0, 0, 0, 0, 0, 0, 0, 0],
      [0, -7, 1, 0, 0, 0, 0, 0, 0],
      [1, 0, 0, 1, 1, 1, 1, 0, 1],
      [0, 6, 5, 0, 0, 0, 0, 0, 0],
      [2, 3, 4, 0, 8, 5, 4, 3, 2]
    ],
    "cap_0": [3, 0, 0, 0, 0, 0, 0],
    "cap_1": [2, 0, 0, 0, 0, 0, 0],
    "turn": 1
  }
}
`
	usi, err := Convert([]byte(json))

	if err != nil || bytes.Equal(usi[0], []byte(r)) {
		t.Errorf("\nResult:   %v\nError:    %v\nExpected: %v", usi, err, r)
	}
}
