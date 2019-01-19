// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package converter

import (
	"strconv"
	"strings"

	"github.com/murosan/shogi-board-server/app/domain/entity/shogi"
	"github.com/murosan/shogi-board-server/app/domain/exception"
)

// ToUSI Go から USI に変換するインターフェース
type ToUSI interface {
	Piece(int) (string, error)
	Position(shogi.Position) ([]byte, error)
}

type toUSI struct{}

// NewToUSI 新しい ToUSI を返す
func NewToUSI() ToUSI {
	return toUSI{}
}

func (tu toUSI) Piece(i int) (s string, e error) {
	switch i {
	case shogi.Fu0:
		s = shogi.UsiFu0
	case shogi.Fu1:
		s = shogi.UsiFu1
	case shogi.Kyou0:
		s = shogi.UsiKyou0
	case shogi.Kyou1:
		s = shogi.UsiKyou1
	case shogi.Kei0:
		s = shogi.UsiKei0
	case shogi.Kei1:
		s = shogi.UsiKei1
	case shogi.Gin0:
		s = shogi.UsiGin0
	case shogi.Gin1:
		s = shogi.UsiGin1
	case shogi.Kin0:
		s = shogi.UsiKin0
	case shogi.Kin1:
		s = shogi.UsiKin1
	case shogi.Kaku0:
		s = shogi.UsiKaku0
	case shogi.Kaku1:
		s = shogi.UsiKaku1
	case shogi.Hisha0:
		s = shogi.UsiHisha0
	case shogi.Hisha1:
		s = shogi.UsiHisha1
	case shogi.Gyoku0:
		s = shogi.UsiGyoku0
	case shogi.Gyoku1:
		s = shogi.UsiGyoku1
	case shogi.To0:
		s = shogi.UsiTo0
	case shogi.To1:
		s = shogi.UsiTo1
	case shogi.NariKyou0:
		s = shogi.UsiNariKyou0
	case shogi.NariKyou1:
		s = shogi.UsiNariKyou1
	case shogi.NariKei0:
		s = shogi.UsiNariKei0
	case shogi.NariKei1:
		s = shogi.UsiNariKei1
	case shogi.NariGin0:
		s = shogi.UsiNariGin0
	case shogi.NariGin1:
		s = shogi.UsiNariGin1
	case shogi.Uma0:
		s = shogi.UsiUma0
	case shogi.Uma1:
		s = shogi.UsiUma1
	case shogi.Ryu0:
		s = shogi.UsiRyu0
	case shogi.Ryu1:
		s = shogi.UsiRyu1
	default:
		e = exception.InvalidPieceID.WithMsg("PieceIDが不正です id=" + strconv.Itoa(i))
	}
	return
}

// TODO: クソコードすぎる
func (tu toUSI) Position(p shogi.Position) ([]byte, error) {
	arr := make([]string, 9)
	for i, r := range p.Pos {
		usir, err := tu.row(r)
		if err != nil {
			return nil, err
		}
		arr[i] = usir
	}

	s := []byte("position sfen " + strings.Join(arr, "/"))
	if p.Turn == 0 {
		s = append(s, []byte(" b ")...)
	} else {
		s = append(s, []byte(" w ")...)
	}

	c0, c1 := p.Cap0, p.Cap1
	if len(c0) == 0 && len(c1) == 0 {
		return append(s, []byte("- 1")...), nil
	}

	// TODO
	for i, c := range c0 {
		if c != 0 {
			p, err := tu.Piece(i + 1)
			if err != nil {
				return nil, err
			}
			s = append(s, []byte(strconv.Itoa(c)+p)...)
		}
	}
	for i, c := range c1 {
		if c != 0 {
			p, err := tu.Piece(-i - 1)
			if err != nil {
				return nil, err
			}
			s = append(s, []byte(strconv.Itoa(c)+p)...)
		}
	}
	return append(s, []byte(" "+strconv.Itoa(p.MoveCount))...), nil
}

func (tu toUSI) row(r [9]int) (s string, e error) {
	emp := 0
	for _, id := range r {
		// クソ
		if id == 0 {
			emp++
			continue
		}
		if emp != 0 {
			s += strconv.Itoa(emp)
			emp = 0
		}
		p, err := tu.Piece(id)
		if err != nil {
			return "", err
		}
		s += p
	}
	// マジクソ
	if emp != 0 {
		s += strconv.Itoa(emp)
	}
	return
}
