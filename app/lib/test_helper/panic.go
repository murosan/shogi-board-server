// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package testhelper

import (
	"fmt"
	"testing"
)

// MustPanic テスト用メソッド
// Panic が発生するかどうかをテストする
// Panic が発生しない場合は、errMsg を添えて Fail とする
func MustPanic(t *testing.T, f func(), errMsg string) {
	t.Helper()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered", r)
		}
	}()
	f()
	t.Error(errMsg)
}
