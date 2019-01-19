// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package converter

import (
	"github.com/murosan/shogi-board-server/app/domain/entity/converter"
)

var tu = converter.NewToUSI()

// UseToUSI ToUSI を返す
func UseToUSI() converter.ToUSI {
	return tu
}
