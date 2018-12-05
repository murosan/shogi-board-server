// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package from_usi

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/murosan/shogi-proxy-server/app/domain/entity/shogi"
	"github.com/murosan/shogi-proxy-server/app/domain/exception"
)

func (fu *FromUsi) Move(s string) (m *shogi.Move, err error) {
	defer func() {
		if rec := recover(); rec != nil {
			m = nil
			err = exception.UnknownCharacter.WithMsg(fmt.Sprintf("%v", rec))
		}
	}()

	a := strings.Split(strings.TrimSpace(s), "")

	if strings.Contains(s, "*") {
		// 持ち駒から
		piece, err := fu.Piece(a[0])
		if err != nil {
			return nil, exception.UnknownCharacter
		}

		return &shogi.Move{
			Source:  []int{-1, -1},
			Dest:    []int{fu.column(a[2]), fu.row(a[3])},
			PieceId: piece,
		}, nil
	}

	return &shogi.Move{
		Source:     []int{fu.column(a[0]), fu.row(a[1])},
		Dest:       []int{fu.column(a[2]), fu.row(a[3])},
		IsPromoted: len(a) == 5 && a[4] == "+",
	}, nil
}

func (fu *FromUsi) column(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	if i < 1 || i > 9 {
		panic(exception.InvalidColumnNumber)
	}
	return i - 1 // 0-8 にする
}

func (fu *FromUsi) row(s string) int {
	if len(s) != 1 {
		panic(exception.InvalidRowNumber)
	}
	r := []rune(s)[0]
	if r < 97 || r > 105 {
		panic(exception.InvalidRowNumber)
	}
	return int(r - 97)
}
