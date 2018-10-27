// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package option

import "bytes"

type FileName struct {
	Name, Val, Default []byte
}

func (f FileName) Usi() []byte {
	return bytes.Join([][]byte{pref, f.Name, val, f.Val}, space)
}

func (f FileName) GetName() []byte { return f.Name }
