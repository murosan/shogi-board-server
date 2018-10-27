// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package option

import "bytes"

type Button struct {
	Name []byte
}

func (b Button) Usi() []byte {
	return bytes.Join([][]byte{pref, b.Name}, space)
}

func (b Button) GetName() []byte { return b.Name }
