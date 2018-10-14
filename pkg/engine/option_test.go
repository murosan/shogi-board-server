// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package engine

import (
	"bytes"
	"strings"
	"testing"

	"github.com/murosan/shogi-proxy-server/pkg/lib"
	"github.com/murosan/shogi-proxy-server/pkg/msg"
)

func TestClient_ParseId(t *testing.T) {
	cases := []struct {
		in   string
		name []byte
		err  error
	}{
		{"id name computer_name", []byte("computer_name"), nil},
		{"id name ", []byte(""), msg.InvalidIdSyntax},
		{"id neimu typo_key", []byte(""), msg.UnknownOption},
	}
	authorCases := []struct {
		in   string
		name []byte
		err  error
	}{
		{"id author computer_author", []byte("computer_author"), nil},
		{"id author", []byte(""), msg.InvalidIdSyntax},
		{"id auther typo_key", []byte(""), msg.UnknownOption},
	}

	for _, c := range cases {
		e := Client{}
		err := e.ParseId([]byte(c.in))
		if err != c.err {
			t.Errorf("Returned error was not as expected.\nInput: %v, Expected: %v, Actual: %v", c.in, c.err, err)
		}
		if !bytes.Equal(e.Name, c.name) {
			t.Errorf("Computer name name was not as expected.\nInput: %v, Expected: %v, Actual: %v", c.in, c.name, e.Name)
		}
	}

	for _, c := range authorCases {
		e := Client{}
		err := e.ParseId([]byte(c.in))
		if err != c.err {
			t.Errorf("Returned error was not as expected.\nInput: %v, Expected: %v, Actual: %v", c.in, c.err, err)
		}
		if !bytes.Equal(e.Author, c.name) {
			t.Errorf("Author name was not as expected.\nInput: %v, Expected: %v, Actual: %v", c.in, c.name, e.Name)
		}
	}
}

func TestClient_ParseOpt(t *testing.T) {
	cases := []struct {
		in   string
		want Check
		err  error
	}{
		{"option name UseBook type check default true", Check{[]byte("UseBook"), true, true}, nil},
		{"   option name UseBook type check default true   ", Check{[]byte("UseBook"), true, true}, nil},
		{"option name UseBook type check default ", Check{}, msg.InvalidOptionSyntax},
		{"option name UseBook type check default not_bool", Check{}, msg.InvalidOptionSyntax},
		{"option name UseBook type check dlft true", Check{}, msg.InvalidOptionSyntax},
	}

	for _, c := range cases {
		e := Client{Options: make(map[string]Option)}
		err := e.ParseOpt([]byte(c.in))
		basicOptionMatching(t, &e, c.in, string(c.want.Name), c.want, c.err, err)
		o, ok := e.Options[string(c.want.Name)]
		if ok {
			switch v := o.(type) {
			case Check:
				if c.want.Default != v.Default || c.want.Val != v.Val {
					t.Errorf("Mismatch values.\nv1: %v\nv2: %v", c.want, v)
				}
			}
		}
	}
}

func TestClient_ParseOpt2(t *testing.T) {
	cases := []struct {
		in   string
		want Spin
		err  error
	}{
		{"option name Selectivity type spin default 2 min 0 max 4", Spin{[]byte("Selectivity"), 2, 2, 0, 4}, nil},
		{"option name Selectivity type spin default -100 min -123456 max 54321 ", Spin{[]byte("Selectivity"), -100, -100, -123456, 54321}, nil},
		{"option name Selectivity type spin min 0 max 4", Spin{}, msg.InvalidOptionSyntax},
		{"option name Selectivity type spin default 2", Spin{}, msg.InvalidOptionSyntax},
		{"option name Selectivity type spin min 0 max 4 default 2", Spin{}, msg.InvalidOptionSyntax},
		{"option name Selectivity type spin default two min 0 max 4", Spin{}, msg.InvalidOptionSyntax},
	}

	for _, c := range cases {
		e := Client{Options: make(map[string]Option)}
		err := e.ParseOpt([]byte(c.in))
		basicOptionMatching(t, &e, c.in, string(c.want.Name), c.want, c.err, err)
		o, ok := e.Options[string(c.want.Name)]
		if ok {
			switch v := o.(type) {
			case Spin:
				if c.want.Default != v.Default || c.want.Val != v.Val || c.want.Min != v.Min || c.want.Max != v.Max {
					t.Errorf("Mismatch values.\nv1: %v\nv2: %v", c.want, v)
				}
			}
		}
	}
}

