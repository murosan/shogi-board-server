// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package engine

import (
	"testing"
)

func TestButton_ToUSI(t *testing.T) {
	b := &Button{Name: "test"}
	res := b.ToUSI()
	usi := "setoption name test"

	if res != usi {
		t.Errorf(`
[app > domain > entity > engine > Button.ToUSI]
Expected: %s
Actual:   %s
`, usi, res)
	}
}

func TestCheck_ToUSI(t *testing.T) {
	cases := []struct {
		opt  *Check
		want string
	}{
		{
			&Check{Name: "test", Value: true, Default: false},
			"setoption name test value true",
		},
		{
			&Check{Name: "test", Value: false, Default: false},
			"setoption name test value false",
		},
	}

	for i, c := range cases {
		res := c.opt.ToUSI()

		if res != c.want {
			t.Errorf(`
[app > domain > entity > engine > Check.ToUSI]
Index:    %d
Expected: %s
Actual:   %s
`, i, c.want, res)
		}
	}
}

func TestRange_ToUSI(t *testing.T) {
	cases := []struct {
		opt  *Range
		want string
	}{
		{
			&Range{Name: "test", Value: 100, Default: 100, Min: 0, Max: 200},
			"setoption name test value 100",
		},
		{
			&Range{Name: "test", Value: -100, Default: 100, Min: 0, Max: 200},
			"setoption name test value -100",
		},
	}

	for i, c := range cases {
		res := c.opt.ToUSI()

		if res != c.want {
			t.Errorf(`
[app > domain > entity > engine > Range.ToUSI]
Index:    %d
Expected: %s
Actual:   %s
`, i, c.want, res)
		}
	}
}

func TestSelect_ToUSI(t *testing.T) {
	cases := []struct {
		opt  *Select
		want string
	}{
		{
			&Select{
				Name:    "test",
				Value:   "a",
				Default: "b",
				Vars:    []string{"a", "b", "c"},
			},
			"setoption name test value a",
		},
		{
			&Select{
				Name:    "test",
				Value:   "",
				Default: "",
				Vars:    []string{""},
			},
			"setoption name test value ",
		},
		{
			&Select{
				Name:    "test",
				Value:   " ",
				Default: "",
				Vars:    []string{"", " "},
			},
			"setoption name test value  ",
		},
	}

	for i, c := range cases {
		res := c.opt.ToUSI()

		if res != c.want {
			t.Errorf(`
[app > domain > entity > engine > Select.ToUSI]
Index:    %d
Expected: %s
Actual:   %s
`, i, c.want, res)
		}
	}
}

func TestText_ToUSI(t *testing.T) {
	cases := []struct {
		opt  *Text
		want string
	}{
		{
			&Text{Name: "test", Value: "abc", Default: "abc"},
			"setoption name test value abc",
		},
		{
			&Text{Name: "test", Value: "abc  ", Default: ""},
			"setoption name test value abc  ",
		},
	}

	for i, c := range cases {
		res := c.opt.ToUSI()

		if res != c.want {
			t.Errorf(`
[app > domain > entity > engine > Text.ToUSI]
Index:    %d
Expected: %s
Actual:   %s
`, i, c.want, res)
		}
	}
}
