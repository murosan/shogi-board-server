// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package from_json

import (
	"encoding/json"
	"fmt"

	"github.com/murosan/shogi-proxy-server/app/domain/entity/shogi"
	"github.com/murosan/shogi-proxy-server/app/domain/exception"
)

func (fj *FromJson) Position(j []byte) (p shogi.Position, e error) {
	if err := json.Unmarshal(j, &p); err != nil {
		e = exception.FailedToParseJson.WithMsg(fmt.Sprintf("caused by:\n%v", e))
	}
	return
}
