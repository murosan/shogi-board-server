// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package connector

type Connector interface {
	// 将棋エンジンと接続する
	Connect() error

	// 接続を切る
	Close() error

	// 将棋エンジンにコマンドを実行する
	Exec(*[]byte) error
}
