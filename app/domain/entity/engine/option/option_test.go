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
		getNameTestHelper(t, i, c.in, c.want)
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
		usiTestHelper(t, i, c.in, c.want)
	}
}

func TestCheck_GetName(t *testing.T) {
	cases := []struct {
		in   Check
		want []byte
	}{
		{Check{[]byte("chk-name"), true, true}, []byte("chk-name")},
		{Check{Name: []byte("")}, []byte("")},
		{Check{[]byte(" "), false, true}, []byte(" ")},
		{Check{Name: []byte("%\n|t\t")}, []byte("%\n|t\t")},
	}

	for i, c := range cases {
		getNameTestHelper(t, i, c.in, c.want)
	}
}

func TestCheck_Usi(t *testing.T) {
	cases := []struct {
		in   Check
		want []byte
	}{
		{Check{[]byte("chk-name"), true, true}, []byte("setoption name chk-name value true")},
		{Check{Name: []byte("")}, []byte("setoption name  value false")},
		{Check{[]byte(" "), false, true}, []byte("setoption name   value false")},
		{Check{Name: []byte("%\n|t\t")}, []byte("setoption name %\n|t\t value false")},
	}

	for i, c := range cases {
		usiTestHelper(t, i, c.in, c.want)
	}
}

func TestFileName_GetName(t *testing.T) {
	cases := []struct {
		in   FileName
		want []byte
	}{
		{FileName{[]byte("file-name"), []byte("engine.exe"), []byte("engine.exe")}, []byte("file-name")},
		{FileName{Name: []byte("")}, []byte("")},
		{FileName{[]byte(" "), []byte("engine.exe"), []byte("engine.exe")}, []byte(" ")},
		{FileName{Name: []byte("%\n|t\t")}, []byte("%\n|t\t")},
	}

	for i, c := range cases {
		getNameTestHelper(t, i, c.in, c.want)
	}
}

func TestFileName_Usi(t *testing.T) {
	cases := []struct {
		in   FileName
		want []byte
	}{
		{FileName{[]byte("file-name"), []byte("engine.exe"), []byte("engine.exe")}, []byte("setoption name file-name value engine.exe")},
		{FileName{Name: []byte("")}, []byte("setoption name  value ")},
		{FileName{[]byte(" "), []byte("engine.exe"), []byte("engine.exe")}, []byte("setoption name   value engine.exe")},
		{FileName{Name: []byte("%\n|t\t")}, []byte("setoption name %\n|t\t value ")},
	}

	for i, c := range cases {
		usiTestHelper(t, i, c.in, c.want)
	}
}

func TestSelect_GetName(t *testing.T) {
	cases := []struct {
		in   Select
		want []byte
	}{
		{Select{[]byte("sel-name"), 1, [][]byte{[]byte("one"), []byte("two"), []byte("three")}}, []byte("sel-name")},
		{Select{Name: []byte("")}, []byte("")},
		{Select{[]byte(" "), 2, [][]byte{[]byte("one"), []byte("two"), []byte("three")}}, []byte(" ")},
		{Select{Name: []byte("%\n|t\t")}, []byte("%\n|t\t")},
	}

	for i, c := range cases {
		getNameTestHelper(t, i, c.in, c.want)
	}
}

func TestSelect_Usi(t *testing.T) {
	cases := []struct {
		in   Select
		want []byte
	}{
		{Select{[]byte("sel-name"), 1, [][]byte{[]byte("one"), []byte("two"), []byte("three")}}, []byte("setoption name sel-name value two")},
		{Select{[]byte(" "), 2, [][]byte{[]byte("one"), []byte("two"), []byte("three")}}, []byte("setoption name   value three")},
	}

	for i, c := range cases {
		usiTestHelper(t, i, c.in, c.want)
	}
}

func TestSpin_GetName(t *testing.T) {
	cases := []struct {
		in   Spin
		want []byte
	}{
		{Spin{[]byte("spn-nm"), 123, 0, -100, 1000}, []byte("spn-nm")},
		{Spin{[]byte("spn-nm2"), -500, -100, -10000, 1000}, []byte("spn-nm2")},
	}

	for i, c := range cases {
		getNameTestHelper(t, i, c.in, c.want)
	}
}

func TestSpin_Usi(t *testing.T) {
	cases := []struct {
		in   Spin
		want []byte
	}{
		{Spin{[]byte("spn-nm"), 123, 0, -100, 1000}, []byte("setoption name spn-nm value 123")},
		{Spin{[]byte("spn-nm2"), -500, -100, -10000, 1000}, []byte("setoption name spn-nm2 value -500")},
	}

	for i, c := range cases {
		usiTestHelper(t, i, c.in, c.want)
	}
}

func TestString_GetName(t *testing.T) {
	cases := []struct {
		in   String
		want []byte
	}{
		{String{[]byte("str-name"), []byte("engine.exe"), []byte("engine.exe")}, []byte("str-name")},
		{String{Name: []byte("")}, []byte("")},
		{String{[]byte(" "), []byte("engine.exe"), []byte("engine.exe")}, []byte(" ")},
		{String{Name: []byte("%\n|t\t")}, []byte("%\n|t\t")},
	}

	for i, c := range cases {
		getNameTestHelper(t, i, c.in, c.want)
	}
}

func TestString_Usi(t *testing.T) {
	cases := []struct {
		in   String
		want []byte
	}{
		{String{[]byte("str-name"), []byte("engine.exe"), []byte("engine.exe")}, []byte("setoption name str-name value engine.exe")},
		{String{Name: []byte("")}, []byte("setoption name  value ")},
		{String{[]byte(" "), []byte("engine.exe"), []byte("engine.exe")}, []byte("setoption name   value engine.exe")},
		{String{Name: []byte("%\n|t\t")}, []byte("setoption name %\n|t\t value ")},
	}

	for i, c := range cases {
		usiTestHelper(t, i, c.in, c.want)
	}
}

func getNameTestHelper(t *testing.T, i int, o Option, want []byte) {
	t.Helper()
	if !bytes.Equal(o.GetName(), want) {
		t.Errorf(`Option.GetName was not as expected
Index: %d
Input: %v
Want: %s
Actual: %s
`, i, o, string(want), string(o.GetName()))
	}
}

func usiTestHelper(t *testing.T, i int, o Option, want []byte) {
	t.Helper()
	if !bytes.Equal(o.Usi(), want) {
		t.Errorf(`Option.Usi was not as expected
Index: %d
Input: %v
Want: %s
Actual: %s
`, i, o, string(want), string(o.Usi()))
	}
}
