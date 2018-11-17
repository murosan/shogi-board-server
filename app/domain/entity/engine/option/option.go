// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package option

import (
	"strconv"
)

var (
	val  = " value "
	pref = "setoption name "
)

type Option interface {
	Usi() (string, error)
	GetName() string
}

type OptMap struct {
	Buttons   map[string]*Button   `json:"buttons"`
	Checks    map[string]*Check    `json:"checks"`
	Spins     map[string]*Spin     `json:"spins"`
	Combos    map[string]*Select   `json:"combos"`
	Strings   map[string]*String   `json:"strings"`
	FileNames map[string]*FileName `json:"file_names"`
}

func EmptyOptMap() *OptMap {
	return &OptMap{
		Buttons:   make(map[string]*Button),
		Checks:    make(map[string]*Check),
		Spins:     make(map[string]*Spin),
		Combos:    make(map[string]*Select),
		Strings:   make(map[string]*String),
		FileNames: make(map[string]*FileName),
	}
}

type Button struct{ name string }

func NewButton(name string) *Button { return &Button{name} }

func (b *Button) Usi() (string, error) {
	return pref + b.name, nil
}

func (b *Button) GetName() string { return b.name }

type Check struct {
	name         string
	val, initial bool
}

func NewCheck(name string, val, init bool) *Check {
	return &Check{name, val, init}
}

func (c *Check) Usi() (string, error) {
	b := strconv.FormatBool(c.val)
	return pref + c.name + val + b, nil
}

func (c *Check) GetName() string { return c.name }

type Spin struct {
	name                   string
	val, initial, min, max int
}

func NewSpin(name string, val, init, min, max int) *Spin {
	return &Spin{name, val, init, min, max}
}

func (s *Spin) Usi() (string, error) {
	b := strconv.Itoa(s.val)
	return pref + s.name + val + b, nil
}

func (s *Spin) GetName() string { return s.name }

// USIのcombo
type Select struct {
	name, val, initial string
	vars               []string
}

func NewSelect(name, val, init string, vars []string) *Select {
	return &Select{name, val, init, vars}
}

func (s *Select) Usi() (string, error) {
	return pref + s.name + val + s.val, nil
}

func (s *Select) GetName() string { return s.name }

type String struct{ name, val, initial string }

func NewString(name, val, init string) *String {
	return &String{name, val, init}
}

func (s *String) Usi() (string, error) {
	return pref + s.name + val + s.val, nil
}

func (s *String) GetName() string { return s.name }

type FileName struct{ name, val, initial string }

func NewFileName(name, val, init string) *FileName {
	return &FileName{name, val, init}
}

func (f *FileName) Usi() (string, error) {
	return pref + f.name + val + f.val, nil
}

func (f *FileName) GetName() string { return f.name }
