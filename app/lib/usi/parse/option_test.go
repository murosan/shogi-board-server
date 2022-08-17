package parse

import (
	"errors"
	"reflect"
	"strings"
	"testing"

	"github.com/murosan/shogi-board-server/app/domain/entity/engine"
)

var (
	errEmp = errors.New("")
)

func TestButton(t *testing.T) {
	cases := []struct {
		in   string
		want *engine.Button
		err  error
	}{
		{"option name ResetLearning type button", &engine.Button{Name: "ResetLearning"}, nil},
		{"option name <empty> type button", &engine.Button{Name: "<empty>"}, nil},
		{"option name ResetLearning type button sur", nil, errEmp},
		{"option name 1 type button", &engine.Button{Name: "1"}, nil},
	}
	for _, c := range cases {
		o, err := Button(c.in)
		basicOptionMatching(t, c.in, o, c.want, err, c.err)
	}
}

func TestCheck(t *testing.T) {
	cases := []struct {
		in   string
		want *engine.Check
		err  error
	}{
		{
			"option name UseBook type check default true",
			&engine.Check{Name: "UseBook", Value: true, Default: true},
			nil,
		},
		{"   option name UseBook type check default true   ",
			&engine.Check{Name: "UseBook", Value: true, Default: true},
			nil,
		},
		{"option name UseBook type check default ", nil, errEmp},
		{"option name UseBook type check default not_bool", nil, errEmp},
		{"option name UseBook type check dlft true", nil, errEmp},
	}

	for _, c := range cases {
		o, err := Check(c.in)
		basicOptionMatching(t, c.in, o, c.want, err, c.err)
	}
}

func TestRange(t *testing.T) {
	cases := []struct {
		in   string
		want *engine.Range
		err  error
	}{
		{
			"option name Selectivity type spin default 2 min 0 max 4",
			&engine.Range{Name: "Selectivity", Value: 2, Default: 2, Min: 0, Max: 4},
			nil,
		},
		{
			"option name Selectivity type spin default -100 min -123456 max 54321 ",
			&engine.Range{
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
			errEmp,
		},
		{
			"option name Selectivity type spin default 2",
			nil,
			errEmp,
		},
		{
			"option name Selectivity type spin min 0 max 4 default 2",
			nil,
			errEmp,
		},
		{
			"option name Selectivity type spin default a min 0 max 4",
			nil,
			errEmp,
		},
		{
			"option name Selectivity type spin default 2 min a max 4",
			nil,
			errEmp,
		},
		{
			"option name Selectivity type spin default 2 min 0 max a",
			nil,
			errEmp,
		},
		{
			"option name Selectivity type spin default 2 min 4 max 0",
			nil,
			errEmp,
		},
		{
			"option name Selectivity type spin default 7 min 0 max 4",
			nil,
			errEmp,
		},
	}
	for _, c := range cases {
		o, err := Range(c.in)
		basicOptionMatching(t, c.in, o, c.want, err, c.err)
	}
}

func TestSelect(t *testing.T) {
	cases := []struct {
		in   string
		want *engine.Select
		err  error
	}{
		{
			"option name Style type combo default Normal var Solid var Normal var Risky",
			&engine.Select{
				Name:    "Style",
				Value:   "Normal",
				Default: "Normal",
				Vars:    []string{"Solid", "Normal", "Risky"},
			},
			nil,
		},
		{
			"option name Style type combo default Nor mal var Sol id var Nor mal var R isky",
			&engine.Select{
				Name:    "Style",
				Value:   "Nor mal",
				Default: "Nor mal",
				Vars:    []string{"Sol id", "Nor mal", "R isky"},
			},
			nil,
		},
		{"option name Style type combo default None var Solid var Normal var Risky",
			nil,
			errEmp,
		},
		{"option name Style type combo var Solid var Normal var Risky",
			nil,
			errEmp,
		},
		{"option name Style type combo default Normal",
			nil,
			errEmp,
		},
		{"option name Style type combo default Normal var ",
			nil,
			errEmp,
		},
	}
	for _, c := range cases {
		o, err := Select(c.in)
		basicOptionMatching(t, c.in, o, c.want, err, c.err)
	}
}

func TestTextFromStringType(t *testing.T) {
	cases := []struct {
		in   string
		want *engine.Text
		err  error
	}{
		{"option name BookFile type string default public.bin",
			&engine.Text{
				Name:    "BookFile",
				Value:   "public.bin",
				Default: "public.bin",
			},
			nil,
		},
		{"option name BookFile type string default public.bin var a",
			&engine.Text{
				Name:    "BookFile",
				Value:   "public.bin var a",
				Default: "public.bin var a",
			},
			nil,
		},
		{"option name BookFile type string default ", &engine.Text{Name: "BookFile"}, nil},
		{"option name BookFile type string", nil, errEmp},
		{"option name BookFile type string public.bin", nil, errEmp},
	}
	for _, c := range cases {
		o, err := TextFromStringType(c.in)
		basicOptionMatching(t, c.in, o, c.want, err, c.err)
	}
}

func TestTextFromFilenameType(t *testing.T) {
	cases := []struct {
		in   string
		want *engine.Text
		err  error
	}{
		{
			"option name LearningFile type filename default <empty>",
			&engine.Text{
				Name:    "LearningFile",
				Value:   "<empty>",
				Default: "<empty>",
			},
			nil,
		},
		{"option name LearningFile type filename default <empty> var a",
			&engine.Text{
				Name:    "LearningFile",
				Value:   "<empty> var a",
				Default: "<empty> var a",
			},
			nil,
		},
		{"option name LearningFile type filename default", &engine.Text{Name: "LearningFile"}, nil},
		{"option name LearningFile type filename", nil, errEmp},
		{"option name LearningFile type filename <empty>", nil, errEmp},
	}
	for _, c := range cases {
		o, err := TextFromFilenameType(c.in)
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
	if e1 != nil && strings.Contains(e1.Error(), e2.Error()) {
		return
	}

	// an error was returned, but was not as expected
	if e1 != nil && !strings.Contains(e1.Error(), e2.Error()) {
		t.Errorf(`
Returned error was not as expected.
Input:    %v
Expected: %v
Actual:   %v`, in, e2, e1)
	}

	if !reflect.DeepEqual(o1, o2) {
		t.Errorf(`
Two options was not equal.
Input:    %s
Expected: %v
Actual:   %v
`, in, o1, o2)
	}
}
