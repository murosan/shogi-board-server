// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package option

import (
	"fmt"
	"github.com/murosan/shogi-proxy-server/app/domain/entity/exception"
	"github.com/murosan/shogi-proxy-server/app/service/logger"
	"go.uber.org/zap"
)

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

// 新しいオプションを追加する
func (om *OptMap) Append(o Option) {
	switch t := o.(type) {
	case *Button:
		om.Buttons[t.GetName()] = t
	case *Check:
		om.Checks[t.GetName()] = t
	case *Spin:
		om.Spins[t.GetName()] = t
	case *Select:
		om.Combos[t.GetName()] = t
	case *String:
		om.Strings[t.GetName()] = t
	case *FileName:
		om.FileNames[t.GetName()] = t
	default:
		panic(exception.UnknownOption)
	}
}

// TODO: オプションの名前をまとめて変数から使うとか整理する
func (om *OptMap) Update(v UpdateOptionValue) (string, error) {
	var (
		opt Option
		ok  bool
	)
	switch v.Type {
	case "button":
		opt, ok = om.Buttons[v.Name]
	case "check":
		opt, ok = om.Checks[v.Name]
	case "spin":
		opt, ok = om.Spins[v.Name]
	case "select":
		opt, ok = om.Combos[v.Name]
	case "string":
		opt, ok = om.Strings[v.Name]
	case "filename":
		opt, ok = om.FileNames[v.Name]
	}

	logger.Use().Debug(
		"OptMap_Update",
		zap.Any("SpecifiedOption", opt),
		zap.Bool("Ok", ok),
	)
	if ok {
		return opt.Update(v.Value)
	}

	msg := fmt.Sprintf("OptionName %s was not found.", v.Name)
	logger.Use().Error("OptMap_Update. Type was not valid", zap.String("msg", msg))
	return "", exception.UnknownOption.WithMsg(msg)
}
