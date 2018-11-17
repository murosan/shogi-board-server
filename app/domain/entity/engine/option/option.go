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

// json化するために Name を大文字始まりにしているが、基本他からは触らない
type Button struct {
	Name string `json:"name"`
}

func NewButton(name string) *Button { return &Button{name} }

func (b *Button) Usi() (string, error) {
	return pref + b.Name, nil
}

func (b *Button) GetName() string { return b.Name }

type Check struct {
	Name    string `json:"name"`
	Val     bool   `json:"val"`
	Default bool   `json:"default"`
}

func NewCheck(name string, val, init bool) *Check {
	return &Check{name, val, init}
}

func (c *Check) Usi() (string, error) {
	b := strconv.FormatBool(c.Val)
	return pref + c.Name + val + b, nil
}

func (c *Check) GetName() string { return c.Name }

type Spin struct {
	Name    string `json:"name"`
	Val     int    `json:"val"`
	Default int    `json:"default"`
	Min     int    `json:"min"`
	Max     int    `json:"max"`
}

func NewSpin(name string, val, init, min, max int) *Spin {
	return &Spin{name, val, init, min, max}
}

func (s *Spin) Usi() (string, error) {
	b := strconv.Itoa(s.Val)
	return pref + s.Name + val + b, nil
}

func (s *Spin) GetName() string { return s.Name }

// USIのcombo
type Select struct {
	Name    string   `json:"name"`
	Val     string   `json:"val"`
	Default string   `json:"default"`
	Vars    []string `json:"vars"`
}

func NewSelect(name, val, init string, vars []string) *Select {
	return &Select{name, val, init, vars}
}

func (s *Select) Usi() (string, error) {
	return pref + s.Name + val + s.Val, nil
}

func (s *Select) GetName() string { return s.Name }

type String struct {
	Name    string `json:"name"`
	Val     string `json:"val"`
	Default string `json:"default"`
}

func NewString(name, val, init string) *String {
	return &String{name, val, init}
}

func (s *String) Usi() (string, error) {
	return pref + s.Name + val + s.Val, nil
}

func (s *String) GetName() string { return s.Name }

type FileName struct {
	Name    string `json:"name"`
	Val     string `json:"val"`
	Default string `json:"default"`
}

func NewFileName(name, val, init string) *FileName {
	return &FileName{name, val, init}
}

func (f *FileName) Usi() (string, error) {
	return pref + f.Name + val + f.Val, nil
}

func (f *FileName) GetName() string { return f.Name }
