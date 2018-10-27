// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package to_usi

import (
	"strconv"
	"strings"

	"github.com/murosan/shogi-proxy-server/app/domain/entity/shogi"
)

// TODO: クソコードすぎる
func (tu *ToUsi) Position(p shogi.Position) ([]byte, error) {
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

func (tu *ToUsi) row(r [9]int) (s string, e error) {
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
