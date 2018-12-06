// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package from_usi

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/murosan/shogi-proxy-server/app/domain/entity/engine/result"
	"github.com/murosan/shogi-proxy-server/app/domain/entity/shogi"
	"github.com/murosan/shogi-proxy-server/app/domain/exception"
)

// Info をパース(info string は渡さない)
// return
//   r *result.Info パースした結果。失敗したら nil
//   mpv int multipvならその値。multipvじゃなければ 0
//   err error エラー
func (fu *FromUsi) Info(s string) (r *result.Info, mpv int, err error) {
	a := strings.Split(s, " ")
	r = result.NewInfo()

	// panic をリカバーしてエラーをセット
	defer func() {
		if rec := recover(); rec != nil {
			r = nil
			err = exception.FailedToParseInfo.WithMsg(fmt.Sprintf("%v", rec))
		}
	}()

	i := 0
	for i < len(a) {
		switch strings.TrimSpace(a[i]) {
		case result.Depth:
			i++
			r.Values[result.Depth] = toInt(a[i])
		case result.SelDepth:
			i++
			r.Values[result.SelDepth] = toInt(a[i])
		case result.Time:
			i++
			r.Values[result.Time] = toInt(a[i])
		case result.Nodes:
			i++
			r.Values[result.Nodes] = toInt(a[i])
		case result.HashFull:
			i++
			r.Values[result.HashFull] = toInt(a[i])
		case result.Nps:
			i++
			r.Values[result.Nps] = toInt(a[i])
		case result.Score:
			if a[i+1] == "cp" || a[i+1] == "mate" {
				r.Score = toInt(a[i+2])
			}
			i += 2
		case result.MultiPv:
			i++
			mpv = toInt(a[i])
		case result.Pv:
			fu.setMoves(r, a[i+1:])
			i += len(a) // pv は 末尾
		}
		i++
	}

	return
}

func (fu *FromUsi) setMoves(r *result.Info, a []string) {
	m := make([]shogi.Move, len(a))
	for i, v := range a {
		mv, err := fu.Move(v)
		if err != nil {
			panic(err)
		}
		m[i] = *mv
	}

	r.Moves = m
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
