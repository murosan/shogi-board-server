package usi

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
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
