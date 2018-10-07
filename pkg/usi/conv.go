// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package usi

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/antonholmquist/jason"
)

type position struct {
	Version uint8  `json:"version"`
	Command string `json:"command"`
	Data    struct {
		Position [9][9]int `json:"position"`
		Cap0     []int     `json:"cap_0"`
		Cap1     []int     `json:"cap_1"`
		Turn     int       `json:"turn"`
	} `json:"data"`
}

// JSONをUSIに変換する
func Convert(b []byte) (s string, e error) {
	v, err := jason.NewObjectFromBytes(b)
	if err != nil {
		e = errors.New("jsonのパースに失敗 json=" + string(b) + "\n" + err.Error())
		return
	}

	// コマンドを一回動的にパースする
	cmd, err2 := v.GetString("command")
	if err2 != nil {
		e = errors.New("コマンドが指定されていません\n" + err2.Error())
		return
	}

	// コマンドによって再度パースする
	switch cmd {
	case "position":
		if p, err := toUsiPosition(b); err != nil {
			e = err
		} else {
			s = p
		}
	default:
		e = errors.New("不明なコマンドです command=" + cmd)
	}

	return
}

func toUsiPosition(b []byte) (string, error) {
	var p position
	if err := json.Unmarshal(b, &p); err != nil {
		return "", errors.New("Positionコマンドに変換できませんでした json=" + string(b) + "\n " + err.Error())
	}

	arr := make([]string, 9)
	for r, row := range p.Data.Position {
		usir, err := rowToUsi(row)
		if err != nil {
			return "", err
		}
		arr[r] = usir
	}

	s := "position sfen " + strings.Join(arr, "/")

	if p.Data.Turn == 0 {
		s += " b "
	} else {
		s += " w "
	}

	c0, c1 := p.Data.Cap0, p.Data.Cap1
	if len(c0) == 0 && len(c1) == 0 {
		return s + "- 1", nil
	}

	// TODO
	for i, c := range c0 {
		if c != 0 {
			p, err := pieceIdToUsi(i + 1)
			if err != nil {
				return "", err
			}
			s += strconv.Itoa(c) + p
		}
	}
	for i, c := range c1 {
		if c != 0 {
			p, err := pieceIdToUsi(-i - 1)
			if err != nil {
				return "", err
			}
			s += strconv.Itoa(c) + p
		}
	}
	return s + " 1", nil
}

func rowToUsi(r [9]int) (s string, e error) {
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
		p, err := pieceIdToUsi(id)
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

func pieceIdToUsi(i int) (s string, e error) {
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
		e = errors.New("pieceIdが不正です id=" + strconv.Itoa(i))
	}
	return
}
