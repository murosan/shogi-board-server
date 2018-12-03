// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package from_usi

import (
	"strings"
	"testing"

	"github.com/murosan/shogi-proxy-server/app/domain/entity/engine/result"
	"github.com/murosan/shogi-proxy-server/app/lib/test_helper"
)

func TestFromUsi_Info(t *testing.T) {
	cases := []struct {
		in   string
		want *result.Info
		err  error
	}{
		{"info time 1141 depth 3 nodes 135125 score cp -1521 pv 3a3b L*4h 4c4d",
			&result.Info{}, nil},
	}

	for i, c := range cases {
		infoHelper(t, i, c.in, c.want, c.err)
	}
}

func infoHelper(t *testing.T, i int, in string, want *result.Info, err error) {
	t.Helper()

	res, e := NewFromUsi().Info(in)

	if (e == nil && err != nil) || (e != nil && err == nil) {
		t.Errorf(`(From Usi: Paese Info) Expected error, but was not as expected.
Index:    %d
Input:    %s
Expected: %v
Actual:   %v
`, i, in, err, e)
	}

	// 想定通りのエラー
	if e != nil && strings.Contains(string(e.Error()), string(err.Error())) {
		return
	}

	// エラーだったが、想定と違った。
	if e != nil && !strings.Contains(string(e.Error()), string(err.Error())) {
		t.Errorf(`(From Usi: Paese Info) Expected error, but was not as expected.
Index:    %d
Input:    %s
Expected: %v
Actual:   %v
`, i, in, err, e)
	}

	if test_helper.InfoEquals(res, want) {
		t.Errorf(`(From Usi: Parse Info) The value was not as expected.
Index:    %d
Input:    %s
Expected: %v
Actual:   %v
`, i, in, want, res)
	}
}
