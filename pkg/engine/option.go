package engine

import (
	"bytes"
	"github.com/murosan/shogi-proxy-server/pkg/msg"
)

type Option interface {
	Usi() string
}

var (
	space  = []byte(" ")
	id     = []byte("id")
	opt    = []byte("option")
	name   = []byte("name")
	author = []byte("author")
	tpe    = []byte("type")
)

// id name <EngineName>
// id author <AuthorName> をEngineにセットする
// EngineNameやAuthorNameにスペースが入る場合もあるのでJoinしている
func ParseId(b []byte) error {
	s := bytes.Split(bytes.TrimSpace(b), space)
	if len(s) < 3 || bytes.Equal(s[0], id) {
		return msg.InvalidIdSyntax
	}

	if bytes.Equal(s[1], name) {
		Engine.Name = bytes.Join(s[2:], space)
		return nil
	}

	if bytes.Equal(s[1], author) {
		Engine.Author = bytes.Join(s[2:], space)
		return nil
	}

	return msg.UnknownOption
}

// 一行受け取って、EngineのOptionMapにセットする
// パースできなかったらエラーを返す
func ParseOpt(b []byte) error {
	s := bytes.Split(bytes.TrimSpace(b), space)
	if len(s) < 5 || bytes.Equal(s[0], opt) || bytes.Equal(s[3], tpe) {
		return msg.InvalidOptionSyntax
	}

	// TODO:
}
