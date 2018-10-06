// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package usi

import (
	"encoding/json"
	"github.com/antonholmquist/jason"
	"strconv"
	"strings"
)

type Position struct {
	Version uint8  `json:"version"`
	Command string `json:"command"`
	Data    struct {
		Position [9][9]int `json:"position"`
		Cap0     []int     `json:"cap_0"`
		Cap1     []int     `json:"cap_1"`
		Turn     int       `json:"turn"`
	} `json:"data"`
}

// JSONをusiに変換する
func Convert(b []byte) string {
	v, err := jason.NewObjectFromBytes(b)
	if err != nil {
		panic("jsonのパースに失敗 json=" + string(b) + "\n" + err.Error())
	}

	cmd, err2 := v.GetString("command")
	if err2 != nil {
		panic("コマンドが指定されていません\n" + err2.Error())
	}

	switch cmd {
	case "position":
		return toUsiPosition(b)
	default:
		panic("不明なコマンドです command=" + cmd)
	}
}

func toUsiPosition(b []byte) string {
	var p Position
	if err := json.Unmarshal(b, &p); err != nil {
		panic("Positionコマンドに変換できませんでした json=" + string(b) + "\n " + err.Error())
	}

	arr := make([]string, 9)
	for r, row := range p.Data.Position {
		arr[r] = rowToUsi(row)
	}

	s := strings.Join(arr, "/")

	if p.Data.Turn == 0 {
		s += " b "
	} else {
		s += " w "
	}

	c0, c1 := p.Data.Cap0, p.Data.Cap1
	if len(c0) == 0 && len(c1) == 0 {
		return s + "- 1"
	}

	for i, c := range c0 {
		if c != 0 {
			s += strconv.Itoa(c) + pieceIdToUsi(i+1)
		}
	}
	for i, c := range c1 {
		if c != 0 {
			s += strconv.Itoa(c) + pieceIdToUsi(-i-1)
		}
	}
	return s + " 1"
}

func rowToUsi(r [9]int) (s string) {
	emp := 0
	for _, id := range r {
		if id == 0 {
			emp++
			continue
		}
		if emp != 0 {
			s += strconv.Itoa(emp)
			emp = 0
		}
		s += pieceIdToUsi(id)
	}
	if emp != 0 {
		s += strconv.Itoa(emp)
	}
	return
}

func pieceIdToUsi(i int) (s string) {
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
		panic("pieceIdが不正です id=" + strconv.Itoa(i))
	}
	return
}
