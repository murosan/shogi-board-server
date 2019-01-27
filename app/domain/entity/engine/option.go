// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package engine

import (
	"fmt"
	"strconv"

	"github.com/murosan/shogi-board-server/app/domain/exception"
	"github.com/murosan/shogi-board-server/app/lib/stringutil"
	pb "github.com/murosan/shogi-board-server/app/proto"
)

var (
	// USI コマンドのパーツ
	pref = "setoption name "
	val  = " value "
)

// AppendOption は Options に指定のオプションを追加します
func AppendOption(opts *pb.Options, opt interface{}) {
	switch o := opt.(type) {
	case *pb.Button:
		opts.Buttons[o.Name] = o
	case *pb.Check:
		opts.Checks[o.Name] = o
	case *pb.Spin:
		opts.Spins[o.Name] = o
	case *pb.Select:
		opts.Selects[o.Name] = o
	case *pb.String:
		opts.Strings[o.Name] = o
	case *pb.Filename:
		opts.Filenames[o.Name] = o
	default:
		panic(exception.UnknownOptionType)
	}
}

// UpdateCheck は Options の Check を更新します
// panics
//   - Name が Options に存在しない
func UpdateCheck(opts *pb.Options, c *pb.Check) {
	if _, ok := opts.Checks[c.Name]; !ok {
		panic(exception.UnknownOption)
	}
	opts.Checks[c.Name] = c
}

// UpdateSpin は Options の Spin を更新します
// panics
//   - Name が Options に存在しない
//   - Min, Max が元の値と違う
//   - Val が範囲外 (Val < Min || Val > Max)
func UpdateSpin(opts *pb.Options, s *pb.Spin) {
	sp, ok := opts.Spins[s.Name]
	if !ok {
		panic(exception.UnknownOption)
	}
	if sp.Min != s.Min || sp.Max != s.Max || s.Val < s.Min || s.Val > s.Max {
		panic(exception.InvalidOptionParameter)
	}

	opts.Spins[s.Name] = s
}

// UpdateSelect は Options の Select を更新します
// panics
//   - Name が Options に存在しない
//   - Vars と Default が元の値と異なる
//   - Val が Vars に存在しない
func UpdateSelect(opts *pb.Options, s *pb.Select) {
	se, ok := opts.Selects[s.Name]
	if !ok {
		panic(exception.UnknownOption)
	}
	if !stringutil.SliceEquals(se.Vars, s.Vars) || se.Default != s.Default {
		panic(exception.InvalidOptionParameter)
	}
	if !stringutil.SliceContains(s.Vars, s.Val) {
		panic(exception.InvalidOptionParameter)
	}

	opts.Selects[s.Name] = s
}

// UpdateString は Options の String を更新します
// panics
//   - Name が Options に存在しない
func UpdateString(opts *pb.Options, s *pb.String) {
	if _, ok := opts.Strings[s.Name]; !ok {
		panic(exception.UnknownOption)
	}

	opts.Strings[s.Name] = s
}

// UpdateFilename は Options の Filename を更新します
// panics
//   - Name が Options に存在しない
func UpdateFilename(opts *pb.Options, s *pb.Filename) {
	if _, ok := opts.Filenames[s.Name]; !ok {
		panic(exception.UnknownOption)
	}

	opts.Filenames[s.Name] = s
}

// ButtonUSI Button の setoption USI コマンドを返す
func ButtonUSI(b *pb.Button) string {
	return pref + b.Name
}

// CheckUSI Check の setoption USI コマンドを返す
func CheckUSI(c *pb.Check) string {
	return pref + c.Name + val + strconv.FormatBool(c.Val)
}

// SpinUSI Spin の setoption USI コマンドを返す
func SpinUSI(s *pb.Spin) string {
	return pref + s.Name + val + fmt.Sprint(s.Val)
}

// SelectUSI Select の setoption USI コマンドを返す
func SelectUSI(s *pb.Select) string {
	return pref + s.Name + val + s.Val
}

// StringUSI String の setoption USI コマンドを返す
func StringUSI(s *pb.String) string {
	return pref + s.Name + val + s.Val
}

// FilenameUSI Filename の setoption USI コマンドを返す
func FilenameUSI(s *pb.Filename) string {
	return pref + s.Name + val + s.Val
}
