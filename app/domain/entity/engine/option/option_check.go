// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package option

import (
	"bytes"
	"strconv"
)

type Check struct {
	Name         []byte
	Val, Default bool
}

func (c Check) Usi() []byte {
	b := []byte(strconv.FormatBool(c.Val))
	return bytes.Join([][]byte{pref, c.Name, val, b}, space)
}

func (c Check) GetName() []byte { return c.Name }
