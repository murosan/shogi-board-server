// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build tools

package tools

// 開発用ライブラリを go modules に認識させるためのファイル

import (
	_ "github.com/oxequa/realize"
	_ "github.com/rakyll/gotest"
	_ "honnef.co/go/tools/cmd/staticcheck"
)
