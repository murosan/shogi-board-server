// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package intutil

import "testing"

func TestSliceEquals(t *testing.T) {
	cases := []struct {
		a, b     []int
		expected bool
	}{
		{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}, true},
		{[]int{1, 2, 3, 4}, []int{1, 2, 3, 7}, false},
		{[]int{-100}, []int{-100}, true},
		{[]int{100}, []int{100, 200}, false},
		{[]int{}, []int{}, true},
	}

	for i, c := range cases {
		res := SliceEquals(c.a, c.b)
		if res != c.expected {
			t.Errorf(`[Int slice] Equals method is not property working.
Index:    %d
InputA:   %v
InputB:   %v
Expected: %t
Actual:   %t
`, i, c.a, c.b, c.expected, res)
		}
	}
}
