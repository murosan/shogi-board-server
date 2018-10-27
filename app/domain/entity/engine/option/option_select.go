// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package option

import "bytes"

// USIのcombo
type Select struct {
	Name  []byte
	Index int
	Vars  [][]byte
}

func (s Select) Usi() []byte {
	return bytes.Join([][]byte{pref, s.Name, val, s.Vars[s.Index]}, space)
}

func (s Select) GetName() []byte { return s.Name }
