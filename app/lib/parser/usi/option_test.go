package usi

import (
	"github.com/pkg/errors"
	"reflect"
	"strings"
	"testing"

	"github.com/murosan/shogi-board-server/app/domain/entity/option"
)

var (
	emptyErr = errors.New("")
)

func TestParseButton(t *testing.T) {
	cases := []struct {
		in   string
		want *option.Button
		err  error
	}{
		{"option name ResetLearning type button", &option.Button{Name: "ResetLearning"}, nil},
		{"option name <empty> type button", &option.Button{Name: "<empty>"}, nil},
		{"option name ResetLearning type button sur", nil, emptyErr},
		{"option name 1 type button", &option.Button{Name: "1"}, nil},
	}
	for _, c := range cases {
		o, err := ParseButton(c.in)
		basicOptionMatching(t, c.in, o, c.want, err, c.err)
	}
}

func TestParseCheck(t *testing.T) {
	cases := []struct {
		in   string
		want *option.Check
		err  error
	}{
		{
			"option name UseBook type check default true",
			&option.Check{Name: "UseBook", Value: true, Default: true},
			nil,
		},
		{"   option name UseBook type check default true   ",
			&option.Check{Name: "UseBook", Value: true, Default: true},
			nil,
		},
		{"option name UseBook type check default ", nil, emptyErr},
		{"option name UseBook type check default not_bool", nil, emptyErr},
		{"option name UseBook type check dlft true", nil, emptyErr},
	}

	for _, c := range cases {
		o, err := ParseCheck(c.in)
		basicOptionMatching(t, c.in, o, c.want, err, c.err)
	}
}

func TestParseRange(t *testing.T) {
	cases := []struct {
		in   string
		want *option.Range
		err  error
	}{
		{
			"option name Selectivity type spin default 2 min 0 max 4",
			&option.Range{Name: "Selectivity", Value: 2, Default: 2, Min: 0, Max: 4},
			nil,
		},
		{
			"option name Selectivity type spin default -100 min -123456 max 54321 ",
			&option.Range{
				Name:    "Selectivity",
				Value:   -100,
				Default: -100,
				Min:     -123456,
				Max:     54321,
			},
			nil,
		},
		{
			"option name Selectivity type spin min 0 max 4",
			nil,
			emptyErr,
		},
		{
			"option name Selectivity type spin default 2",
			nil,
			emptyErr,
		},
		{
			"option name Selectivity type spin min 0 max 4 default 2",
			nil,
			emptyErr,
		},
		{
			"option name Selectivity type spin default a min 0 max 4",
			nil,
			emptyErr,
		},
		{
			"option name Selectivity type spin default 2 min a max 4",
			nil,
			emptyErr,
		},
		{
			"option name Selectivity type spin default 2 min 0 max a",
			nil,
			emptyErr,
		},
		{
			"option name Selectivity type spin default 2 min 4 max 0",
			nil,
			emptyErr,
		},
		{
			"option name Selectivity type spin default 7 min 0 max 4",
			nil,
			emptyErr,
		},
	}
	for _, c := range cases {
		o, err := ParseRange(c.in)
		basicOptionMatching(t, c.in, o, c.want, err, c.err)
	}
}

func TestParseSelect(t *testing.T) {
	cases := []struct {
		in   string
		want *option.Select
		err  error
	}{
		{
			"option name Style type combo default Normal var Solid var Normal var Risky",
			&option.Select{
				Name:    "Style",
				Value:   "Normal",
				Default: "Normal",
				Vars:    []string{"Solid", "Normal", "Risky"},
			},
			nil,
		},
		{
			"option name Style type combo default Nor mal var Sol id var Nor mal var R isky",
			&option.Select{
				Name:    "Style",
				Value:   "Nor mal",
				Default: "Nor mal",
				Vars:    []string{"Sol id", "Nor mal", "R isky"},
			},
			nil,
		},
		{"option name Style type combo default None var Solid var Normal var Risky",
			nil,
			emptyErr,
		},
		{"option name Style type combo var Solid var Normal var Risky",
			nil,
			emptyErr,
		},
		{"option name Style type combo default Normal",
			nil,
			emptyErr,
		},
		{"option name Style type combo default Normal var ",
			nil,
			emptyErr,
		},
	}
	for _, c := range cases {
		o, err := ParseSelect(c.in)
		basicOptionMatching(t, c.in, o, c.want, err, c.err)
	}
}

func TestParseTextFromStringType(t *testing.T) {
	cases := []struct {
		in   string
		want *option.Text
		err  error
	}{
		{"option name BookFile type string default public.bin",
			&option.Text{
				Name:    "BookFile",
				Value:   "public.bin",
				Default: "public.bin",
			},
			nil,
		},
		{"option name BookFile type string default public.bin var a",
			&option.Text{
				Name:    "BookFile",
				Value:   "public.bin var a",
				Default: "public.bin var a",
			},
			nil,
		},
		{"option name BookFile type string",
			nil,
			emptyErr,
		},
		{"option name BookFile type string public.bin",
			nil,
			emptyErr,
		},
	}
	for _, c := range cases {
		o, err := ParseTextFromStringType(c.in)
		basicOptionMatching(t, c.in, o, c.want, err, c.err)
	}
}

func TestParseTextFromFilenameType(t *testing.T) {
	cases := []struct {
		in   string
		want *option.Text
		err  error
	}{
		{
			"option name LearningFile type filename default <empty>",
			&option.Text{
				Name:    "LearningFile",
				Value:   "<empty>",
				Default: "<empty>",
			},
			nil,
		},
		{"option name LearningFile type filename default <empty> var a",
			&option.Text{
				Name:    "LearningFile",
				Value:   "<empty> var a",
				Default: "<empty> var a",
			},
			nil,
		},
		{"option name LearningFile type filename",
			nil,
			emptyErr,
		},
		{"option name LearningFile type filename <empty>",
			nil,
			emptyErr,
		},
	}
	for _, c := range cases {
		o, err := ParseTextFromFilenameType(c.in)
		basicOptionMatching(t, c.in, o, c.want, err, c.err)
	}
}

// in: input
// o1: Returned Option
// o2: Expected Option
// e1: Returned Error
// e2: Expected Error
func basicOptionMatching(t *testing.T, in string, o1, o2 interface{}, e1, e2 error) {
	t.Helper()
	if (e1 == nil && e2 != nil) || (e1 != nil && e2 == nil) {
		t.Errorf(`
Returned error was not as expected.
Input:    %v
Expected: %v
Actual:   %v`, in, e2, e1)
	}

	// returned expected error
	if e1 != nil && strings.Contains(string(e1.Error()), string(e2.Error())) {
		return
	}

	// an error was returned, but was not as expected
	if e1 != nil && !strings.Contains(string(e1.Error()), string(e2.Error())) {
		t.Errorf(`
Returned error was not as expected.
Input:    %v
Expected: %v
Actual:   %v`, in, e2, e1)
	}

	if !reflect.DeepEqual(o1, o2) {
		t.Errorf(`
Marshaled value (json bytes) was not as expected.
Input:    %s
Expected: %v
Actual:   %v
`, in, o1, o2)
	}
}
