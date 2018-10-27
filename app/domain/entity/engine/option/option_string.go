// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package option

import "bytes"

type String struct {
	Name, Val, Default []byte
}

func (s String) Usi() []byte {
	return bytes.Join([][]byte{pref, s.Name, val, s.Val}, space)
}

func (s String) GetName() []byte { return s.Name }
