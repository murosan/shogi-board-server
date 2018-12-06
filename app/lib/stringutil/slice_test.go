// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stringutil

import "testing"

func TestSliceContains(t *testing.T) {
	cases := []struct {
		a        []string
		b        string
		expected bool
	}{
		{[]string{"one", "two", "three"}, "one", true},
		{[]string{"one", "two", "three"}, "threee", false},
		{[]string{"one"}, "one", true},
		{[]string{}, "", false},
		{nil, "", false},
	}

	for i, c := range cases {
		res := SliceContains(c.a, c.b)
		if res != c.expected {
			t.Errorf(`[String slice] Contains method is not property working.
Index:    %d
InputA:   %v
InputB:   %v
Expected: %t
Actual:   %t
`, i, c.a, c.b, c.expected, res)
		}
	}
}

func TestSliceEquals(t *testing.T) {
	cases := []struct {
		a, b     []string
		expected bool
	}{
		{[]string{"one", "two", "three"}, []string{"one", "two", "three"}, true},
		{[]string{"one", "two", "three"}, []string{"one", "two", "threee"}, false},
		{[]string{"one"}, []string{"one"}, true},
		{[]string{"one", "two"}, []string{"one", "two", "three"}, false},
		{[]string{}, []string{}, true},
		{[]string{}, nil, false},
		{nil, []string{}, false},
		{nil, nil, true},
	}

	for i, c := range cases {
		res := SliceEquals(c.a, c.b)
		if res != c.expected {
			t.Errorf(`[String slice] Equals method is not property working.
Index:    %d
InputA:   %v
InputB:   %v
Expected: %t
Actual:   %t
`, i, c.a, c.b, c.expected, res)
		}
	}
}
