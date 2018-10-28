// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package option

import (
	"bytes"
	"strconv"
)

var (
	space = []byte(" ")
	val   = []byte("value")
	pref  = []byte("setoption name")
)

type Option interface {
	Usi() []byte
	GetName() []byte
}

type Button struct {
	Name []byte
}

func (b Button) Usi() []byte {
	return bytes.Join([][]byte{pref, b.Name}, space)
}

func (b Button) GetName() []byte { return b.Name }

type Check struct {
	Name         []byte
	Val, Default bool
}

func (c Check) Usi() []byte {
	b := []byte(strconv.FormatBool(c.Val))
	return bytes.Join([][]byte{pref, c.Name, val, b}, space)
}

func (c Check) GetName() []byte { return c.Name }

type Spin struct {
	Name                   []byte
	Val, Default, Min, Max int
}

func (s Spin) Usi() []byte {
	b := strconv.Itoa(s.Val)
	return bytes.Join([][]byte{pref, s.Name, val, []byte(b)}, space)
}

func (s Spin) GetName() []byte { return s.Name }

// USIのcombo
type Select struct {
	Name  []byte
	Index int
	Vars  [][]byte
}

func (s Select) Usi() []byte {
	return bytes.Join([][]byte{pref, s.Name, val, s.Vars[s.Index]}, space)
}

func (s Select) GetName() []byte { return s.Name }

type String struct {
	Name, Val, Default []byte
}

func (s String) Usi() []byte {
	return bytes.Join([][]byte{pref, s.Name, val, s.Val}, space)
}

func (s String) GetName() []byte { return s.Name }

type FileName struct {
	Name, Val, Default []byte
}

func (f FileName) Usi() []byte {
	return bytes.Join([][]byte{pref, f.Name, val, f.Val}, space)
}

func (f FileName) GetName() []byte { return f.Name }
