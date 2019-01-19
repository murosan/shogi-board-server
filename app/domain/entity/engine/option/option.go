// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package option

import (
	"fmt"
	"strconv"

	"github.com/murosan/shogi-board-server/app/domain/exception"
	"github.com/murosan/shogi-board-server/app/lib/stringutil"
)

var (
	pref = "setoption name "
	val  = " value "
)

// Option 将棋エンジンのオプションのインターフェース
type Option interface {
	// オプション名を返す
	GetName() string

	// USI形式の文字列を返す
	Usi() string

	// 新しい値をオプションにセットして、更新されたUSI形式の文字列を返す
	// 新しい値が不正ならエラー
	Update(interface{}) (string, error)
}

// Button USI の button。値は持たない。on/off の機能はなく、
// 単に押すだけということ。
type Button struct {
	Name string `json:"name"`
}

// NewButton 新しい Button を返す
func NewButton(name string) *Button { return &Button{name} }

// GetName Button のオプション名を返す
func (b *Button) GetName() string { return b.Name }

// Usi USI の setoption コマンドを返す
func (b *Button) Usi() string { return pref + b.Name }

// Update Button を更新する。特に何もしない
func (b *Button) Update(_ interface{}) (string, error) {
	return b.Usi(), nil
}

// Check USI の check。bool(true|false) の値を持つ
type Check struct {
	Name    string `json:"name"`
	Val     bool   `json:"val"`
	Default bool   `json:"default"`
}

// NewCheck 新しい Check を返す
func NewCheck(name string, val, init bool) *Check {
	return &Check{name, val, init}
}

// GetName Check のオプション名を返す
func (c *Check) GetName() string { return c.Name }

// Usi USI の setoption コマンドを返す
func (c *Check) Usi() string {
	return pref + c.Name + val + strconv.FormatBool(c.Val)
}

// Update Check の値を更新する bool を受ける。bool 以外はエラー
func (c *Check) Update(i interface{}) (string, error) {
	b, ok := i.(bool)
	if !ok {
		return "", exception.InvalidOptionParameter.WithMsg("[check] Value type must be bool")
	}

	c.Val = b
	return c.Usi(), nil
}

// Spin USI の spin。int の値を持つ。初期値・最小値・最大値が必要。
// 値は必ずその範囲内にある必要がある。将棋エンジンからその値が渡されないと、ちゃんと動作しない
type Spin struct {
	Name    string `json:"name"`
	Val     int    `json:"val"`
	Default int    `json:"default"`
	Min     int    `json:"min"`
	Max     int    `json:"max"`
}

// NewSpin 新しい Spin を返す
func NewSpin(name string, val, init, min, max int) *Spin {
	return &Spin{name, val, init, min, max}
}

// GetName オプション名を返す
func (s *Spin) GetName() string { return s.Name }

// Usi USI の setoption コマンドを返す
func (s *Spin) Usi() string {
	return pref + s.Name + val + strconv.Itoa(s.Val)
}

// Update 値を更新する。 int 以外はエラー。値が範囲内にない場合もエラー
func (s *Spin) Update(i interface{}) (string, error) {
	fn, ok := i.(float64) // json を interface{} でパースすると float64 になってしまう
	if !ok {
		return "", exception.InvalidOptionParameter.WithMsg("[spin] Value type must be int.")
	}

	n := int(fn)
	if n < s.Min || n > s.Max {
		return "", exception.InvalidOptionParameter.WithMsg(fmt.Sprintf(
			"[spin] Value must be greater than or equal to %d, "+
				"and lesser than or equal to %d. But got: %d", s.Min, s.Max, n))
	}

	s.Val = n
	return s.Usi(), nil
}

// Select USI の combo。Vars が選択肢で、値(Val)や Default はそのVarsから選ぶ。
// Vars に存在しない値は持つことができない。
type Select struct {
	Name    string   `json:"name"`
	Val     string   `json:"val"`
	Default string   `json:"default"`
	Vars    []string `json:"vars"`
}

// NewSelect 新しい Select を返す
func NewSelect(name, val, init string, vars []string) *Select {
	return &Select{name, val, init, vars}
}

// GetName オプション名を返す
func (s *Select) GetName() string { return s.Name }

// Usi USI の setoption コマンドを返す
func (s *Select) Usi() string { return pref + s.Name + val + s.Val }

// Update 値(Val) を更新する。選択肢にない時や、文字列以外が渡るとエラー
func (s *Select) Update(i interface{}) (string, error) {
	v, ok := i.(string)
	if !ok {
		return "", exception.InvalidOptionParameter.WithMsg("[select] Value type must be string.")
	}

	if !stringutil.SliceContains(s.Vars, v) {
		return "", exception.InvalidOptionParameter.WithMsg(fmt.Sprintf(
			"[select] Value was not in vars."+
				"Value: %s\nVars: %v", v, s.Vars))
	}

	s.Val = v
	return s.Usi(), nil
}

// String USi の string。単純な文字列を値に持つ
type String struct {
	Name    string `json:"name"`
	Val     string `json:"val"`
	Default string `json:"default"`
}

// NewString 新しい String を返す
func NewString(name, val, init string) *String {
	return &String{name, val, init}
}

// GetName オプション名を返す
func (s *String) GetName() string { return s.Name }

// Usi USI の setoption コマンドを返す
func (s *String) Usi() string { return pref + s.Name + val + s.Val }

// Update 値(Val) を更新する。string 以外はエラー
func (s *String) Update(i interface{}) (string, error) {
	v, ok := i.(string)
	if !ok {
		return "", exception.InvalidOptionParameter.WithMsg("[string] Value type must be string.")
	}
	s.Val = v
	return s.Usi(), nil
}

// FileName USI の filename。ファイル名とあるが、実態は String と全く同じただの文字列である。
// 分ける意味が全くわからないが、USI でそうなっているので仕方ない。
type FileName struct {
	Name    string `json:"name"`
	Val     string `json:"val"`
	Default string `json:"default"`
}

// NewFileName 新しい FileName を返す
func NewFileName(name, val, init string) *FileName {
	return &FileName{name, val, init}
}

// GetName オプション名を返す
func (f *FileName) GetName() string { return f.Name }

// Usi USI の setoption コマンドを返す
func (f *FileName) Usi() string { return pref + f.Name + val + f.Val }

// Update 値(Val) を更新する。string 以外はエラー
func (f *FileName) Update(i interface{}) (string, error) {
	v, ok := i.(string)
	if !ok {
		return "", exception.InvalidOptionParameter.WithMsg("[filename] Value type must be string.")
	}
	f.Val = v
	return f.Usi(), nil
}
