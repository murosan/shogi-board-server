package engine

import (
	"bytes"
	"strconv"
)

// 将棋エンジンが出力するオプションを定義する
// 数値はとりあえず int にしておく
// 文字列は string より []byte の方が都合がいいので
// []byte で統一しておく
// TODO: bytes.Join のパフォーマンスを調べる

var (
	pref = []byte("setoption name")
	val  = []byte("value")
)

type Option interface {
	// USIコマンドを返す
	Usi() []byte
}

type Check struct {
	Name         []byte
	Val, Default bool
}

func (c Check) Usi() []byte {
	b := []byte(strconv.FormatBool(c.Val))
	return bytes.Join([][]byte{pref, c.Name, val, b}, space)
}

type Spin struct {
	Name                   []byte
	Val, Default, Min, Max int
}

func (s Spin) Usi() []byte {
	b := strconv.Itoa(s.Val)
	return bytes.Join([][]byte{pref, s.Name, val, []byte(b)}, space)
}

// USIのcombo
type Select struct {
	Name  []byte
	Index int
	Vars  [][]byte
}

func (c Select) Usi() []byte {
	return bytes.Join([][]byte{pref, c.Name, val, c.Vars[c.Index]}, space)
}

type Button struct {
	Name []byte
}

func (b Button) Usi() []byte {
	return bytes.Join([][]byte{pref, b.Name}, space)
}

type String struct {
	Name, Val, Default []byte
}

func (s String) Usi() []byte {
	return bytes.Join([][]byte{pref, s.Name, val, s.Val}, space)
}

type FileName struct {
	Name, Val, Default []byte
}

func (f FileName) Usi() []byte {
	return bytes.Join([][]byte{pref, f.Name, val, f.Val}, space)
}
