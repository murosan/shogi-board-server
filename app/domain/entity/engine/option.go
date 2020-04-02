// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package engine

import (
	"fmt"
	"strconv"
)

const (
	set = "setoption name "
	val = " value "
)

// Option represents engine option given on initializing engine.
type Option interface {
	fmt.Stringer
	ToUSI() string   // ToUSI returns USI setoption command string.
	Validate() error // Validate validates and return true if valid.
}

// Button is a Button type in USI expression.
type Button struct {
	Name string `json:"name"`
}

func (b *Button) ToUSI() string   { return set + b.Name }
func (b *Button) Validate() error { return nil }
func (b *Button) String() string  { return fmt.Sprintf("Button{name:%s}", b.Name) }
func (b *Button) Copy() *Button   { return &Button{Name: b.Name} }

// Check is a Check type in USI expression.
type Check struct {
	Name    string `json:"name"`
	Value   bool   `json:"value"`
	Default bool   `json:"default"`
}

func (c *Check) ToUSI() string   { return set + c.Name + val + strconv.FormatBool(c.Value) }
func (c *Check) Validate() error { return nil }
func (c *Check) String() string {
	return fmt.Sprintf(
		"Check{name:%s,value:%t,default:%t}",
		c.Name, c.Value, c.Default,
	)
}
func (c *Check) Copy() *Check {
	return &Check{
		Name:    c.Name,
		Value:   c.Value,
		Default: c.Default,
	}
}

// Range is a Spin type in USI expression.
type Range struct {
	Name    string `json:"name"`
	Value   int    `json:"value"`
	Default int    `json:"default"`
	Min     int    `json:"min"`
	Max     int    `json:"max"`
}

func (r *Range) ToUSI() string { return set + r.Name + val + fmt.Sprint(r.Value) }
func (r *Range) Validate() error {
	if r.Value < r.Min || r.Value > r.Max {
		return fmt.Errorf(
			"[Check.Validate] out of range. value=%d, min=%d, max=%d",
			r.Value, r.Min, r.Max,
		)
	}
	return nil
}
func (r *Range) String() string {
	return fmt.Sprintf(
		"Range{name:%s,value:%d,default:%d,min:%d,max:%d}",
		r.Name, r.Value, r.Default, r.Min, r.Max,
	)
}
func (r *Range) Copy() *Range {
	return &Range{
		Name:    r.Name,
		Value:   r.Value,
		Default: r.Default,
		Min:     r.Min,
		Max:     r.Max,
	}
}

// Select is a Combo type in USI expression.
type Select struct {
	Name    string   `json:"name"`
	Value   string   `json:"value"`
	Default string   `json:"default"`
	Vars    []string `json:"vars"`
}

func (s *Select) ToUSI() string { return set + s.Name + val + s.Value }
func (s *Select) Validate() error {
	for _, val := range s.Vars {
		if val == s.Value {
			return nil
		}
	}
	return fmt.Errorf(
		"[Select.Validate] vars does not contain the value. value=%s, vars=%v",
		s.Value, s.Vars,
	)
}
func (s *Select) String() string {
	return fmt.Sprintf(
		"Select{name:%s,value:%s,default:%s,vars:%v}",
		s.Name, s.Value, s.Default, s.Vars,
	)
}
func (s *Select) Copy() *Select {
	return &Select{
		Name:    s.Name,
		Value:   s.Value,
		Default: s.Default,
		Vars:    s.Vars,
	}
}

// Text is a String or Filename type in USI expression.
// We treat both String and Filename type as a single 'Text' type.
type Text struct {
	Name    string `json:"name"`
	Value   string `json:"value"`
	Default string `json:"default"`
}

func (t *Text) ToUSI() string   { return set + t.Name + val + t.Value }
func (t *Text) Validate() error { return nil }
func (t *Text) String() string {
	return fmt.Sprintf(
		"Text{name:%s,value:%s,default:%s}",
		t.Name, t.Value, t.Default,
	)
}
func (t *Text) Copy() *Text {
	return &Text{
		Name:    t.Name,
		Value:   t.Value,
		Default: t.Default,
	}
}
