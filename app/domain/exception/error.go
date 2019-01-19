// Copyright 2018 murosan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package exception

import (
	"fmt"
)

// Error エラー。Code と Msg を持つ
type Error struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

// Error error インターフェースを満たすメソッド
// Code と Msg を文字列化する
func (e *Error) Error() string {
	return fmt.Sprintf("%s %s", e.Code, e.Msg)
}

// WithMsg Msg 受け取り、その Msg を含んだ新しいエラーを返す
func (e *Error) WithMsg(m string) *Error {
	return &Error{e.Code, m}
}

var (
	// http

	// NotFound 404 APIのパスが違うなど
	NotFound = &Error{Code: "NotFound"}
	// MethodNotAllowed 405 POST/GET を間違えている
	MethodNotAllowed = &Error{Code: "MethodNotAllowed"}
	// InternalServerError 503 サーバー側でエラーが発生
	InternalServerError = &Error{Code: "InternalServerError"}
	// ContentLengthRequired 411 ヘッダに Content-Length が見つからなかった
	ContentLengthRequired = &Error{Code: "ContentLengthRequired"}
	// FailedToReadBody body を必要としていたが読み取れなかった
	FailedToReadBody = &Error{Code: "FailedToReadBody"}
	// FailedToParseJSON JSON をパースしようとしたが、値が不正だった
	FailedToParseJSON = &Error{Code: "FailedToParseJSON"}

	// engine

	// ConnectionTimeout 将棋エンジンからの応答がなかったり、必要以上に時間がかかってタイムアウト
	ConnectionTimeout = &Error{Code: "ConnectionTimeout"}
	// FailedToConnect 将棋エンジンへの接続に失敗した
	FailedToConnect = &Error{Code: "FailedToConnect"}
	// FailedToClose 将棋エンジンとの接続解除に失敗した
	FailedToClose = &Error{Code: "FailedToClose"}
	// EngineIsNotRunning 将棋エンジンに接続している前提の処理を実行しようとしたが、
	// 接続前だった
	EngineIsNotRunning = &Error{Code: "EngineIsNotRunning"}
	// EngineIsAlreadyRunning 将棋エンジンに接続しようとしたが、すでに接続済だった
	EngineIsAlreadyRunning = &Error{Code: "EngineIsAlreadyRunning"}
	// FailedToUpdateOption 将棋エンジンのオプションの値を更新しようとしたが失敗した
	FailedToUpdateOption = &Error{Code: "FailedToUpdateOption"}
	// InvalidOptionSyntax オプションのシンタックスが間違っている
	// オプションを更新するときや、将棋エンジンから最初に出力されるオプション情報が不正なときに起こる
	InvalidOptionSyntax = &Error{Code: "InvalidOptionSyntax"}
	// InvalidOptionParameter オプションに渡そうとしたパラメータが不正
	// 型が違うときや、範囲外が指定された
	InvalidOptionParameter = &Error{Code: "InvalidOptionParameter"}
	// UnknownOption 不明なオプションの型
	UnknownOption = &Error{Code: "UnknownOption"}

	// usi

	// InvalidPieceID 駒のIDが不正
	InvalidPieceID = &Error{Code: "InvalidPieceID"}
	// InvalidRowNumber Row の値が範囲外
	InvalidRowNumber = &Error{Code: "InvalidRowNumber"}
	// InvalidColumnNumber Column の値が範囲外
	InvalidColumnNumber = &Error{Code: "InvalidColumnNumber"}
	// UnknownCharacter 不明な文字列
	UnknownCharacter = &Error{Code: "UnknownCharacter"}
	// FailedToParseInfo 将棋エンジンから出力される info のパースに失敗した
	FailedToParseInfo = &Error{Code: "FailedToParseInfo"}
)
