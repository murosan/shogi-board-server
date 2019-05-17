// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package engine

import "fmt"

// Options is a option holder.
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
