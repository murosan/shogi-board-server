// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package converter

import (
	"encoding/json"
	"fmt"

	"github.com/murosan/shogi-board-server/app/domain/entity/shogi"
	"github.com/murosan/shogi-board-server/app/domain/exception"
)

// FromJSON JSON から Go に変換するインターフェース
type FromJSON interface {
	Position([]byte) (p shogi.Position, e error)
}

type fromJSON struct{}

// NewFromJSON 新しい FromJSON を返す
func NewFromJSON() FromJSON {
	return fromJSON{}
}

func (fj fromJSON) Position(j []byte) (p shogi.Position, e error) {
	if err := json.Unmarshal(j, &p); err != nil {
		e = exception.FailedToParseJSON.WithMsg(fmt.Sprintf("caused by:\n%v", e))
	}
	return
}
