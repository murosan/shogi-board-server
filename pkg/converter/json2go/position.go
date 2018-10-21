// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package json2go

import (
	"encoding/json"

	"github.com/murosan/shogi-proxy-server/pkg/converter/models"
)

func ToPosition(j []byte) (p models.Position, e error) {
	if err := json.Unmarshal(j, &p); err != nil {
		e = err
		return
	}
	return
}
