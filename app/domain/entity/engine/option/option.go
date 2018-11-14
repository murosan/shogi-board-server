// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package option

import (
	"fmt"
	"strconv"
)

var (
	val  = "value"
	pref = "setoption name"
)

type Option interface {
	UpdateAndGetUsi() (string, error)
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

func (b Button) UpdateAndGetUsi() (string, error) {
	return fmt.Sprintf("%s %s", pref, b.Name), nil
}

func (b Button) GetName() string { return b.Name }

type Check struct {
	Name    string `json:"name"`
	Val     bool   `json:"val"`
	Default bool   `json:"default"`
}

func (c Check) UpdateAndGetUsi() (string, error) {
	b := []byte(strconv.FormatBool(c.Val))
	return fmt.Sprintf("%s %s %s %s", pref, c.Name, val, b), nil
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

func (s Spin) UpdateAndGetUsi() (string, error) {
	b := strconv.Itoa(s.Val)
	return fmt.Sprintf("%s %s %s %s", pref, s.Name, val, b), nil
}

func (s Spin) GetName() string { return s.Name }

// USIのcombo
type Select struct {
	Name  string   `json:"name"`
	Index int      `json:"index"`
	Vars  []string `json:"vars"`
}

func (s Select) UpdateAndGetUsi() (string, error) {
	return fmt.Sprintf("%s %s %s %s", pref, s.Name, val, s.Vars[s.Index]), nil
}

func (s Select) GetName() string { return s.Name }

type String struct {
	Name    string `json:"name"`
	Val     string `json:"val"`
	Default string `json:"default"`
}

func (s String) UpdateAndGetUsi() (string, error) {
	return fmt.Sprintf("%s %s %s %s", pref, s.Name, val, s.Val), nil
}

func (s String) GetName() string { return s.Name }

type FileName struct {
	Name    string `json:"name"`
	Val     string `json:"val"`
	Default string `json:"default"`
}

func (f FileName) UpdateAndGetUsi() (string, error) {
	return fmt.Sprintf("%s %s %s %s", pref, f.Name, val, f.Val), nil
}

func (f FileName) GetName() string { return f.Name }
