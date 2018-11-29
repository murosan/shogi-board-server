// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package from_usi

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"github.com/murosan/shogi-proxy-server/app/domain/entity/engine/option"
	"github.com/murosan/shogi-proxy-server/app/domain/exception"
)

var emp = ""

func TestFromUsi_EngineID(t *testing.T) {
	fu := NewFromUsi()
	cases := []struct {
		in, name, val string
		err           error
	}{
		{"id name computer_name", name, "computer_name", nil},
		{"id name ", name, emp, nil},
		{"id name", emp, emp, exception.UnknownOption},
		{"id neimu typo_key", emp, emp, exception.UnknownOption},
		{"id author computer_author", author, "computer_author", nil},
		{"id author", emp, emp, exception.UnknownOption},
		{"id auther typo_key", emp, emp, exception.UnknownOption},
	}

	for _, c := range cases {
		n, v, e := fu.EngineID(c.in)
		if e != c.err {
			t.Errorf(`
Returned error was not as expected.
Input:    %v
Expected: %v
Actual:   %v`, c.in, c.err, e)
		}
		if n != c.name {
			t.Errorf(`
name of OptionId was not as expected.
Input:    %v
Expected: %v
Actual:   %v`, c.in, c.name, n)
		}
		if v != c.val {
			t.Errorf(`
Value of OptionId was not as expected.
Input:    %v
Expected: %v
Actual:   %v`, c.in, c.val, v)
		}
	}
}

func TestFromUsi_Option(t *testing.T) {
	fu := NewFromUsi()
	cases := []struct {
		in   string
		want *option.Button
		err  error
	}{
		{"option name ResetLearning type button", option.NewButton("ResetLearning"), nil},
		{"option name <empty> type button", option.NewButton("<empty>"), nil}, // まぁいい
		{"option name ResetLearning type button sur", nil, exception.InvalidOptionSyntax},
		{"option name 1 type button", option.NewButton("1"), nil},
	}
	for _, c := range cases {
		o, err := fu.Option(c.in)
		basicOptionMatching(t, c.in, o, c.want, err, c.err)
	}
}

func TestFromUsi_Option2(t *testing.T) {
	fu := NewFromUsi()
	cases := []struct {
		in   string
		want *option.Check
		err  error
	}{
		{"option name UseBook type check default true", option.NewCheck("UseBook", true, true), nil},
		{"   option name UseBook type check default true   ", option.NewCheck("UseBook", true, true), nil},
		{"option name UseBook type check default ", nil, exception.InvalidOptionSyntax},
		{"option name UseBook type check default not_bool", nil, exception.InvalidOptionSyntax},
		{"option name UseBook type check dlft true", nil, exception.InvalidOptionSyntax},
	}

	for _, c := range cases {
		o, err := fu.Option(c.in)
		basicOptionMatching(t, c.in, o, c.want, err, c.err)
	}
}

func TestFromUsi_Option3(t *testing.T) {
	fu := NewFromUsi()
	cases := []struct {
		in   string
		want *option.Spin
		err  error
	}{
		{"option name Selectivity type spin default 2 min 0 max 4", option.NewSpin("Selectivity", 2, 2, 0, 4), nil},
		{"option name Selectivity type spin default -100 min -123456 max 54321 ", option.NewSpin("Selectivity", -100, -100, -123456, 54321), nil},
		{"option name Selectivity type spin min 0 max 4", nil, exception.InvalidOptionSyntax},
		{"option name Selectivity type spin default 2", nil, exception.InvalidOptionSyntax},
		{"option name Selectivity type spin min 0 max 4 default 2", nil, exception.InvalidOptionSyntax},
		{"option name Selectivity type spin default two min 0 max 4", nil, exception.InvalidOptionSyntax},
	}
	for _, c := range cases {
		o, err := fu.Option(c.in)
		basicOptionMatching(t, c.in, o, c.want, err, c.err)
	}
}

func TestFromUsi_Option4(t *testing.T) {
	fu := NewFromUsi()
	cases := []struct {
		in   string
		want *option.Select
		err  error
	}{
		{
			"option name Style type combo default Normal var Solid var Normal var Risky",
			option.NewSelect("Style", "Normal", "Normal", []string{"Solid", "Normal", "Risky"}),
			nil,
		},
		{
			"option name Style type combo default Nor mal var Sol id var Nor mal var R isky",
			option.NewSelect("Style", "Nor mal", "Nor mal", []string{"Sol id", "Nor mal", "R isky"}),
			nil,
		},
		{"option name Style type combo default None var Solid var Normal var Risky", nil, exception.InvalidOptionSyntax},
		{"option name Style type combo var Solid var Normal var Risky", nil, exception.InvalidOptionSyntax},
		{"option name Style type combo default Normal", nil, exception.InvalidOptionSyntax},
	}
	for _, c := range cases {
		o, err := fu.Option(c.in)
		basicOptionMatching(t, c.in, o, c.want, err, c.err)
	}
}

func TestFromUsi_Option5(t *testing.T) {
	fu := NewFromUsi()
	cases := []struct {
		in   string
		want *option.String
		err  error
	}{
		{"option name BookFile type string default public.bin", option.NewString("BookFile", "public.bin", "public.bin"), nil},
		{"option name BookFile type string default public.bin var a", option.NewString("BookFile", "public.bin var a", "public.bin var a"), nil},
		{"option name BookFile type string", nil, exception.InvalidOptionSyntax},
		{"option name BookFile type string public.bin", nil, exception.InvalidOptionSyntax},
	}
	for _, c := range cases {
		o, err := fu.Option(c.in)
		basicOptionMatching(t, c.in, o, c.want, err, c.err)
	}
}

func TestFromUsi_Option6(t *testing.T) {
	fu := NewFromUsi()
	cases := []struct {
		in   string
		want *option.FileName
		err  error
	}{
		{"option name LearningFile type filename default <empty>", option.NewFileName("LearningFile", "<empty>", "<empty>"), nil},
		{"option name LearningFile type filename default <empty> var a", option.NewFileName("LearningFile", "<empty> var a", "<empty> var a"), nil},
		{"option name LearningFile type filename", nil, exception.InvalidOptionSyntax},
		{"option name LearningFile type filename <empty>", nil, exception.InvalidOptionSyntax},
	}
	for _, c := range cases {
		o, err := fu.Option(c.in)
		basicOptionMatching(t, c.in, o, c.want, err, c.err)
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
		t.Errorf(`
Returned error was not as expected.
Input:    %v
Expected: %v
Actual:   %v`, in, e2, e1)
	}

	// 予想通りのエラーが返った
	if e1 != nil && strings.Contains(string(e1.Error()), string(e2.Error())) {
		return
	}

	// エラーは返ったが、想定と違った
	if e1 != nil && !strings.Contains(string(e1.Error()), string(e2.Error())) {
		t.Errorf(`
Returned error was not as expected.
Input:    %v
Expected: %v
Actual:   %v`, in, e2, e1)
	}

	// USIコマンドが想定通りかどうか
	usi1 := o1.Usi()
	usi2 := o2.Usi()
	if usi1 != usi2 {
		t.Errorf(`
Update value was not as expected.
Input:    %v
Expected: %v
Actual:   %v`, in, usi2, usi1)
	}

	// json化した値が同等かどうか
	json1, _ := json.MarshalIndent(o1, "", "  ")
	json2, _ := json.MarshalIndent(o2, "", "  ")
	if !bytes.Equal(json1, json2) {
		t.Errorf(`
Marshaled value (json bytes) was not as expected.
Input:    %s
Expected: %s
Actual:   %s
`, in, string(json2), string(json1))
	}
}
