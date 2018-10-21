// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package go2usi

import (
	"github.com/murosan/shogi-proxy-server/pkg/msg"
	"strconv"
)

func convPieceId(i int) (s string, e error) {
	switch i {
	case 1:
		s = "P"
	case -1:
		s = "p"
	case 2:
		s = "L"
	case -2:
		s = "l"
	case 3:
		s = "N"
	case -3:
		s = "n"
	case 4:
		s = "S"
	case -4:
		s = "s"
	case 5:
		s = "G"
	case -5:
		s = "g"
	case 6:
		s = "B"
	case -6:
		s = "b"
	case 7:
		s = "R"
	case -7:
		s = "r"
	case 8:
		s = "K"
	case -8:
		s = "k"
	case 11:
		s = "+P"
	case -11:
		s = "+p"
	case 12:
		s = "+L"
	case -12:
		s = "+l"
	case 13:
		s = "+N"
	case -13:
		s = "+n"
	case 14:
		s = "+S"
	case -14:
		s = "+s"
	case 16:
		s = "+B"
	case -16:
		s = "+b"
	case 17:
		s = "+R"
	case -17:
		s = "+r"
	default:
		e = msg.InvalidPieceId.WithMsg("PieceIdが不正です id=" + strconv.Itoa(i))
	}
	return
}