func TestClient_ParseOpt3(t *testing.T) {
	cases := []struct {
		in   string
		want Select
		err  error
	}{
		{
			"option name Style type combo default Normal var Solid var Normal var Risky",
			Select{[]byte("Style"), 1, [][]byte{[]byte("Solid"), []byte("Normal"), []byte("Risky")}},
			nil,
		},
		{"option name Style type combo default None var Solid var Normal var Risky", Select{}, msg.InvalidOptionSyntax},
		{"option name Style type combo var Solid var Normal var Risky", Select{}, msg.InvalidOptionSyntax},
		{"option name Style type combo default Normal", Select{}, msg.InvalidOptionSyntax},
	}

	for _, c := range cases {
		e := Client{Options: make(map[string]Option)}
		err := e.ParseOpt([]byte(c.in))
		basicOptionMatching(t, &e, c.in, string(c.want.Name), c.want, c.err, err)
		o, ok := e.Options[string(c.want.Name)]
		if ok {
			switch v := o.(type) {
			case Select:
				if c.want.Index != v.Index || !lib.EqualBytes(c.want.Vars, v.Vars) {
					t.Errorf("Mismatch values.\nv1: %v\nv2: %v", c.want, v)
				}
			}
		}
	}
}

func TestClient_ParseOpt4(t *testing.T) {
	cases := []struct {
		in   string
		want Button
		err  error
	}{
		{"option name ResetLearning type button", Button{[]byte("ResetLearning")}, nil},
		{"option name <empty> type button", Button{[]byte("<empty>")}, nil}, // まぁいい
		{"option name ResetLearning type button sur", Button{}, msg.InvalidOptionSyntax},
		{"option name 1 type button", Button{[]byte("1")}, nil},
	}

	for _, c := range cases {
		e := Client{Options: make(map[string]Option)}
		err := e.ParseOpt([]byte(c.in))
		basicOptionMatching(t, &e, c.in, string(c.want.Name), c.want, c.err, err)
	}
}

func TestClient_ParseOpt5(t *testing.T) {
	cases := []struct {
		in   string
		want String
		err  error
	}{
		{"option name BookFile type string default public.bin", String{[]byte("BookFile"), []byte("public.bin"), []byte("public.bin")}, nil},
		{"option name BookFile type string default public.bin var a", String{}, msg.InvalidOptionSyntax},
		{"option name BookFile type string", String{}, msg.InvalidOptionSyntax},
		{"option name BookFile type string public.bin", String{}, msg.InvalidOptionSyntax},
	}

	for _, c := range cases {
		e := Client{Options: make(map[string]Option)}
		err := e.ParseOpt([]byte(c.in))
		basicOptionMatching(t, &e, c.in, string(c.want.Name), c.want, c.err, err)
		o, ok := e.Options[string(c.want.Name)]
		if ok {
			switch v := o.(type) {
			case String:
				if !bytes.Equal(v.Val, c.want.Val) || !bytes.Equal(v.Default, c.want.Default) {
					t.Errorf("Mismatch values.\nv1: %v\nv2: %v", c.want, v)
				}
			}
		}
	}
}

func TestClient_ParseOpt6(t *testing.T) {
	cases := []struct {
		in   string
		want FileName
		err  error
	}{
		{"option name LearningFile type filename default <empty>", FileName{[]byte("LearningFile"), []byte("<empty>"), []byte("<empty>")}, nil},
		{"option name LearningFile type filename default <empty> var a", FileName{}, msg.InvalidOptionSyntax},
		{"option name LearningFile type filename", FileName{}, msg.InvalidOptionSyntax},
		{"option name LearningFile type filename <empty>", FileName{}, msg.InvalidOptionSyntax},
	}

	for _, c := range cases {
		e := Client{Options: make(map[string]Option)}
		err := e.ParseOpt([]byte(c.in))
		basicOptionMatching(t, &e, c.in, string(c.want.Name), c.want, c.err, err)
		o, ok := e.Options[string(c.want.Name)]
		if ok {
			switch v := o.(type) {
			case FileName:
				if !bytes.Equal(v.Val, c.want.Val) || !bytes.Equal(v.Default, c.want.Default) {
					t.Errorf("Mismatch values.\nv1: %v\nv2: %v", c.want, v)
				}
			}
		}
	}
}

// in: input
// na: Name
// o:  Option
// e1: cases のエラー(Expected Error)
// e2: ParseOpt の戻り値のエラー(Actual Error)
func basicOptionMatching(t *testing.T, e *Client, in, na string, o Option, ce, pe error) {
	t.Helper()
	if (ce == nil && pe != nil) || (ce != nil && pe == nil) {
		t.Errorf("Returned error was not as expected.\nInput: %v\nExpected: %v\nActual: %v", in, ce, pe)
	}

	// 予想通りのエラーが返った
	if pe != nil && strings.Contains(string(pe.Error()), string(ce.Error())) {
		return
	}

	// エラーは返ったが、想定と違った
	if pe != nil && !strings.Contains(string(pe.Error()), string(ce.Error())) {
		t.Errorf("Returned error was not as expected.\nInput: %v\nExpected: %v\nActual: %v", in, ce, pe)
	}

	// 正しくOptionがセットされたかどうか
	opt, ok := e.Options[na]
	if !ok {
		t.Errorf("Option was not set.\nInput: %v", in)
	}

	// USIコマンドが想定通りかどうか
	if !bytes.Equal(opt.Usi(), o.Usi()) {
		t.Errorf("Set value was not as expected.\nInput: %v\nExpected: %v\nActual: %v", in, string(opt.Usi()), string(o.Usi()))
	}
}
