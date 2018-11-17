// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package option

import (
	"testing"
)

func TestButton_GetName(t *testing.T) {
	cases := []struct {
		in   *Button
		want string
	}{
		{NewButton("btn-name"), "btn-name"},
		{NewButton(""), ""},
		{NewButton(" "), " "},
		{NewButton("%\n|t\t"), "%\n|t\t"},
	}

	for i, c := range cases {
		getNameTestHelper(t, i, c.in, c.want)
	}
}

func TestButton_Usi(t *testing.T) {
	cases := []struct {
		in   *Button
		want string
	}{
		{NewButton("btn-name"), "setoption name btn-name"},
		{NewButton(""), "setoption name "},
		{NewButton(" "), "setoption name  "},
		{NewButton("%\n|t\t"), "setoption name %\n|t\t"},
	}

	for i, c := range cases {
		usiTestHelper(t, i, c.in, c.want)
	}
}

func TestCheck_GetName(t *testing.T) {
	cases := []struct {
		in   *Check
		want string
	}{
		{NewCheck("chk-name", true, true), "chk-name"},
		{NewCheck("", false, false), ""},
		{NewCheck(" ", false, true), " "},
		{NewCheck("%\n|t\t", false, false), "%\n|t\t"},
	}

	for i, c := range cases {
		getNameTestHelper(t, i, c.in, c.want)
	}
}

func TestCheck_Usi(t *testing.T) {
	cases := []struct {
		in   *Check
		want string
	}{
		{NewCheck("chk-name", true, true), "setoption name chk-name value true"},
		{NewCheck("", false, false), "setoption name  value false"},
		{NewCheck(" ", false, true), "setoption name   value false"},
		{NewCheck("%\n|t\t", false, false), "setoption name %\n|t\t value false"},
	}

	for i, c := range cases {
		usiTestHelper(t, i, c.in, c.want)
	}
}

func TestFileName_GetName(t *testing.T) {
	cases := []struct {
		in   *FileName
		want string
	}{
		{NewFileName("file-name", "engine.exe", "engine.exe"), "file-name"},
		{NewFileName("", "", ""), ""},
		{NewFileName(" ", "engine.exe", "engine.exe"), " "},
		{NewFileName("%\n|t\t", "", ""), "%\n|t\t"},
	}

	for i, c := range cases {
		getNameTestHelper(t, i, c.in, c.want)
	}
}

func TestFileName_Usi(t *testing.T) {
	cases := []struct {
		in   *FileName
		want string
	}{
		{NewFileName("file-name", "engine.exe", "engine.exe"), "setoption name file-name value engine.exe"},
		{NewFileName("", "", ""), "setoption name  value "},
		{NewFileName(" ", "engine.exe", "engine.exe"), "setoption name   value engine.exe"},
		{NewFileName("%\n|t\t", "", ""), "setoption name %\n|t\t value "},
	}

	for i, c := range cases {
		usiTestHelper(t, i, c.in, c.want)
	}
}

func TestSelect_GetName(t *testing.T) {
	cases := []struct {
		in   *Select
		want string
	}{
		{NewSelect("sel-name", "one", "one", []string{"one", "two", "three"}), "sel-name"},
		{NewSelect("", "", "", []string{}), ""},
		{NewSelect(" ", "three", "three", []string{"one", "two", "three"}), " "},
		{NewSelect("%\n|t\t", "", "", []string{}), "%\n|t\t"},
	}

	for i, c := range cases {
		getNameTestHelper(t, i, c.in, c.want)
	}
}

func TestSelect_Usi(t *testing.T) {
	cases := []struct {
		in   *Select
		want string
	}{
		{NewSelect("sel-name", "two", "one", []string{"one", "two", "three"}), "setoption name sel-name value two"},
		{NewSelect(" ", "three", "three", []string{"one", "two", "three"}), "setoption name   value three"},
	}

	for i, c := range cases {
		usiTestHelper(t, i, c.in, c.want)
	}
}

func TestSpin_GetName(t *testing.T) {
	cases := []struct {
		in   *Spin
		want string
	}{
		{NewSpin("spn-nm", 123, 0, -100, 1000), "spn-nm"},
		{NewSpin("spn-nm2", -500, -100, -10000, 1000), "spn-nm2"},
	}

	for i, c := range cases {
		getNameTestHelper(t, i, c.in, c.want)
	}
}

func TestSpin_Usi(t *testing.T) {
	cases := []struct {
		in   *Spin
		want string
	}{
		{NewSpin("spn-nm", 123, 0, -100, 1000), "setoption name spn-nm value 123"},
		{NewSpin("spn-nm2", -500, -100, -10000, 1000), "setoption name spn-nm2 value -500"},
	}

	for i, c := range cases {
		usiTestHelper(t, i, c.in, c.want)
	}
}

func TestString_GetName(t *testing.T) {
	cases := []struct {
		in   *String
		want string
	}{
		{NewString("str-name", "engine.exe", "engine.exe"), "str-name"},
		{NewString("", "", ""), ""},
		{NewString(" ", "engine.exe", "engine.exe"), " "},
		{NewString("%\n|t\t", "", ""), "%\n|t\t"},
	}

	for i, c := range cases {
		getNameTestHelper(t, i, c.in, c.want)
	}
}

func TestString_Usi(t *testing.T) {
	cases := []struct {
		in   *String
		want string
	}{
		{NewString("str-name", "engine.exe", "engine.exe"), "setoption name str-name value engine.exe"},
		{NewString("", "", ""), "setoption name  value "},
		{NewString(" ", "engine.exe", "engine.exe"), "setoption name   value engine.exe"},
		{NewString("%\n|t\t", "", ""), "setoption name %\n|t\t value "},
	}

	for i, c := range cases {
		usiTestHelper(t, i, c.in, c.want)
	}
}

func getNameTestHelper(t *testing.T, i int, o Option, want string) {
	t.Helper()
	if o.GetName() != want {
		t.Errorf(`Option.GetName was not as expected
index: %d
Input: %v
Want: %s
Actual: %s
`, i, o, string(want), string(o.GetName()))
	}
}

func usiTestHelper(t *testing.T, i int, o Option, want string) {
	t.Helper()
	usi, _ := o.Usi() // TODO
	if usi != want {
		t.Errorf(`Option.Usi was not as expected
index: %d
Input: %v
Want: %s
Actual: %s
`, i, o, string(want), usi)
	}
}
