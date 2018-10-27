// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package option

import (
	"bytes"
	"strconv"
)

type Spin struct {
	Name                   []byte
	Val, Default, Min, Max int
}

func (s Spin) Usi() []byte {
	b := strconv.Itoa(s.Val)
	return bytes.Join([][]byte{pref, s.Name, val, []byte(b)}, space)
}

func (s Spin) GetName() []byte { return s.Name }
