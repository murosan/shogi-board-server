// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package option

import (
	"fmt"
	"strconv"

	"github.com/murosan/shogi-proxy-server/app/domain/entity/exception"
	"github.com/murosan/shogi-proxy-server/app/lib/stringutil"
)

var (
	pref = "setoption name "
	val  = " value "
)

type Option interface {
	// オプション名を返す
	GetName() string

	// USI形式の文字列を返す
	Usi() string

	// 新しい値をオプションにセットして、更新されたUSI形式の文字列を返す
	// 新しい値が不正ならエラー
	Update(interface{}) (string, error)
}

// json化するために Name を大文字始まりにしているが、基本他からは触らない
type Button struct {
	Name string `json:"name"`
}

func NewButton(name string) *Button { return &Button{name} }
func (b *Button) GetName() string   { return b.Name }
func (b *Button) Usi() string       { return pref + b.Name }
func (b *Button) Update(_ interface{}) (string, error) {
	return b.Usi(), nil
}

type Check struct {
	Name    string `json:"name"`
	Val     bool   `json:"val"`
	Default bool   `json:"default"`
}

func NewCheck(name string, val, init bool) *Check {
	return &Check{name, val, init}
}
func (c *Check) GetName() string { return c.Name }
func (c *Check) Usi() string {
	return pref + c.Name + val + strconv.FormatBool(c.Val)
}
func (c *Check) Update(i interface{}) (string, error) {
	switch b := i.(type) {
	case bool:
		c.Val = b
		return c.Usi(), nil
	}
	return "", exception.InvalidOptionParameter
}

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
func (s *Spin) GetName() string { return s.Name }
func (s *Spin) Usi() string {
	return pref + s.Name + val + strconv.Itoa(s.Val)
}
func (s *Spin) Update(i interface{}) (string, error) {
	fn, ok := i.(float64)
	if !ok {
		return "", exception.InvalidOptionParameter.WithMsg("[spin] Value type must be int.")
	}

	n := int(fn)
	if s.Min <= n && n <= s.Max {
		s.Val = n
		return s.Usi(), nil
	}
	format := "[spin] Value must be greater than or equal to %d" +
		"and lesser than or equal to %d. But got: %d"
	msg := fmt.Sprintf(format, s.Min, s.Max, n)
	return "", exception.InvalidOptionParameter.WithMsg(msg)
}

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
func (s *Select) GetName() string { return s.Name }
func (s *Select) Usi() string     { return pref + s.Name + val + s.Val }
func (s *Select) Update(i interface{}) (string, error) {
	switch v := i.(type) {
	case string:
		if stringutil.SliceContains(s.Vars, v) {
			s.Val = v
			return s.Usi(), nil
		}
	}
	return "", exception.InvalidOptionParameter
}

type String struct {
	Name    string `json:"name"`
	Val     string `json:"val"`
	Default string `json:"default"`
}

func NewString(name, val, init string) *String {
	return &String{name, val, init}
}
func (s *String) GetName() string { return s.Name }
func (s *String) Usi() string     { return pref + s.Name + val + s.Val }
func (s *String) Update(i interface{}) (string, error) {
	switch v := i.(type) {
	case string:
		s.Val = v
		return s.Usi(), nil
	}
	return "", exception.InvalidOptionParameter
}

type FileName struct {
	Name    string `json:"name"`
	Val     string `json:"val"`
	Default string `json:"default"`
}

func NewFileName(name, val, init string) *FileName {
	return &FileName{name, val, init}
}
func (f *FileName) GetName() string { return f.Name }
func (f *FileName) Usi() string     { return pref + f.Name + val + f.Val }
func (f *FileName) Update(i interface{}) (string, error) {
	switch v := i.(type) {
	case string:
		f.Val = v
		return f.Usi(), nil
	}
	return "", exception.InvalidOptionParameter
}
