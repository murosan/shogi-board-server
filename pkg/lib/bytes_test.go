// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lib

import "testing"

var (
	one   = []byte("one")
	two   = []byte("two")
	three = []byte("three")

	arr = [][]byte{one, two, three}
)

func TestIndexOfBytes(t *testing.T) {
	cases := []struct {
		b    [][]byte
		t    []byte
		want int
	}{
		{arr, one, 0},
		{arr, three, 2},
		{arr, []byte("four"), -1},
	}

	for i, c := range cases {
		if r := IndexOfBytes(c.b, c.t); r != c.want {
			t.Errorf("Result was not as expected.\nIndex: %v\nExpected: %v\nActual: %v", i, c.want, r)
		}
	}
}

func TestEqualBytes(t *testing.T) {
	cases := []struct {
		in1, in2 [][]byte
		want     bool
	}{
		{arr, arr, true},
		{[][]byte{}, [][]byte{}, true},
		{arr, [][]byte{one, two, []byte("threee")}, false,},
		{arr, [][]byte{one, two}, false,},
	}

	for i, c := range cases {
		if r := EqualBytes(c.in1, c.in2); r != c.want {
			t.Errorf("Result was not as expected.\nIndex: %v\nExpected: %v\nActual: %v", i, c.want, r)
		}
	}
}
