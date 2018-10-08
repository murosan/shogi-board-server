// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package engine

import (
	"bufio"
	"log"
)

// Engine の出力を読み取る人
func CatchEngineOutput() {
	defer func() {}()
	s := bufio.NewScanner(Engine.Stdout)

	for s.Scan() {
		b := s.Bytes()
		Engine.EngineOut <- b
	}

	if s.Err() != nil {
		log.Println("scan: ", s.Err())
	}
	close(Engine.Done)
}
