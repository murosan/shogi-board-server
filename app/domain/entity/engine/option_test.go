// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package engine

import "testing"

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

func TestCheck_Set(t *testing.T) {
	cases := []struct {
		opt *Check
		new bool
	}{
		{&Check{Name: "test", Value: true, Default: false}, true},
		{&Check{Name: "test", Value: false, Default: false}, true},
		{&Check{Name: "test", Value: true, Default: false}, false},
		{&Check{Name: "test", Value: false, Default: false}, false},
	}

	for i, c := range cases {
		c.opt.Set(c.new)
		res := c.opt.Value

		if res != c.new {
			t.Errorf(`
[app > domain > entity > engine > Check.Set]
Index:    %d
Expected: %t
Actual:   %t
`, i, c.new, res)
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

func TestRange_Set(t *testing.T) {
	cases := []struct {
		opt     *Range
		new     int
		want    int
		isError bool
	}{
		{
			&Range{Name: "test", Value: 100, Default: 100, Min: 0, Max: 200},
			100,
			100,
			false,
		},
		{
			&Range{Name: "test", Value: -100, Default: 100, Min: 0, Max: 200},
			200,
			200,
			false,
		},
		{
			&Range{Name: "test", Value: -100, Default: 100, Min: 0, Max: 200},
			300,
			-100,
			true,
		},
		{
			&Range{Name: "test", Value: -100, Default: 100, Min: -11, Max: 200},
			-12,
			-100,
			true,
		},
	}

	for i, c := range cases {
		err := c.opt.Set(c.new)
		isErr := err != nil

		if isErr != c.isError {
			t.Errorf(`
[app > domain > entity > engine > Range.Set]
Index:         %d
ErrorExpected: %t
WasError:      %t
`, i, c.isError, isErr)
		}

		res := c.opt.Value

		if res != c.want {
			t.Errorf(`
[app > domain > entity > engine > Range.Set]
Index:    %d
Expected: %d
Actual:   %d
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

func TestSelect_Set(t *testing.T) {
	cases := []struct {
		opt     *Select
		new     string
		want    string
		isError bool
	}{
		{
			&Select{
				Name:    "test",
				Value:   "a",
				Default: "b",
				Vars:    []string{"a", "b", "c"},
			},
			"c",
			"c",
			false,
		},
		{
			&Select{
				Name:    "test",
				Value:   "a",
				Default: "b",
				Vars:    []string{"a", "b", "c"},
			},
			"a",
			"a",
			false,
		},
		{
			&Select{
				Name:    "test",
				Value:   "a",
				Default: "b",
				Vars:    []string{"a", "b", "c"},
			},
			"d",
			"a",
			true,
		},
	}

	for i, c := range cases {
		err := c.opt.Set(c.new)
		isErr := err != nil

		if isErr != c.isError {
			t.Errorf(`
[app > domain > entity > engine > Select.Set]
Index:         %d
ErrorExpected: %t
WasError:      %t
`, i, c.isError, isErr)
		}

		res := c.opt.Value

		if res != c.want {
			t.Errorf(`
[app > domain > entity > engine > Select.Set]
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

func TestText_Set(t *testing.T) {
	cases := []struct {
		opt *Text
		new string
	}{
		{&Text{Name: "test", Value: "abc", Default: "abc"}, "def"},
		{&Text{Name: "test", Value: "abc", Default: "abc"}, "  "},
	}

	for i, c := range cases {
		c.opt.Set(c.new)
		res := c.opt.Value

		if res != c.new {
			t.Errorf(`
[app > domain > entity > engine > Text.Set]
Index:    %d
Expected: %s
Actual:   %s
`, i, c.new, res)
		}
	}
}
