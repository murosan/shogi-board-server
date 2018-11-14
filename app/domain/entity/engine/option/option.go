// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package option

import (
	"fmt"
	"strconv"

	"github.com/murosan/shogi-proxy-server/app/domain/entity/exception"
)

var (
	val  = "value"
	pref = "setoption name"
)

type Option interface {
	Usi() []byte
	GetName() string
}

type OptMap struct {
	Buttons   map[string]Button   `json:"buttons"`
	Checks    map[string]Check    `json:"checks"`
	Spins     map[string]Spin     `json:"spins"`
	Combos    map[string]Select   `json:"combos"`
	Strings   map[string]String   `json:"strings"`
	FileNames map[string]FileName `json:"file_names"`
}

func EmptyOptMap() *OptMap {
	return &OptMap{
		Buttons:   make(map[string]Button),
		Checks:    make(map[string]Check),
		Spins:     make(map[string]Spin),
		Combos:    make(map[string]Select),
		Strings:   make(map[string]String),
		FileNames: make(map[string]FileName),
	}
}

type Button struct {
	Name string `json:"name"`
}

func (b Button) Usi() []byte {
	return []byte(fmt.Sprintf("%s %s", pref, b.Name))
}

func (b Button) GetName() string { return b.Name }

type Check struct {
	Name    string `json:"name"`
	Val     bool   `json:"val"`
	Default bool   `json:"default"`
}

func (c Check) Usi() []byte {
	b := []byte(strconv.FormatBool(c.Val))
	return []byte(fmt.Sprintf("%s %s %s %s", pref, c.Name, val, b))
}

func (c Check) GetName() string { return c.Name }

func (c *Check) Update(b bool) { c.Val = b }

type Spin struct {
	Name    string `json:"name"`
	Val     int    `json:"val"`
	Default int    `json:"default"`
	Min     int    `json:"min"`
	Max     int    `json:"max"`
}

func (s Spin) Usi() []byte {
	b := strconv.Itoa(s.Val)
	return []byte(fmt.Sprintf("%s %s %s %s", pref, s.Name, val, b))
}

func (s Spin) GetName() string { return s.Name }

func (s *Spin) Update(i int) error {
	if i < s.Min || i > s.Max {
		return exception.InvalidOptionParameter
	}
	s.Val = i
	return nil
}

// USIのcombo
type Select struct {
	Name  string   `json:"name"`
	Index int      `json:"index"`
	Vars  []string `json:"vars"`
}

func (s Select) Usi() []byte {
	return []byte(fmt.Sprintf("%s %s %s %s", pref, s.Name, val, s.Vars[s.Index]))
}

func (s Select) GetName() string { return s.Name }

func (s *Select) Update(v string) error {
	// TODO: Indexで持つのはよくないかもしれない
	return nil
}

type String struct {
	Name    string `json:"name"`
	Val     string `json:"val"`
	Default string `json:"default"`
}

func (s String) Usi() []byte {
	return []byte(fmt.Sprintf("%s %s %s %s", pref, s.Name, val, s.Val))
}

func (s String) GetName() string { return s.Name }

func (s *String) Update(v string) { s.Val = v }

type FileName struct {
	Name    string `json:"name"`
	Val     string `json:"val"`
	Default string `json:"default"`
}

func (f FileName) Usi() []byte {
	return []byte(fmt.Sprintf("%s %s %s %s", pref, f.Name, val, f.Val))
}

func (f FileName) GetName() string { return f.Name }

func (f *FileName) Update(v string) { f.Val = v }
