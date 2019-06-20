// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package option

import (
	"reflect"
	"testing"
)

var (
	b = &Button{Name: "test"}
	c = &Check{Name: "test", Value: true, Default: false}
	r = &Range{Name: "test", Value: 100, Default: 100, Min: 0, Max: 200}
	s = &Select{
		Name:    "test",
		Value:   "a",
		Default: "b",
		Vars:    []string{"a", "b", "c"},
	}
	txt = &Text{Name: "test", Value: "abc", Default: "abc"}

	hasValue = &Options{
		Buttons: map[string]*Button{"test": b},
		Checks:  map[string]*Check{"test": c},
		Ranges:  map[string]*Range{"test": r},
		Selects: map[string]*Select{"test": s},
		Texts:   map[string]*Text{"test": txt},
	}
)

func TestOptions_Push(t *testing.T) {
	opts := NewOptions()

	opts.PutButton(b)
	opts.PutCheck(c)
	opts.PutRange(r)
	opts.PutSelect(s)
	opts.PutText(txt)

	if !reflect.DeepEqual(opts, hasValue) {
		t.Errorf(`
[app > domain > entity > engine > Options.Push]
Expected: %v
Actual:   %v
`, hasValue, opts)
	}
}
