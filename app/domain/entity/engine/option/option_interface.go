// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package option

var (
	space = []byte(" ")
	val   = []byte("value")
	pref  = []byte("setoption name")
)

type Option interface {
	Usi() []byte
	GetName() []byte
}
