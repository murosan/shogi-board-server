// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package option

import (
	"bytes"
	"testing"
)

func TestButton_GetName(t *testing.T) {
	cases := []struct {
		in   Button
		want []byte
	}{
		{Button{[]byte("btn-name")}, []byte("btn-name")},
		{Button{[]byte("")}, []byte("")},
		{Button{[]byte(" ")}, []byte(" ")},
		{Button{[]byte("%\n|t\t")}, []byte("%\n|t\t")},
	}

	for i, c := range cases {
		if !bytes.Equal(c.in.GetName(), c.want) {
			t.Errorf("Button.GetName was not as Expected\nIndex: %d\nWant: %s\nActual: %s", i, string(c.in.GetName()), string(c.want))
		}
	}
}
func TestButton_Usi(t *testing.T) {
	cases := []struct {
		in   Button
		want []byte
	}{
		{Button{[]byte("btn-name")}, []byte("setoption name btn-name")},
		{Button{[]byte("")}, []byte("setoption name ")},
		{Button{[]byte(" ")}, []byte("setoption name  ")},
		{Button{[]byte("%\n|t\t")}, []byte("setoption name %\n|t\t")},
	}

	for i, c := range cases {
		if !bytes.Equal(c.in.Usi(), c.want) {
			t.Errorf("Button.Usi was not as Expected\nIndex: %d\nWant: %s\nActual: %s", i, string(c.in.Usi()), string(c.want))
		}
	}
}
