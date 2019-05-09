// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package shogi

// Piece は駒
type Piece int

// USIPiece は USI 形式の駒
type USIPiece string

const (
	// Empty 空白マス
	Empty Piece = 0

	// Fu0 先手の歩
	Fu0 Piece = 1

	// Kyou0 先手の香車
	Kyou0 Piece = 2

	// Kei0 先手の桂馬
	Kei0 Piece = 3

	// Gin0 先手の銀
	Gin0 Piece = 4

	// Kin0 先手の金
	Kin0 Piece = 5

	// Kaku0 先手の角
	Kaku0 Piece = 6

	// Hisha0 先手の飛車
	Hisha0 Piece = 7

	// Gyoku0 先手の玉
	Gyoku0 Piece = 8

	// To0 Piece 先手のと金
	To0 Piece = 11

	// NariKyou0 先手の成香
	NariKyou0 Piece = 12

	// NariKei0 先手の成桂
	NariKei0 Piece = 13

	// NariGin0 先手の成銀
	NariGin0 Piece = 14

	// Uma0 先手の馬
	Uma0 Piece = 16

	// Ryu0 先手の龍
	Ryu0 Piece = 17

	// Fu1 後手の歩
	Fu1 = -Fu0

	// Kyou1 後手の香車
	Kyou1 = -Kyou0

	// Kei1 後手の桂馬
	Kei1 = -Kei0

	// Gin1 後手の銀
	Gin1 = -Gin0

	// Kin1 後手の金
	Kin1 = -Kin0

	// Kaku1 後手の角
	Kaku1 = -Kaku0

	// Hisha1 後手の飛車
	Hisha1 = -Hisha0

	// Gyoku1 後手の玉
	Gyoku1 = -Gyoku0

	// To1 後手のと金
	To1 = -To0

	// NariKyou1 後手の成香
	NariKyou1 = -NariKyou0

	// NariKei1 後手の成桂
	NariKei1 = -NariKei0

	// NariGin1 後手の成銀
	NariGin1 = -NariGin0

	// Uma1 後手の馬
	Uma1 = -Uma0

	// Ryu1 後手の龍
	Ryu1 = -Ryu0
)

const (
	// UsiFu0 USIの先手の歩
	UsiFu0 USIPiece = "P"

	// UsiKyou0 USIの先手の香車
	UsiKyou0 USIPiece = "L"

	// UsiKei0 USIの先手の桂馬
	UsiKei0 USIPiece = "N"

	// UsiGin0 USIの先手の銀
	UsiGin0 USIPiece = "S"

	// UsiKin0 USIの先手の金
	UsiKin0 USIPiece = "G"

	// UsiKaku0 USIの先手の角
	UsiKaku0 USIPiece = "B"

	// UsiHisha0 USIの先手の飛車
	UsiHisha0 USIPiece = "R"

	// UsiGyoku0 USIの先手の玉
	UsiGyoku0 USIPiece = "K"

	// UsiTo0 USIの先手のと金
	UsiTo0 USIPiece = "+P"

	// UsiNariKyou0 USIの先手の成香
	UsiNariKyou0 USIPiece = "+L"

	// UsiNariKei0 USIの先手の成桂
	UsiNariKei0 USIPiece = "+N"

	// UsiNariGin0 USIの先手の成銀
	UsiNariGin0 USIPiece = "+S"

	// UsiUma0 USIの先手の馬
	UsiUma0 USIPiece = "+B"

	// UsiRyu0 USIの先手の龍
	UsiRyu0 USIPiece = "+R"

	// UsiFu1 USIの後手の歩
	UsiFu1 USIPiece = "p"

	// UsiKyou1 USIの後手の香車
	UsiKyou1 USIPiece = "l"

	// UsiKei1 USIの後手の桂馬
	UsiKei1 USIPiece = "n"

	// UsiGin1 USIの後手の銀
	UsiGin1 USIPiece = "s"

	// UsiKin1 USIの後手の金
	UsiKin1 USIPiece = "g"

	// UsiKaku1 USIの後手の角
	UsiKaku1 USIPiece = "b"

	// UsiHisha1 USIの後手の飛車
	UsiHisha1 USIPiece = "r"

	// UsiGyoku1 USIの後手の玉
	UsiGyoku1 USIPiece = "k"

	// UsiTo1 USIの後手のと金
	UsiTo1 USIPiece = "+p"

	// UsiNariKyou1 USIの後手の成香
	UsiNariKyou1 USIPiece = "+l"

	// UsiNariKei1 USIの後手の成桂
	UsiNariKei1 USIPiece = "+n"

	// UsiNariGin1 USIの後手の成銀
	UsiNariGin1 USIPiece = "+s"

	// UsiUma1 USIの後手の馬
	UsiUma1 USIPiece = "+b"

	// UsiRyu1 USIの後手の龍
	UsiRyu1 USIPiece = "+r"
)
