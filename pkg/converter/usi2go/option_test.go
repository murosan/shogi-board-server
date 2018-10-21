// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package usi2go

import (
	"bytes"
	"strings"
	"testing"

	"github.com/murosan/shogi-proxy-server/pkg/converter/models"
	"github.com/murosan/shogi-proxy-server/pkg/lib"
	"github.com/murosan/shogi-proxy-server/pkg/msg"
)

var emp = []byte("")

func TestParseId(t *testing.T) {
	cases := []struct {
		in, name, val []byte
		err           error
	}{
		{[]byte("id name computer_name"), name, []byte("computer_name"), nil},
		{[]byte("id name "), emp, emp, msg.InvalidIdSyntax},
		{[]byte("id neimu typo_key"), emp, emp, msg.UnknownOption},
		{[]byte("id author computer_author"), author, []byte("computer_author"), nil},
		{[]byte("id author"), []byte(""), emp, msg.InvalidIdSyntax},
		{[]byte("id auther typo_key"), emp, emp, msg.UnknownOption},
	}

	for _, c := range cases {
		n, v, e := ParseId(c.in)
		if e != c.err {
			t.Errorf("Returned error was not as expected.\nInput: %v, Expected: %v, Actual: %v", string(c.in), c.err, e)
		}
		if !bytes.Equal(n, c.name) {
			t.Errorf("Name of OptionId was not as expected.\nInput: %v, Expected: %v, Actual: %v", string(c.in), string(c.name), string(n))
		}
		if !bytes.Equal(v, c.val) {
			t.Errorf("Value of OptionId was not as expected.\nInput: %v, Expected: %v, Actual: %v", string(c.in), string(c.val), string(v))
		}
	}
}

func TestParseOpt(t *testing.T) {
	cases := []struct {
		in   string
		want models.Check
		err  error
	}{
		{"option name UseBook type check default true", models.Check{Name: []byte("UseBook"), Val: true, Default: true}, nil},
		{"   option name UseBook type check default true   ", models.Check{Name: []byte("UseBook"), Val: true, Default: true}, nil},
		{"option name UseBook type check default ", models.Check{}, msg.InvalidOptionSyntax},
		{"option name UseBook type check default not_bool", models.Check{}, msg.InvalidOptionSyntax},
		{"option name UseBook type check dlft true", models.Check{}, msg.InvalidOptionSyntax},
	}

	for _, c := range cases {
		o, err := ParseOpt([]byte(c.in))
		basicOptionMatching(t, c.in, o, c.want, err, c.err)
		switch v := o.(type) {
		case models.Check:
			if c.want.Default != v.Default || c.want.Val != v.Val {
				t.Errorf("Mismatch values.\nInput: %v\nExpected: %v\nActual: %v",c.in, c.want, v)
			}
		}
	}
}

func TestParseOpt2(t *testing.T) {
	cases := []struct {
		in   string
		want models.Spin
		err  error
	}{
		{"option name Selectivity type spin default 2 min 0 max 4", models.Spin{[]byte("Selectivity"), 2, 2, 0, 4}, nil},
		{"option name Selectivity type spin default -100 min -123456 max 54321 ", models.Spin{[]byte("Selectivity"), -100, -100, -123456, 54321}, nil},
		{"option name Selectivity type spin min 0 max 4", models.Spin{}, msg.InvalidOptionSyntax},
		{"option name Selectivity type spin default 2", models.Spin{}, msg.InvalidOptionSyntax},
		{"option name Selectivity type spin min 0 max 4 default 2", models.Spin{}, msg.InvalidOptionSyntax},
		{"option name Selectivity type spin default two min 0 max 4", models.Spin{}, msg.InvalidOptionSyntax},
	}
	for _, c := range cases {
		o, err := ParseOpt([]byte(c.in))
		basicOptionMatching(t, c.in, o, c.want, err, c.err)
		switch v := o.(type) {
		case models.Spin:
			if c.want.Default != v.Default || c.want.Val != v.Val || c.want.Min != v.Min || c.want.Max != v.Max {
				t.Errorf("Mismatch values.\nInput: %v\nExpected: %v\nActual: %v",c.in, c.want, v)
			}
		}
	}
}

