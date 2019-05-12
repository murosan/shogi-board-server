// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package engine

import "fmt"

// Options is a option holder.
// Although each maps are exported to convert to JSON,
// do not touch directly from other packages.
type Options struct {
	Buttons map[string]*Button `json:"buttons"`
	Checks  map[string]*Check  `json:"checks"`
	Ranges  map[string]*Range  `json:"ranges"`
	Selects map[string]*Select `json:"selects"`
	Texts   map[string]*Text   `json:"texts"`
}

// Push appends given option to each options map.
// Returns error if given option is not valid type.
// If the name is already exists, Push overrides it.
func (o *Options) Push(v Option) error {
	switch a := v.(type) {
	case *Button:
		o.Buttons[a.Name] = a
	case *Check:
		o.Checks[a.Name] = a
	case *Range:
		o.Ranges[a.Name] = a
	case *Select:
		o.Selects[a.Name] = a
	case *Text:
		o.Texts[a.Name] = a
	default:
		return fmt.Errorf(`%v is not a valid option type`, v)
	}

	return nil
}

// GetButton return a Button from Options as a first value.
// The second value is true if the Button exists, false otherwise.
func (o *Options) GetButton(name string) (*Button, bool) {
	b, ok := o.Buttons[name]
	return b, ok
}

// GetCheck return a Check from Options as a first value.
// The second value is true if the Check exists, false otherwise.
func (o *Options) GetCheck(name string) (*Check, bool) {
	c, ok := o.Checks[name]
	return c, ok
}

// GetRange return a Range from Options as a first value.
// The second value is true if the Range exists, false otherwise.
func (o *Options) GetRange(name string) (*Range, bool) {
	r, ok := o.Ranges[name]
	return r, ok
}

// GetSelect return a Select from Options as a first value.
// The second value is true if the Select exists, false otherwise.
func (o *Options) GetSelect(name string) (*Select, bool) {
	s, ok := o.Selects[name]
	return s, ok
}

// GetText return a Text from Options as a first value.
// The second value is true if the Text exists, false otherwise.
func (o *Options) GetText(name string) (*Text, bool) {
	t, ok := o.Texts[name]
	return t, ok
}
