// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package option

import (
	"strings"
	"testing"

	"github.com/murosan/shogi-proxy-server/app/domain/entity/exception"
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

func TestButton_Update(t *testing.T) {
	cases := []struct {
		in   *Button
		set  interface{}
		want string
		err  error
	}{
		{NewButton("btn-name"), "string", "setoption name btn-name", nil},
		{NewButton(""), 100, "setoption name ", nil},
		{NewButton(" "), true, "setoption name  ", nil},
		{NewButton("%\n|t\t"), []int{1, 2, 3}, "setoption name %\n|t\t", nil},
	}

	for i, c := range cases {
		setTestHelper(t, i, c.in, c.set, c.want, c.err)
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

func TestCheck_Update(t *testing.T) {
	cases := []struct {
		in   *Check
		set  interface{}
		want string
		err  error
	}{
		{NewCheck("chk-name", true, true), false, "setoption name chk-name value false", nil},
		{NewCheck("", false, false), false, "setoption name  value false", nil},
		{NewCheck(" ", false, true), true, "setoption name   value true", nil},
		{NewCheck("%\n|t\t", false, false), false, "setoption name %\n|t\t value false", nil},
		{NewCheck("name", true, false), 1, "", exception.InvalidOptionParameter},
		{NewCheck("name", true, false), 0, "", exception.InvalidOptionParameter},
		{NewCheck("name", true, false), "true", "", exception.InvalidOptionParameter},
	}

	for i, c := range cases {
		setTestHelper(t, i, c.in, c.set, c.want, c.err)
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

func TestSpin_Update(t *testing.T) {
	cases := []struct {
		in   *Spin
		set  interface{} // json.Unmarshal すると float64 になる
		want string
		err  error
	}{
		{NewSpin("spn-nm", 123, 0, -100, 1000), float64(567), "setoption name spn-nm value 567", nil},
		{NewSpin("spn-nm2", -500, -100, -10000, 1000), float64(-101), "setoption name spn-nm2 value -101", nil},
		{NewSpin("spn-nm2", -500, -100, -10000, 1000), float64(1000), "setoption name spn-nm2 value 1000", nil},
		{NewSpin("spn-nm2", -500, -100, -10000, 1000), float64(-10000), "setoption name spn-nm2 value -10000", nil},
		{NewSpin("spn-nm2", -500, -100, -10000, 1000), float64(1001), "", exception.InvalidOptionParameter},
		{NewSpin("spn-nm2", -500, -100, -10000, 1000), float64(-10001), "", exception.InvalidOptionParameter},
		{NewSpin("spn-nm2", -500, -100, -10000, 1000), "string", "", exception.InvalidOptionParameter},
	}

	for i, c := range cases {
		setTestHelper(t, i, c.in, c.set, c.want, c.err)
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

func TestSelect_Update(t *testing.T) {
	cases := []struct {
		in   *Select
		set  interface{}
		want string
		err  error
	}{
		{NewSelect("sel-name", "two", "one", []string{"one", "two", "three"}), "three", "setoption name sel-name value three", nil},
		{NewSelect(" ", "three", "three", []string{"one", "two", "three"}), "two", "setoption name   value two", nil},
		{NewSelect(" ", "three", "three", []string{"one", "two", "three"}), "none", "", exception.InvalidOptionParameter},
		{NewSelect(" ", "one", "one", []string{"one"}), "one", "setoption name   value one", nil},
		{NewSelect(" ", "three", "three", []string{"one", "two", "three"}), 100, "", exception.InvalidOptionParameter},
	}

	for i, c := range cases {
		setTestHelper(t, i, c.in, c.set, c.want, c.err)
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

func TestString_Update(t *testing.T) {
	cases := []struct {
		in   *String
		set  interface{}
		want string
		err  error
	}{
		{NewString("str-name", "engine.exe", "engine.exe"), "updated", "setoption name str-name value updated", nil},
		{NewString("", "", ""), "updated", "setoption name  value updated", nil},
		{NewString(" ", "engine.exe", "engine.exe"), "up up", "setoption name   value up up", nil},
		{NewString("%\n|t\t", "", ""), "set set", "setoption name %\n|t\t value set set", nil},
		{NewString("str-name", "engine.exe", "engine.exe"), "", "setoption name str-name value ", nil},
		{NewString("str-name", "engine.exe", "engine.exe"), 100, "", exception.InvalidOptionParameter},
	}

	for i, c := range cases {
		setTestHelper(t, i, c.in, c.set, c.want, c.err)
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

func TestFileName_Update(t *testing.T) {
	cases := []struct {
		in   *FileName
		set  interface{}
		want string
		err  error
	}{
		{NewFileName("str-name", "engine.exe", "engine.exe"), "updated", "setoption name str-name value updated", nil},
		{NewFileName("", "", ""), "updated", "setoption name  value updated", nil},
		{NewFileName(" ", "engine.exe", "engine.exe"), "up up", "setoption name   value up up", nil},
		{NewFileName("%\n|t\t", "", ""), "set set", "setoption name %\n|t\t value set set", nil},
		{NewFileName("str-name", "engine.exe", "engine.exe"), "", "setoption name str-name value ", nil},
		{NewFileName("str-name", "engine.exe", "engine.exe"), 100, "", exception.InvalidOptionParameter},
	}

	for i, c := range cases {
		setTestHelper(t, i, c.in, c.set, c.want, c.err)
	}
}

func getNameTestHelper(t *testing.T, i int, o Option, want string) {
	t.Helper()
	if o.GetName() != want {
		t.Errorf(`Option.GetName was not as expected
index:  %d
Input:  %v
Want:   %s
Actual: %s
`, i, o, want, o.GetName())
	}
}

func usiTestHelper(t *testing.T, i int, o Option, want string) {
	t.Helper()
	usi := o.Usi()
	if usi != want {
		t.Errorf(`Option.Usi was not as expected
index:  %d
Input:  %v
Want:   %s
Actual: %s
`, i, o, want, usi)
	}
}

func setTestHelper(t *testing.T, i int, o Option, set interface{}, want string, e error) {
	t.Helper()
	usi, err := o.Update(set)
	errMatches := err == e
	if err != nil && e != nil {
		errMatches = strings.Contains(err.Error(), e.Error())
	}
	if usi != want || !errMatches {
		t.Errorf(`Option.Update was not as expected
index:       %d
Input:       %v
SetVal:      %v
WantUsi:     %s
ActualUsi:   %s
WantError:   %v
ActualError: %v
`, i, o, set, want, usi, e, err)
	}
}
