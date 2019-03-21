// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tools

// 開発用ライブラリを go modules に認識させるためのファイル

import (
	_ "google.golang.org/grpc"
	_ "github.com/golang/protobuf/protoc-gen-go"
	_ "honnef.co/go/tools/cmd/staticcheck"
	_ "github.com/rakyll/gotest"
)
