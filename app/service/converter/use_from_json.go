// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package converter

import (
	"github.com/murosan/shogi-board-server/app/domain/entity/converter"
)

var fj = converter.NewFromJSON()

// UseFromJSON FromJSON を返す
func UseFromJSON() converter.FromJSON {
	return fj
}
