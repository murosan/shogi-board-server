// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package converter

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/murosan/shogi-board-server/app/domain/entity/shogi"
	"github.com/murosan/shogi-board-server/app/domain/exception"
	pb "github.com/murosan/shogi-board-server/app/proto"
)

// ToUSI Go から USI に変換するインターフェース
type ToUSI interface {
	Piece(int32) (string, error)
	Position(*pb.Position) ([]byte, error)
}

type toUSI struct{}

// NewToUSI 新しい ToUSI を返す
func NewToUSI() ToUSI {
	return toUSI{}
}

func (tu toUSI) Piece(i int32) (s string, e error) {
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
		e = exception.InvalidPieceID.WithMsg("PieceIDが不正です id=" + fmt.Sprint(i))
	}
	return
}

// TODO: クソコードすぎる
func (tu toUSI) Position(p *pb.Position) ([]byte, error) {
	arr := make([]string, 9)
	for i, r := range p.Pos {
		usir, err := tu.row(r)
		if err != nil {
			return nil, err
		}
		arr[i] = usir
	}

	s := []byte("position sfen " + strings.Join(arr, "/"))
	if p.Turn == 1 {
		s = append(s, []byte(" b ")...)
	} else {
		s = append(s, []byte(" w ")...)
	}

	c0, c1 := p.Cap0, p.Cap1

	// TODO
	caps := []byte("")
	for i, c := range c0 {
		if c != 0 {
			p, err := tu.Piece(int32(i + 1))
			if err != nil {
				return nil, err
			}
			caps = append(caps, []byte(fmt.Sprint(c)+p)...)
		}
	}
	for i, c := range c1 {
		if c != 0 {
			p, err := tu.Piece(int32(-i - 1))
			if err != nil {
				return nil, err
			}
			caps = append(caps, []byte(fmt.Sprint(c)+p)...)
		}
	}
	if len(caps) == 0 {
		s = append(s, []byte("-")...)
	} else {
		s = append(s, caps...)
	}
	return append(s, []byte(" "+fmt.Sprint(p.MoveCount))...), nil
}

func (tu toUSI) row(r *pb.Row) (s string, e error) {
	emp := 0
	for _, id := range r.Row {
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
