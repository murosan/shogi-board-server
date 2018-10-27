// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package converter

import (
	"github.com/murosan/shogi-proxy-server/app/domain/entity/converter/from_json"
)

var fj = from_json.NewFromJson()

func UseFromJson() *from_json.FromJson {
	return fj
}
