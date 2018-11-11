// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package from_usi

import (
	"bytes"
	"github.com/murosan/shogi-proxy-server/app/lib/stringutil"
	"strings"
	"testing"

	"github.com/murosan/shogi-proxy-server/app/domain/entity/engine/option"
	"github.com/murosan/shogi-proxy-server/app/domain/entity/exception"
)

var emp = []byte("")

func TestFromUsi_EngineID(t *testing.T) {
	fu := NewFromUsi()
	cases := []struct {
		in, name, val []byte
		err           error
	}{
		{[]byte("id name computer_name"), name, []byte("computer_name"), nil},
		{[]byte("id name "), emp, emp, exception.InvalidIdSyntax},
		{[]byte("id neimu typo_key"), emp, emp, exception.UnknownOption},
		{[]byte("id author computer_author"), author, []byte("computer_author"), nil},
		{[]byte("id author"), []byte(""), emp, exception.InvalidIdSyntax},
		{[]byte("id auther typo_key"), emp, emp, exception.UnknownOption},
	}

	for _, c := range cases {
		n, v, e := fu.EngineID(c.in)
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

func TestFromUsi_Option(t *testing.T) {
	fu := NewFromUsi()
	cases := []struct {
		in   string
		want option.Check
		err  error
	}{
		{"option name UseBook type check default true", option.Check{Name: "UseBook", Val: true, Default: true}, nil},
		{"   option name UseBook type check default true   ", option.Check{Name: "UseBook", Val: true, Default: true}, nil},
		{"option name UseBook type check default ", option.Check{}, exception.InvalidOptionSyntax},
		{"option name UseBook type check default not_bool", option.Check{}, exception.InvalidOptionSyntax},
		{"option name UseBook type check dlft true", option.Check{}, exception.InvalidOptionSyntax},
	}

	for _, c := range cases {
		o, err := fu.Option([]byte(c.in))
		basicOptionMatching(t, c.in, o, c.want, err, c.err)
		switch v := o.(type) {
		case option.Check:
			if c.want.Default != v.Default || c.want.Val != v.Val {
				t.Errorf("Mismatch values.\nInput: %v\nExpected: %v\nActual: %v", c.in, c.want, v)
			}
		}
	}
}

func TestParseOpt2(t *testing.T) {
	fu := NewFromUsi()
	cases := []struct {
		in   string
		want option.Spin
		err  error
	}{
		{"option name Selectivity type spin default 2 min 0 max 4", option.Spin{Name: "Selectivity", Val: 2, Default: 2, Min: 0, Max: 4}, nil},
		{"option name Selectivity type spin default -100 min -123456 max 54321 ", option.Spin{Name: "Selectivity", Val: -100, Default: -100, Min: -123456, Max: 54321}, nil},
		{"option name Selectivity type spin min 0 max 4", option.Spin{}, exception.InvalidOptionSyntax},
		{"option name Selectivity type spin default 2", option.Spin{}, exception.InvalidOptionSyntax},
		{"option name Selectivity type spin min 0 max 4 default 2", option.Spin{}, exception.InvalidOptionSyntax},
		{"option name Selectivity type spin default two min 0 max 4", option.Spin{}, exception.InvalidOptionSyntax},
	}
	for _, c := range cases {
		o, err := fu.Option([]byte(c.in))
		basicOptionMatching(t, c.in, o, c.want, err, c.err)
		switch v := o.(type) {
		case option.Spin:
			if c.want.Default != v.Default || c.want.Val != v.Val || c.want.Min != v.Min || c.want.Max != v.Max {
				t.Errorf("Mismatch values.\nInput: %v\nExpected: %v\nActual: %v", c.in, c.want, v)
			}
		}
	}
}

func TestParseOpt3(t *testing.T) {
	fu := NewFromUsi()
	cases := []struct {
		in   string
		want option.Select
		err  error
	}{
		{
			"option name Style type combo default Normal var Solid var Normal var Risky",
			option.Select{Name: "Style", Index: 1, Vars: []string{"Solid", "Normal", "Risky"}},
			nil,
		},
		{"option name Style type combo default None var Solid var Normal var Risky", option.Select{}, exception.InvalidOptionSyntax},
		{"option name Style type combo var Solid var Normal var Risky", option.Select{}, exception.InvalidOptionSyntax},
		{"option name Style type combo default Normal", option.Select{}, exception.InvalidOptionSyntax},
	}
	for _, c := range cases {
		o, err := fu.Option([]byte(c.in))
		basicOptionMatching(t, c.in, o, c.want, err, c.err)
		switch v := o.(type) {
		case option.Select:
			if c.want.Index != v.Index || !stringutil.SliceEquals(c.want.Vars, v.Vars) {
				t.Errorf("Mismatch values.\nInput: %v\nExpected: %v\nActual: %v", c.in, c.want, v)
			}
		}
	}
}

func TestParseOpt4(t *testing.T) {
	fu := NewFromUsi()
	cases := []struct {
		in   string
		want option.Button
		err  error
	}{
		{"option name ResetLearning type button", option.Button{Name: "ResetLearning"}, nil},
		{"option name <empty> type button", option.Button{Name: "<empty>"}, nil}, // まぁいい
		{"option name ResetLearning type button sur", option.Button{}, exception.InvalidOptionSyntax},
		{"option name 1 type button", option.Button{Name: "1"}, nil},
	}
	for _, c := range cases {
		o, err := fu.Option([]byte(c.in))
		basicOptionMatching(t, c.in, o, c.want, err, c.err)
	}
}

func TestParseOpt5(t *testing.T) {
	fu := NewFromUsi()
	cases := []struct {
		in   string
		want option.String
		err  error
	}{
		{"option name BookFile type string default public.bin", option.String{Name: "BookFile", Val: "public.bin", Default: "public.bin"}, nil},
		{"option name BookFile type string default public.bin var a", option.String{}, exception.InvalidOptionSyntax},
		{"option name BookFile type string", option.String{}, exception.InvalidOptionSyntax},
		{"option name BookFile type string public.bin", option.String{}, exception.InvalidOptionSyntax},
	}
	for _, c := range cases {
		o, err := fu.Option([]byte(c.in))
		basicOptionMatching(t, c.in, o, c.want, err, c.err)
		switch v := o.(type) {
		case option.String:
			if v.Val != c.want.Val || v.Default != c.want.Default {
				t.Errorf("Mismatch values.\nInput: %v\nExpected: %v\nActual: %v", c.in, c.want, v)
			}
		}
	}
}

func TestParseOpt6(t *testing.T) {
	fu := NewFromUsi()
	cases := []struct {
		in   string
		want option.FileName
		err  error
	}{
		{"option name LearningFile type filename default <empty>", option.FileName{Name: "LearningFile", Val: "<empty>", Default: "<empty>"}, nil},
		{"option name LearningFile type filename default <empty> var a", option.FileName{}, exception.InvalidOptionSyntax},
		{"option name LearningFile type filename", option.FileName{}, exception.InvalidOptionSyntax},
		{"option name LearningFile type filename <empty>", option.FileName{}, exception.InvalidOptionSyntax},
	}
	for _, c := range cases {
		o, err := fu.Option([]byte(c.in))
		basicOptionMatching(t, c.in, o, c.want, err, c.err)
		switch v := o.(type) {
		case option.FileName:
			if v.Val != c.want.Val || v.Default != c.want.Default {
				t.Errorf("Mismatch values.\nInput: %v\nExpected: %v\nActual: %v", c.in, c.want, v)
			}
		}
	}
}

// in: input
// o1: Returned Option
// o2: Expected Option
// e1: Returned Error
// e2: Expected Error
func basicOptionMatching(t *testing.T, in string, o1, o2 option.Option, e1, e2 error) {
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
