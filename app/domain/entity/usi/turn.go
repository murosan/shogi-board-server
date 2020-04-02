// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package usi

// Turn is turn of the game in USI expression.
type Turn string

const (
	// Sente is the first player in USI expression.
	Sente Turn = "b"

	// Gote is the second player in USI expression.
	Gote Turn = "w"
)
