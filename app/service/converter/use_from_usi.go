// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package converter

import "github.com/murosan/shogi-board-server/app/domain/entity/converter"

// FromUsi を interface にして nil をぶっこむ
// lazy 的な感じで initialize する
var fu = converter.NewFromUSI()

// UseFromUSI FromUSI を返す
func UseFromUSI() converter.FromUSI {
	return fu
}
