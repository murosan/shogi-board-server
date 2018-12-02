// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package from_usi

import (
	"strconv"
	"strings"

	"github.com/murosan/shogi-proxy-server/app/domain/entity/shogi"
	"github.com/murosan/shogi-proxy-server/app/domain/exception"
)

func (fu *FromUsi) Move(s string) (*shogi.Move, error) {
	a := strings.Split(strings.TrimSpace(s), "")
	if len(a) < 4 {
		return nil, exception.UnknownCharacter
	}

	if strings.Contains(s, "*") {
		// 持ち駒から
		piece, err := fu.Piece(a[0])
		if err != nil {
			return nil, exception.UnknownCharacter
		}

		column, err := fu.column(a[2])
		if err != nil {
			return nil, exception.UnknownCharacter
		}

		row, err := fu.row(a[3])
		if err != nil {
			return nil, exception.UnknownCharacter
		}

		return &shogi.Move{
			Source:  []int{-1, -1},
			Dest:    []int{column, row},
			PieceId: piece,
			Extra:   shogi.FromCaptured,
		}, nil
	}

	sourceCol, err := fu.column(a[0])
	if err != nil {
		return nil, exception.UnknownCharacter
	}

	sourceRow, err := fu.row(a[1])
	if err != nil {
		return nil, exception.UnknownCharacter
	}

	destCol, err := fu.column(a[2])
	if err != nil {
		return nil, exception.UnknownCharacter
	}

	destRow, err := fu.row(a[3])
	if err != nil {
		return nil, exception.UnknownCharacter
	}

	extra := shogi.None
	if len(a) == 5 {
		if a[4] != "+" {
			return nil, exception.UnknownCharacter
		}
		extra = shogi.Promote
	}

	return &shogi.Move{
		Source: []int{sourceCol, sourceRow},
		Dest:   []int{destCol, destRow},
		Extra:  extra,
	}, nil
}

func (fu *FromUsi) column(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	if i < 1 || i > 9 {
		return 0, exception.InvalidColumnNumber
	}
	return i - 1, nil // 0-8 にする
}

func (fu *FromUsi) row(s string) (int, error) {
	if len(s) != 1 {
		return 0, exception.InvalidRowNumber
	}
	r := []rune(s)[0]
	if r < 97 || r > 105 {
		return 0, exception.InvalidRowNumber
	}
	return int(r - 97), nil
}
