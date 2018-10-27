// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package converter

import (
	"github.com/murosan/shogi-proxy-server/app/domain/entity/converter/to_usi"
)

var tu = to_usi.NewToUsi()

func UseToUsi() *to_usi.ToUsi {
	return tu
}
