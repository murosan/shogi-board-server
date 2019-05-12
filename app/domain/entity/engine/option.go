// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package engine

import (
	"fmt"
	"strconv"

	sslice "github.com/murosan/goutils/slice/strings"
)

const (
	// These are the parts of USI command.
	set = "setoption name "
	val = " value "
)

// Option is a shogi enigne option.
// Option is given when initializing.
type Option interface {
	// ToUSI returns
	ToUSI() string
}

// Button is a Button type in USI expression.
type Button struct {
	Name string `json:"name"`
}

// ToUSI returns USI setoption command string.
func (b *Button) ToUSI() string {
	return set + b.Name
}

// Check is a Check type in USI expression.
type Check struct {
	Name    string `json:"name"`
	Value   bool   `json:"value"`
	Default bool   `json:"default"`
}

// ToUSI returns USI setoption command string.
func (c *Check) ToUSI() string {
	return set + c.Name + val + strconv.FormatBool(c.Value)
}

// Set sets the given value to Check.Value
func (c *Check) Set(b bool) { c.Value = b }

// Range is a Spin type in USI expression.
type Range struct {
	Name    string `json:"name"`
	Value   int    `json:"value"`
	Default int    `json:"default"`
	Min     int    `json:"min"`
	Max     int    `json:"max"`
}

// ToUSI returns USI setoption command string.
func (r *Range) ToUSI() string {
	return set + r.Name + val + fmt.Sprint(r.Value)
}

// Set sets given value to Range.Value.
// Returns an error if given value was out of range, otherwise returns nil.
func (r *Range) Set(i int) error {
	if i < r.Min || i > r.Max {
		f := "[Check.Set] Given value is out of range. Given=%d, Min=%d, Max=%d"
		return fmt.Errorf(f, i, r.Min, r.Max)
	}

	r.Value = i
	return nil
}

// Select is a Combo type in USI expression.
type Select struct {
	Name    string   `json:"name"`
	Value   string   `json:"value"`
	Default string   `json:"default"`
	Vars    []string `json:"vars"`
}

// ToUSI returns USI setoption command string.
func (s *Select) ToUSI() string {
	return set + s.Name + val + s.Value
}

// Set sets given value to Select.Value.
// Returns an error if given value was not in Set.Vars, otherwise returns nil.
func (s *Select) Set(v string) error {
	if sslice.NotContain(s.Vars, v) {
		f := "[Select.Set] Given string is not in vars. Given=%s, Vars=%v"
		return fmt.Errorf(f, v, s.Vars)
	}

	s.Value = v
	return nil
}

// Text is a String or Filename type in USI expression.
// We treat the String and Filename type as a single type named Text.
type Text struct {
	Name    string `json:"name"`
	Value   string `json:"value"`
	Default string `json:"default"`
}

// ToUSI returns USI setoption command string.
func (t *Text) ToUSI() string {
	return set + t.Name + val + t.Value
}

// Set sets given value to Text.Value
func (t *Text) Set(s string) {
	t.Value = s
}