func TestParseOpt3(t *testing.T) {
	cases := []struct {
		in   string
		want models.Select
		err  error
	}{
		{
			"option name Style type combo default Normal var Solid var Normal var Risky",
			models.Select{[]byte("Style"), 1, [][]byte{[]byte("Solid"), []byte("Normal"), []byte("Risky")}},
			nil,
		},
		{"option name Style type combo default None var Solid var Normal var Risky", models.Select{}, msg.InvalidOptionSyntax},
		{"option name Style type combo var Solid var Normal var Risky", models.Select{}, msg.InvalidOptionSyntax},
		{"option name Style type combo default Normal", models.Select{}, msg.InvalidOptionSyntax},
	}
	for _, c := range cases {
		o, err := ParseOpt([]byte(c.in))
		basicOptionMatching(t, c.in, o, c.want, err, c.err)
		switch v := o.(type) {
		case models.Select:
			if c.want.Index != v.Index || !lib.EqualBytes(c.want.Vars, v.Vars) {
				t.Errorf("Mismatch values.\nInput: %v\nExpected: %v\nActual: %v",c.in, c.want, v)
			}
		}
	}
}

func TestParseOpt4(t *testing.T) {
	cases := []struct {
		in   string
		want models.Button
		err  error
	}{
		{"option name ResetLearning type button", models.Button{[]byte("ResetLearning")}, nil},
		{"option name <empty> type button", models.Button{[]byte("<empty>")}, nil}, // まぁいい
		{"option name ResetLearning type button sur", models.Button{}, msg.InvalidOptionSyntax},
		{"option name 1 type button", models.Button{[]byte("1")}, nil},
	}
	for _, c := range cases {
		o, err := ParseOpt([]byte(c.in))
		basicOptionMatching(t, c.in, o, c.want, err, c.err)
	}
}

func TestParseOpt5(t *testing.T) {
	cases := []struct {
		in   string
		want models.String
		err  error
	}{
		{"option name BookFile type string default public.bin", models.String{[]byte("BookFile"), []byte("public.bin"), []byte("public.bin")}, nil},
		{"option name BookFile type string default public.bin var a", models.String{}, msg.InvalidOptionSyntax},
		{"option name BookFile type string", models.String{}, msg.InvalidOptionSyntax},
		{"option name BookFile type string public.bin", models.String{}, msg.InvalidOptionSyntax},
	}
	for _, c := range cases {
		o, err := ParseOpt([]byte(c.in))
		basicOptionMatching(t, c.in, o, c.want, err, c.err)
		switch v := o.(type) {
		case models.String:
			if !bytes.Equal(v.Val, c.want.Val) || !bytes.Equal(v.Default, c.want.Default) {
				t.Errorf("Mismatch values.\nInput: %v\nExpected: %v\nActual: %v",c.in, c.want, v)
			}
		}
	}
}

func TestParseOpt6(t *testing.T) {
	cases := []struct {
		in   string
		want models.FileName
		err  error
	}{
		{"option name LearningFile type filename default <empty>", models.FileName{[]byte("LearningFile"), []byte("<empty>"), []byte("<empty>")}, nil},
		{"option name LearningFile type filename default <empty> var a", models.FileName{}, msg.InvalidOptionSyntax},
		{"option name LearningFile type filename", models.FileName{}, msg.InvalidOptionSyntax},
		{"option name LearningFile type filename <empty>", models.FileName{}, msg.InvalidOptionSyntax},
	}
	for _, c := range cases {
		o, err := ParseOpt([]byte(c.in))
		basicOptionMatching(t, c.in, o, c.want, err, c.err)
		switch v := o.(type) {
		case models.FileName:
			if !bytes.Equal(v.Val, c.want.Val) || !bytes.Equal(v.Default, c.want.Default) {
				t.Errorf("Mismatch values.\nInput: %v\nExpected: %v\nActual: %v",c.in, c.want, v)
			}
		}
	}
}

// in: input
// o1: Returned Option
// o2: Expected Option
// e1: Returned Error
// e2: Expected Error
func basicOptionMatching(t *testing.T, in string, o1, o2 models.Option, e1, e2 error) {
	t.Helper()
	if (e1 == nil && e2 != nil) || (e1 != nil && e2 == nil) {
		t.Errorf("Returned error was not as expected.\nInput: %v\nExpected: %v\nActual: %v", in, e2, e1)
	}

	// 予想通りのエラーが返った
	if e1 != nil && strings.Contains(string(e1.Error()), string(e2.Error())) {
		return
	}

	// エラーは返ったが、想定と違った
	if e1 != nil && !strings.Contains(string(e1.Error()), string(e2.Error())) {
		t.Errorf("Returned error was not as expected.\nInput: %v\nExpected: %v\nActual: %v", in, e2, e1)
	}

	// USIコマンドが想定通りかどうか
	if !bytes.Equal(o1.Usi(), o2.Usi()) {
		t.Errorf("Set value was not as expected.\nInput: %v\nExpected: %v\nActual: %v", in, string(o1.Usi()), string(o2.Usi()))
	}
}
