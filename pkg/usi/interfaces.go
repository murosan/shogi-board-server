// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package usi

type position struct {
	Version uint8  `json:"version"`
	Command string `json:"command"`
	Data    struct {
		Position [9][9]int `json:"position"`
		Cap0     []int     `json:"cap_0"`
		Cap1     []int     `json:"cap_1"`
		Turn     int       `json:"turn"`
	} `json:"data"`
}
