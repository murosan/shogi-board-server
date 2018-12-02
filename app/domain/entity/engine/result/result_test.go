// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package result

import "testing"

func TestResult_Flush(t *testing.T) {
	v := make(map[int]*Info)
	v[1] = &Info{}
	v[2] = &Info{}
	r := &Result{v}

	l1 := len(r.Values)
	if l1 != 2 {
		t.Errorf(`Failed to setup.`)
	}

	r.Flush()
	if len(r.Values) != 0 {
		t.Errorf(`Failed to flush values. 
Expected length is 0, but got %d`, len(r.Values))
	}
}

func TestResult_Set(t *testing.T) {
	v := make(map[int]*Info)
	r := &Result{v}
	if len(r.Values) != 0 {
		t.Errorf(`Failed to setup.`)
	}

	r.Set(&Info{}, 11)
	if len(r.Values) != 1 {
		t.Errorf("Failed to set info. Expected length is 1, but got 0")
	}
}
