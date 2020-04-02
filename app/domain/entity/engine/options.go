// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package engine

import (
	"fmt"
	"sync"
)

// Options is a option holder.
type Options struct {
	sync.RWMutex
	Buttons map[string]*Button `json:"buttons"`
	Checks  map[string]*Check  `json:"checks"`
	Ranges  map[string]*Range  `json:"ranges"`
	Selects map[string]*Select `json:"selects"`
	Texts   map[string]*Text   `json:"texts"`
}

// NewOptions returns new Options
func NewOptions() *Options {
	return &Options{
		Buttons: make(map[string]*Button),
		Checks:  make(map[string]*Check),
		Ranges:  make(map[string]*Range),
		Selects: make(map[string]*Select),
		Texts:   make(map[string]*Text),
	}
}

func (o *Options) String() string {
	return fmt.Sprintf(
		"Options{buttons:%v,checks:%v,ranges:%v,selects:%v,texts:%v}",
		o.Buttons, o.Checks, o.Ranges, o.Selects, o.Texts,
	)
}

// PutButton put the Button to Buttons map
func (o *Options) PutButton(b *Button) {
	o.Lock()
	o.Buttons[b.Name] = b
	o.Unlock()
}

// GetButton finds and returns a Button from the given name
func (o *Options) GetButton(name string) (*Button, bool) {
	o.RLock()
	b, ok := o.Buttons[name]
	o.RUnlock()
	return b, ok
}

// PutCheck put the Check to Checks map
func (o *Options) PutCheck(b *Check) {
	o.Lock()
	o.Checks[b.Name] = b
	o.Unlock()
}

// GetCheck finds and returns a Check from the given name
func (o *Options) GetCheck(name string) (*Check, bool) {
	o.RLock()
	b, ok := o.Checks[name]
	o.RUnlock()
	return b, ok
}

// PutRange put the Range to Ranges map
func (o *Options) PutRange(b *Range) {
	o.Lock()
	o.Ranges[b.Name] = b
	o.Unlock()
}

// GetRange finds and returns a Range from the given name
func (o *Options) GetRange(name string) (*Range, bool) {
	o.RLock()
	b, ok := o.Ranges[name]
	o.RUnlock()
	return b, ok
}

// PutSelect put the Select to Selects map
func (o *Options) PutSelect(b *Select) {
	o.Lock()
	o.Selects[b.Name] = b
	o.Unlock()
}

// GetSelect finds and returns a Select from the given name
func (o *Options) GetSelect(name string) (*Select, bool) {
	o.RLock()
	b, ok := o.Selects[name]
	o.RUnlock()
	return b, ok
}

// PutText put the Text to Texts map
func (o *Options) PutText(b *Text) {
	o.Lock()
	o.Texts[b.Name] = b
	o.Unlock()
}

// GetText finds and returns a Text from the given name
func (o *Options) GetText(name string) (*Text, bool) {
	o.RLock()
	b, ok := o.Texts[name]
	o.RUnlock()
	return b, ok
}

// Copy creates cloned one.
func (o *Options) Copy() *Options {
	o.RLock()
	defer o.RUnlock()

	opts := NewOptions()

	for _, v := range o.Buttons {
		opts.PutButton(v.Copy())
	}
	for _, v := range o.Checks {
		opts.PutCheck(v.Copy())
	}
	for _, v := range o.Ranges {
		opts.PutRange(v.Copy())
	}
	for _, v := range o.Selects {
		opts.PutSelect(v.Copy())
	}
	for _, v := range o.Texts {
		opts.PutText(v.Copy())
	}

	return opts
}
