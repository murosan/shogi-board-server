// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package converter

import "github.com/murosan/shogi-proxy-server/app/domain/entity/converter/from_usi"

// FromUsi を interface にして nil をぶっこむ
// lazy 的な感じで initialize する
var fu = from_usi.NewFromUsi()

func UseFromUsi() *from_usi.FromUsi {
	return fu
}
