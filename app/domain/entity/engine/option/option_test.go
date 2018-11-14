// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package option

import (
	"testing"
)

func TestButton_GetName(t *testing.T) {
	cases := []struct {
		in   Button
		want string
	}{
		{Button{"btn-name"}, "btn-name"},
		{Button{""}, ""},
		{Button{" "}, " "},
		{Button{"%\n|t\t"}, "%\n|t\t"},
	}

	for i, c := range cases {
		getNameTestHelper(t, i, c.in, c.want)
	}
}

func TestButton_Usi(t *testing.T) {
	cases := []struct {
		in   Button
		want string
	}{
		{Button{"btn-name"}, "setoption name btn-name"},
		{Button{""}, "setoption name "},
		{Button{" "}, "setoption name  "},
		{Button{"%\n|t\t"}, "setoption name %\n|t\t"},
	}

	for i, c := range cases {
		usiTestHelper(t, i, c.in, c.want)
	}
}

func TestCheck_GetName(t *testing.T) {
	cases := []struct {
		in   Check
		want string
	}{
		{Check{"chk-name", true, true}, "chk-name"},
		{Check{Name: ""}, ""},
		{Check{" ", false, true}, " "},
		{Check{Name: "%\n|t\t"}, "%\n|t\t"},
	}

	for i, c := range cases {
		getNameTestHelper(t, i, c.in, c.want)
	}
}

func TestCheck_Usi(t *testing.T) {
	cases := []struct {
		in   Check
		want string
	}{
		{Check{"chk-name", true, true}, "setoption name chk-name value true"},
		{Check{Name: ""}, "setoption name  value false"},
		{Check{" ", false, true}, "setoption name   value false"},
		{Check{Name: "%\n|t\t"}, "setoption name %\n|t\t value false"},
	}

	for i, c := range cases {
		usiTestHelper(t, i, c.in, c.want)
	}
}

func TestFileName_GetName(t *testing.T) {
	cases := []struct {
		in   FileName
		want string
	}{
		{FileName{"file-name", "engine.exe", "engine.exe"}, "file-name"},
		{FileName{Name: ""}, ""},
		{FileName{" ", "engine.exe", "engine.exe"}, " "},
		{FileName{Name: "%\n|t\t"}, "%\n|t\t"},
	}

	for i, c := range cases {
		getNameTestHelper(t, i, c.in, c.want)
	}
}

func TestFileName_Usi(t *testing.T) {
	cases := []struct {
		in   FileName
		want string
	}{
		{FileName{"file-name", "engine.exe", "engine.exe"}, "setoption name file-name value engine.exe"},
		{FileName{Name: ""}, "setoption name  value "},
		{FileName{" ", "engine.exe", "engine.exe"}, "setoption name   value engine.exe"},
		{FileName{Name: "%\n|t\t"}, "setoption name %\n|t\t value "},
	}

	for i, c := range cases {
		usiTestHelper(t, i, c.in, c.want)
	}
}

func TestSelect_GetName(t *testing.T) {
	cases := []struct {
		in   Select
		want string
	}{
		{Select{"sel-name", 1, []string{"one", "two", "three"}}, "sel-name"},
		{Select{Name: ""}, ""},
		{Select{" ", 2, []string{"one", "two", "three"}}, " "},
		{Select{Name: "%\n|t\t"}, "%\n|t\t"},
	}

	for i, c := range cases {
		getNameTestHelper(t, i, c.in, c.want)
	}
}

func TestSelect_Usi(t *testing.T) {
	cases := []struct {
		in   Select
		want string
	}{
		{Select{"sel-name", 1, []string{"one", "two", "three"}}, "setoption name sel-name value two"},
		{Select{" ", 2, []string{"one", "two", "three"}}, "setoption name   value three"},
	}

	for i, c := range cases {
		usiTestHelper(t, i, c.in, c.want)
	}
}

func TestSpin_GetName(t *testing.T) {
	cases := []struct {
		in   Spin
		want string
	}{
		{Spin{"spn-nm", 123, 0, -100, 1000}, "spn-nm"},
		{Spin{"spn-nm2", -500, -100, -10000, 1000}, "spn-nm2"},
	}

	for i, c := range cases {
		getNameTestHelper(t, i, c.in, c.want)
	}
}

func TestSpin_Usi(t *testing.T) {
	cases := []struct {
		in   Spin
		want string
	}{
		{Spin{"spn-nm", 123, 0, -100, 1000}, "setoption name spn-nm value 123"},
		{Spin{"spn-nm2", -500, -100, -10000, 1000}, "setoption name spn-nm2 value -500"},
	}

	for i, c := range cases {
		usiTestHelper(t, i, c.in, c.want)
	}
}

func TestString_GetName(t *testing.T) {
	cases := []struct {
		in   String
		want string
	}{
		{String{"str-name", "engine.exe", "engine.exe"}, "str-name"},
		{String{Name: ""}, ""},
		{String{" ", "engine.exe", "engine.exe"}, " "},
		{String{Name: "%\n|t\t"}, "%\n|t\t"},
	}

	for i, c := range cases {
		getNameTestHelper(t, i, c.in, c.want)
	}
}

func TestString_Usi(t *testing.T) {
	cases := []struct {
		in   String
		want string
	}{
		{String{"str-name", "engine.exe", "engine.exe"}, "setoption name str-name value engine.exe"},
		{String{Name: ""}, "setoption name  value "},
		{String{" ", "engine.exe", "engine.exe"}, "setoption name   value engine.exe"},
		{String{Name: "%\n|t\t"}, "setoption name %\n|t\t value "},
	}

	for i, c := range cases {
		usiTestHelper(t, i, c.in, c.want)
	}
}

func getNameTestHelper(t *testing.T, i int, o Option, want string) {
	t.Helper()
	if o.GetName() != want {
		t.Errorf(`Option.GetName was not as expected
Index: %d
Input: %v
Want: %s
Actual: %s
`, i, o, string(want), string(o.GetName()))
	}
}

func usiTestHelper(t *testing.T, i int, o Option, want string) {
	t.Helper()
	usi, _ := o.UpdateAndGetUsi() // TODO
	if usi != want {
		t.Errorf(`Option.UpdateAndGetUsi was not as expected
Index: %d
Input: %v
Want: %s
Actual: %s
`, i, o, string(want), usi)
	}
}
